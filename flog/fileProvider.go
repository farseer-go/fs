package flog

import (
	"bufio"
	"fmt"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/parse"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// FileProvider 控制台打印
type FileProvider struct {
	config fileConfig // 配置
}

func (r *FileProvider) CreateLogger(categoryName string, formatter IFormatter, logLevel eumLogLevel.Enum) ILoggerPersistent {
	// path必须是"/"结尾
	if !strings.HasSuffix(r.config.Path, "/") {
		r.config.Path += "/"
	}

	// 检查目录文件是否存在
	_, err := os.Stat(r.config.Path)
	// 创建目录
	if err != nil {
		_ = os.MkdirAll(r.config.Path, 0766)
	}

	// 刷新文件的时间间隔不能小于1
	if r.config.RefreshInterval < 1 {
		r.config.RefreshInterval = 1
	}

	persistent := &fileLoggerPersistent{formatter, logLevel, make(chan string, 10000), r.config}
	// 异步定时写入日志文件
	go persistent.writeFile()
	// 开启日志文件数量限制
	if r.config.FileCountLimit > 0 {
		go persistent.removeLimitFile()
	}
	return persistent
}

type fileLoggerPersistent struct {
	formatter  IFormatter
	logLevel   eumLogLevel.Enum
	logsBuffer chan string // 写入文件的日志内容
	config     fileConfig  // 配置
}

func (r *fileLoggerPersistent) IsEnabled(logLevel eumLogLevel.Enum) bool {
	return logLevel >= r.logLevel
}

func (r *fileLoggerPersistent) Log(LogLevel eumLogLevel.Enum, log *LogData, exception error) {
	var logContent string
	if log.newLine {
		logContent = r.formatter.Formatter(log) + "\r\n"
	} else {
		logContent = r.formatter.Formatter(log)
	}
	r.logsBuffer <- mustCompile.ReplaceAllString(logContent, "")
}

// 将缓冲区的日志，每隔1秒，写入文件
func (r *fileLoggerPersistent) writeFile() {
	ticker := time.NewTicker(time.Second * time.Duration(r.config.RefreshInterval))
	for range ticker.C {
		// 组装要写入的日志内容
		var logs []string
		for len(r.logsBuffer) > 0 {
			logs = append(logs, <-r.logsBuffer)
		}

		// 没有日志内容
		if len(logs) == 0 {
			continue
		}

		// 根据文件间隔来确定文件名称前缀
		fileName := r.getFilename()

		// 如果开启了日志文件的大小限制，则要拆分文件
		if r.config.FileSizeLimitMb > 0 {
			maxFileIndex := r.getFileIndex(fileName)

			// 拼接最大索引号名称
			fileName += strconv.Itoa(maxFileIndex)
		}

		// 写日志到文件
		file, err := os.OpenFile(r.config.Path+fileName+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			continue
		}
		// 写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)
		_, _ = write.WriteString(strings.Join(logs, ""))
		//Flush将缓存的文件真正写入到文件中
		_ = write.Flush()
		//及时关闭file句柄
		_ = file.Close()
	}
}

// 获取文件名称
func (r *fileLoggerPersistent) getFilename() string {
	// 根据文件间隔来确定文件名称前缀
	var fileName string
	switch strings.ToLower(r.config.RollingInterval) {
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

	// 如果有限制文件大小，则要加后缀
	if r.config.FileSizeLimitMb > 0 {
		fileName += "_"
	}
	return fileName
}

// 获取文件索引号
func (r *fileLoggerPersistent) getFileIndex(fileName string) (maxFileIndex int) {
	// 获取目录下的日志数量，用来确定FileCountLimit限制
	logFiles := getFiles(r.config.Path, fileName+"*.log")
	for _, file := range logFiles {
		if !strings.HasPrefix(file, "./") && !strings.HasPrefix(file, "/") {
			file = "./" + file
		}

		s := file[len(r.config.Path+fileName):] // 移除文件名称前缀，只要文件索引号部份
		s = s[:len(s)-4]                        // 移除.log后缀
		fileIndex := parse.Convert(s, 0)

		// 取最大的索引号
		if fileIndex > maxFileIndex {
			maxFileIndex = fileIndex
		}
	}

	// 获取最大索引号的文件大小
	fileInfo, _ := os.Stat(r.config.Path + fileName + strconv.Itoa(maxFileIndex) + ".log")
	// 如果文件超出大小限制
	if fileInfo != nil && fileInfo.Size()/1024/1024 >= r.config.FileSizeLimitMb {
		maxFileIndex++ // 增加索引号
	}
	return
}

// 移除超过数量的日志文件
func (r *fileLoggerPersistent) removeLimitFile() {
	// 1分钟判断一次
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		// 获取目录下的日志数量，用来确定FileCountLimit限制
		logFiles := getFiles(r.config.Path, "*.log")
		// 文件没有超出限制
		if len(logFiles) < r.config.FileCountLimit {
			continue
		}

		var oldFile string                       // 最旧的文件路径
		oldFileModTime := time.Now().UnixMilli() // 最旧的文件修改时间（用来比较哪个是最旧的）
		// 找出最早的文件，并删除
		for _, file := range logFiles {
			if !strings.HasPrefix(file, "./") && !strings.HasPrefix(file, "/") {
				file = "./" + file
			}
			fileInfo, _ := os.Stat(file)
			// 找到最早修改的文件
			if oldFileModTime > fileInfo.ModTime().UnixMilli() {
				oldFileModTime = fileInfo.ModTime().UnixMilli()
				oldFile = file
			}
		}

		_ = os.Remove(oldFile)
	}
}

// GetFiles 读取指定目录下的文件
// path：目录路径
// searchPattern：匹配文件名要包含的名称，搜索全部，传入""即可
// searchSubDir：是否要搜索子目录
func getFiles(path string, searchPattern string) []string {
	var files []string
	_ = filepath.WalkDir(path, func(filePath string, dirInfo fs.DirEntry, err error) error {
		if path == filePath {
			return nil
		}
		// 目录不需要判断，filepath.Walk执行就包含递归了
		if !dirInfo.IsDir() {
			match := true
			if searchPattern != "" {
				match, _ = filepath.Match(filepath.Join(filepath.Dir(filePath), searchPattern), filePath)
			}
			if match {
				files = append(files, filePath)
			}
		} else if dirInfo.IsDir() {
			return fs.SkipDir
		}
		return nil
	})
	return files
}
