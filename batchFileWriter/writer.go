package batchFileWriter

import (
	"bufio"
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
)

type BatchFileWriter struct {
	dataChan         chan []byte   // 明确只接收处理好的字节流
	dir              string        // 存储目录
	fileExtension    string        // 文件扩展名
	rollingInterval  string        // 文件滚动间隔（year/month/day/week/hour）
	fileSizeLimitMb  int64         // 文件大小限制（MB），设置之后，只有当文件大小超过限制时才会滚动
	fileCountLimit   int           // 文件数量限制
	interval         time.Duration // 多长时间刷盘一次
	currentFileIndex int           // 缓存当前的索引号
	currentPath      string        // current文件的完整路径
	bufferSize       int
	exitChan         chan struct{}
	wg               sync.WaitGroup
	closeOnce        sync.Once // 确保 Close 方法只执行一次
	// ... 其他字段
	closed atomic.Bool
}

// dir: 存储目录
// fileExtension: 文件扩展名
// rollingInterval: 文件滚动间隔（year/month/day/week/hour）
// fileSizeLimitMb: 限制文件大小
// fileCountLimit: 限制文件数量
// interval: 多长时间生成一个文件。PS: 如果开启了文件大小限制，则不生效
func NewWriter(dir string, fileExtension string, rollingInterval string, fileSizeLimitMb int64, fileCountLimit int, interval time.Duration) *BatchFileWriter {
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
		exitChan:        make(chan struct{}),
		currentPath:     filepath.Join(dir, fmt.Sprintf("%s.%s", "current", fileExtension)),
	}
	// 初始化：获取文件索引和当前大小
	w.currentFileIndex = w.getFileIndex()

	// 开启翻转协程
	w.wg.Add(1)
	go w.daemon()
	return w
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
		payload, err = snc.Marshal(v)
		if err != nil {
			fmt.Println("BatchFileWriter转成json时失败: ", err.Error())
		}
	}

	// 2. 将处理好的字节流放入异步队列
	if len(payload) > 0 && !w.closed.Load() {
		select {
		case w.dataChan <- payload:
		case <-w.exitChan:
			fmt.Println("BatchFileWriter.Write: 已关闭，丢弃数据")
		default:
			fmt.Println("BatchFileWriter.Write: 由于队列已满，丢弃数据")
		}
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
			bufWriter.Write(line)
			bufWriter.WriteByte('\n')
			fileSize += int64(len(line) + 1)

			// 如果文件大小限制开启，且当前文件大小超过限制，则进行翻转
			if w.fileSizeLimitMb > 0 && fileSize >= w.fileSizeLimitMb*1024*1024 {
				f, fileSize = w.rotate(bufWriter, f)
			}
		case <-ticker.C:
			// 如果文件大小限制没有开启，则按照时间间隔翻转
			if w.fileSizeLimitMb <= 0 && w.interval > 0 { // 判断w.interval > 0,是因为有可能没有设置时间间隔
				f, fileSize = w.rotate(bufWriter, f)
			}
		case <-w.exitChan:
			// 退出前排空 Channel
			close(w.dataChan)
			for line := range w.dataChan {
				bufWriter.Write(line)
				bufWriter.WriteByte('\n')
				fileSize += int64(len(line) + 1)
			}
			// 如果文件大小限制开启，且当前文件大小超过限制，则进行翻转
			if w.fileSizeLimitMb > 0 && fileSize >= w.fileSizeLimitMb*1024*1024 {
				f, fileSize = w.rotate(bufWriter, f)
			} else {
				bufWriter.Flush()
				_ = f.Sync()
				_ = f.Close()

				// 如果开启了文件数量限制，则在翻转时检查是否需要删除旧文件
				if w.fileCountLimit > 0 {
					w.removeLimitFile()
				}
			}
			return
		}
	}
}

func (w *BatchFileWriter) rotate(bufWriter *bufio.Writer, f *os.File) (*os.File, int64) {
	bufWriter.Flush()
	_ = f.Sync()
	_ = f.Close()

	if info, err := os.Stat(w.currentPath); err == nil && info.Size() > 0 {
		dispatchName := fmt.Sprintf("%s_%d.%s", w.getFilename(), w.currentFileIndex, w.fileExtension)
		_ = os.Rename(w.currentPath, filepath.Join(w.dir, dispatchName))
		w.currentFileIndex++
	}

	// 支持重试打开文件，避免因磁盘暂时不可用导致的翻转失败
	var err error
	for {
		if f, _, err = w.openFile(w.currentPath); err == nil {
			break
		}
		fmt.Println("BatchFileWriter.rotate: 打开文件失败，3秒后重试: " + err.Error())
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

// 获取文件索引号
func (w *BatchFileWriter) getFileIndex() (maxFileIndex int) {
	// 文件名和索引之间有个间隔符号
	fileName := w.getFilename() + "_"
	// 获取目录下的日志数量，用来确定FileCountLimit限制
	logFiles := getFiles(w.dir, fmt.Sprintf("%s*.%s", fileName, w.fileExtension))
	for _, file := range logFiles {
		if !strings.HasPrefix(file, "./") && !strings.HasPrefix(file, "/") {
			file = "./" + file
		}

		s := file[len(filepath.Join(w.dir, fileName)):] // 移除文件名称前缀，只要文件索引号部份
		s = s[:len(s)-len("."+w.fileExtension)]         // 移除扩展名后缀
		fileIndex := parse.Convert(s, 0)

		// 取最大的索引号
		if fileIndex > maxFileIndex {
			maxFileIndex = fileIndex
		}
	}

	// 增加索引号
	maxFileIndex++
	return
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

	// 如果文件总数未超标，直接返回（注意排除 current 文件）
	// 这里减 1 是因为 getFiles 可能会扫到正在写的 current 文件
	if len(files) <= w.fileCountLimit {
		return
	}

	// 2. 批量获取文件详情（避免在循环中重复调用 Stat）
	var fileList []fileInfo
	for _, f := range files {
		// 排除当前正在写入的文件，不参与清理
		if strings.HasSuffix(f, "current."+w.fileExtension) {
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
	// 剩余需要保留的文件数量 = w.fileCountLimit
	overCount := len(fileList) - w.fileCountLimit
	if overCount > 0 {
		for i := 0; i < overCount; i++ {
			_ = os.Remove(fileList[i].path)
		}
	}
}
