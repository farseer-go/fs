package batchFileWriter

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/snc"
	"github.com/vmihailenco/msgpack/v5"
	// "google.golang.org/protobuf/proto"
)

// SerializeType 序列化方式
type SerializeType int

const (
	SerializeJSON        SerializeType = iota // 使用 JSON 序列化（默认）
	SerializeMessagePack                      // 使用 MessagePack 序列化
	//SerializeProtobuf                         // 使用 Protobuf 序列化
)

// BatchFileWriter 是一个高性能的批量文件写入器，支持基于时间和大小的文件滚动，以及文件数量限制。它通过异步队列和缓冲写入来优化性能，并且在关闭时确保所有数据都被正确处理。
type BatchFileWriter struct {
	dataChan         chan []byte    // 明确只接收处理好的字节流
	dir              string         // 存储目录
	fileExtension    string         // 文件扩展名
	rollingInterval  string         // 文件滚动间隔（year/month/day/week/hour）
	fileSizeLimitMb  int64          // 文件大小限制（MB），设置之后，只有当文件大小超过限制时才会滚动
	fileCountLimit   int            // 文件数量限制
	interval         time.Duration  // 多长时间刷盘一次
	currentFileIndex int            // 当前时间段的文件索引号
	currentFileName  string         // 当前文件名（不含扩展名）
	currentPath      string         // 当前文件的完整路径
	bufferSize       int            // 写入缓冲区大小
	appendNewLine    bool           // 是否在每条日志后添加换行符
	serializeType    SerializeType  // 序列化方式（JSON 或 Protobuf）
	wg               sync.WaitGroup // 调用Close退出时等待所有翻转协程完成
	closeOnce        sync.Once      // 确保 Close 方法只执行一次
	closed           atomic.Bool    // 标志位，表示是否已关闭
	exitChan         chan struct{}  // 退出信号
}

// dir: 存储目录
// fileExtension: 文件扩展名
// rollingInterval: 文件的名称定义规则（year/month/day/week/hour）
// fileSizeLimitMb: 限制文件大小(单位MB)
// fileCountLimit: 限制文件数量
// interval: 多长时间刷盘(不翻转文件)
// serializeType: 序列化方式（SerializeJSON 或 SerializeProtobuf）
func NewWriter(dir string, fileExtension string, rollingInterval string, fileSizeLimitMb int64, fileCountLimit int, interval time.Duration, appendNewLine bool, serializeType SerializeType) *BatchFileWriter {
	// 创建目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0755)
	}
	// 移除扩展名前缀.
	fileExtension = strings.TrimPrefix(fileExtension, ".")

	w := &BatchFileWriter{
		dataChan:        make(chan []byte, 20000),
		dir:             dir,
		fileExtension:   fileExtension,
		rollingInterval: rollingInterval,
		fileSizeLimitMb: fileSizeLimitMb,
		fileCountLimit:  fileCountLimit,
		interval:        interval,
		bufferSize:      256 * 1024,
		appendNewLine:   appendNewLine,
		serializeType:   serializeType,
		exitChan:        make(chan struct{}),
	}

	// 初始化：获取当前时间段的文件名和索引
	w.currentFileName = w.getFilename()
	w.currentFileIndex, w.currentPath = w.initCurrentFile()

	// 开启翻转协程
	w.wg.Add(1)
	go w.daemon()
	return w
}

// getCurrentFilePath 获取当前文件的完整路径
func (w *BatchFileWriter) getCurrentFilePath() string {
	return filepath.Join(w.dir, fmt.Sprintf("%s_%d.%s", w.currentFileName, w.currentFileIndex, w.fileExtension))
}

// Write 方法在业务协程中完成类型转换，利用并发 CPU 提高吞吐量
func (w *BatchFileWriter) Write(data any) {
	var payload []byte

	// 1. 立即进行类型判定和转换
	switch v := data.(type) {
	case []byte:
		payload = v
	case string:
		payload = []byte(v)
	default:
		var err error
		switch w.serializeType {
		// case SerializeProtobuf:
		// 	// Protobuf 序列化：要求数据实现 proto.Message 接口
		// 	if msg, ok := data.(proto.Message); ok {
		// 		payload, err = proto.Marshal(msg)
		// 		if err != nil {
		// 			fmt.Println("BatchFileWriter转成protobuf时失败: ", err.Error())
		// 		}
		// 	} else {
		// 		fmt.Println("BatchFileWriter转成protobuf时失败: 数据未实现 proto.Message 接口")
		// 	}
		case SerializeMessagePack:
			// MessagePack 序列化
			payload, err = msgpack.Marshal(v)
			if err != nil {
				fmt.Println("BatchFileWriter转成msgpack时失败: ", err.Error())
			}
		default:
			// 默认 JSON 序列化
			payload, err = snc.Marshal(v)
			if err != nil {
				fmt.Println("BatchFileWriter转成json时失败: ", err.Error())
			}
		}
	}

	if len(payload) == 0 || w.closed.Load() {
		return
	}

	// 2. 将处理好的字节流放入异步队列
	select {
	case w.dataChan <- payload:
	case <-w.exitChan:
		fmt.Println("BatchFileWriter.Write: 已关闭，丢弃数据")
	default:
		fmt.Println("BatchFileWriter.Write: 由于队列已满，丢弃数据")
	}
}

// 拿到文件句柄和当前文件大小
func (w *BatchFileWriter) openFile(fileName string) (*os.File, int64, error) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, 0, errors.New("BatchFileWriter.openFile 打开文件失败: " + err.Error())
	}
	// 得到当前文件的大小
	fInfo, err := f.Stat()
	if err != nil {
		return nil, 0, errors.New("BatchFileWriter.openFile.Stat 统计文件失败: " + err.Error())
	}
	return f, fInfo.Size(), nil
}

func (w *BatchFileWriter) write(bufWriter *bufio.Writer, line []byte) int64 {
	if w.serializeType == SerializeMessagePack {
		// msgpack 是二进制格式，内部可能含有 0x0A（'\n'）字节。
		// 不能依赖换行符分隔记录，改用 4 字节大端长度前缀分帧：
		//   [uint32 BE: N][N bytes msgpack payload]
		// 读取端只需先读 4 字节得到长度，再精确读取 N 字节即可还原完整记录。
		var lenBuf [4]byte
		binary.BigEndian.PutUint32(lenBuf[:], uint32(len(line)))
		bufWriter.Write(lenBuf[:])
		bufWriter.Write(line)
		return int64(4 + len(line))
	}
	if w.appendNewLine {
		bufWriter.Write(line)
		bufWriter.WriteByte('\n')
		return int64(len(line) + 1)
	}
	bufWriter.Write(line)
	return int64(len(line))
}

// 翻转文件
func (w *BatchFileWriter) daemon() {
	defer w.wg.Done()

	f, fileSize, err := w.openFile(w.currentPath)
	if err != nil {
		panic("BatchFileWriter.daemon: 无法打开文件，退出翻转协程:" + err.Error())
	}
	bufWriter := bufio.NewWriterSize(f, w.bufferSize)

	var ticker *time.Ticker
	if w.interval == 0 {
		// 创建一个非常长的定时器，基本上不会触发，避免不必要的资源占用
		ticker = time.NewTicker(time.Hour * 24 * 365)
	} else {
		ticker = time.NewTicker(w.interval)
	}
	defer ticker.Stop()

	for {
		select {
		case line := <-w.dataChan:
			fileSize += w.write(bufWriter, line)

			// 如果文件大小限制开启，且当前文件大小超过限制，则进行翻转
			if w.fileSizeLimitMb > 0 && fileSize >= w.fileSizeLimitMb*1024*1024 {
				f, fileSize = w.rotate(bufWriter, f)
			}
		case <-ticker.C:
			// 检查时间段是否变化
			newFileName := w.getFilename()
			if newFileName != w.currentFileName {
				// 时间段变化，翻转到新文件
				f, fileSize = w.rotate(bufWriter, f)
			} else if w.interval > 0 { // 时间到了,则只刷盘,不用翻转文件
				if err := bufWriter.Flush(); err != nil {
					fmt.Println("BatchFileWriter: Flush失败:", err.Error())
				}
				if err := f.Sync(); err != nil {
					fmt.Println("BatchFileWriter: Sync失败:", err.Error())
				}
			}
		case <-w.exitChan:
			// 退出前排空 Channel
			close(w.dataChan)
			for line := range w.dataChan {
				fileSize += w.write(bufWriter, line)

				// 如果文件大小限制开启，且当前文件大小超过限制，则进行翻转
				if w.fileSizeLimitMb > 0 && fileSize >= w.fileSizeLimitMb*1024*1024 {
					f, fileSize = w.rotate(bufWriter, f)
				}
			}

			if err := bufWriter.Flush(); err != nil {
				fmt.Println("BatchFileWriter: Flush失败:", err.Error())
			}
			if err := f.Sync(); err != nil {
				fmt.Println("BatchFileWriter: Sync失败:", err.Error())
			}
			_ = f.Close()

			// 如果开启了文件数量限制，则在翻转时检查是否需要删除旧文件
			if w.fileCountLimit > 0 {
				w.removeLimitFile()
			}
			return
		}
	}
}

// 写入到文件，并在需要时翻转文件
func (w *BatchFileWriter) rotate(bufWriter *bufio.Writer, f *os.File) (*os.File, int64) {
	if err := bufWriter.Flush(); err != nil {
		fmt.Println("BatchFileWriter: Flush失败:", err.Error())
	}
	if err := f.Sync(); err != nil {
		fmt.Println("BatchFileWriter: Sync失败:", err.Error())
	}
	_ = f.Close()

	// 检查时间段是否变化
	newFileName := w.getFilename()
	if newFileName != w.currentFileName {
		// 时间段变化，查找新时间段的最新文件
		w.currentFileName = newFileName
		w.currentFileIndex, w.currentPath = w.initCurrentFile()
	} else {
		// 时间段没变，索引+1
		w.currentFileIndex++
		w.currentPath = w.getCurrentFilePath()
	}

	// 支持重试打开文件，避免因磁盘暂时不可用导致的翻转失败
	var err error
	for {
		if f, _, err = w.openFile(w.currentPath); err == nil {
			break
		}
		fmt.Println("BatchFileWriter.rotate: 打开文件失败, 3秒后重试:", err.Error())
		time.Sleep(3 * time.Second)
	}
	bufWriter.Reset(f)

	// 如果开启了文件数量限制，则在翻转时检查是否需要删除旧文件
	if w.fileCountLimit > 0 {
		w.removeLimitFile()
	}
	return f, 0
}

func (w *BatchFileWriter) Close() {
	w.closeOnce.Do(func() {
		w.closed.Store(true) // 先设置标志
		close(w.exitChan)
		w.wg.Wait()
	})
}

// 获取文件名称
func (w *BatchFileWriter) getFilename() string {
	// 根据文件间隔来确定文件名称前缀
	var fileName string
	switch strings.ToLower(w.rollingInterval) {
	case "hour":
		fileName = time.Now().Format("2006-01-02-15")
	case "day":
		fileName = time.Now().Format("2006-01-02")
	case "week":
		year, week := time.Now().ISOWeek()
		fileName = fmt.Sprint(year, "-", week)
	case "month":
		fileName = time.Now().Format("2006-01")
	case "year":
		fileName = time.Now().Format("2006")
	default:
		fileName = time.Now().Format("2006-01-02") // day
	}
	return fileName
}

// 初始化当前文件：查找最新文件，如果未超过大小限制则复用，否则创建新文件
func (w *BatchFileWriter) initCurrentFile() (int, string) {
	fileName := w.currentFileName + "_"
	logFiles := getFiles(w.dir, fmt.Sprintf("%s*.%s", fileName, w.fileExtension))
	prefix := filepath.Join(w.dir, fileName)

	// 找到最大索引号的文件
	maxIndex := 0
	for _, file := range logFiles {
		s := file[len(prefix):]                 // 移除文件名称前缀，只要文件索引号部份
		s = s[:len(s)-len("."+w.fileExtension)] // 移除扩展名后缀
		fileIndex := parse.Convert(s, 0)
		if fileIndex > maxIndex {
			maxIndex = fileIndex
		}
	}

	// 如果没有文件，创建索引为1的新文件
	if maxIndex == 0 {
		path := filepath.Join(w.dir, fmt.Sprintf("%s_1.%s", w.currentFileName, w.fileExtension))
		return 1, path
	}

	// 检查最大索引文件的大小
	latestPath := filepath.Join(w.dir, fmt.Sprintf("%s_%d.%s", w.currentFileName, maxIndex, w.fileExtension))
	if info, err := os.Stat(latestPath); err == nil {
		// 如果文件大小未超过限制，复用该文件
		if w.fileSizeLimitMb <= 0 || info.Size() < w.fileSizeLimitMb*1024*1024 {
			return maxIndex, latestPath
		}
	}

	// 文件大小超过限制或无法获取文件信息，创建新文件
	newIndex := maxIndex + 1
	path := filepath.Join(w.dir, fmt.Sprintf("%s_%d.%s", w.currentFileName, newIndex, w.fileExtension))
	return newIndex, path
}

// 优化后的 getFiles，专注于高性能单层目录扫描
func getFiles(path string, searchPattern string) []string {
	var files []string

	// 1. 使用 ReadDir 替代 WalkDir，仅读取当前层级，效率更高
	entries, err := os.ReadDir(path)
	if err != nil {
		return files
	}

	for _, entry := range entries {
		// 2. 过滤掉目录，只处理文件
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()

		// 3. 简化匹配逻辑
		// 直接使用 filepath.Match 匹配文件名，避免复杂的 Join 操作
		if searchPattern != "" {
			if match, _ := filepath.Match(searchPattern, fileName); !match {
				continue
			}
		}

		// 4. 拼接完整路径返回
		files = append(files, filepath.Join(path, fileName))
	}

	return files
}

type fileInfo struct {
	path    string
	modTime int64
}

// 优化后的移除旧文件逻辑
func (w *BatchFileWriter) removeLimitFile() {
	// 1. 获取目录下所有符合后缀的文件
	files := getFiles(w.dir, "*."+w.fileExtension)

	// 如果文件总数未超标，直接返回
	if len(files) <= w.fileCountLimit {
		return
	}

	// 2. 批量获取文件详情（避免在循环中重复调用 Stat）
	var fileList []fileInfo
	for _, f := range files {
		// 排除当前正在写入的文件，不参与清理
		if f == w.currentPath {
			continue
		}

		if info, err := os.Stat(f); err == nil {
			fileList = append(fileList, fileInfo{
				path:    f,
				modTime: info.ModTime().UnixNano(),
			})
		}
	}

	// 3. 按修改时间从旧到新排序 (升序)
	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].modTime < fileList[j].modTime
	})

	// 4. 计算需要删除的数量并批量处理
	// 当前文件必须保留，所以其他文件最多保留 fileCountLimit - 1 个
	overCount := len(fileList) - (w.fileCountLimit - 1)
	if overCount > 0 {
		for i := range overCount {
			fmt.Println("BatchFileWriter: 移除文件", fileList[i].path)
			_ = os.Remove(fileList[i].path)
		}
	}
}
