package flog

import (
	"time"

	"github.com/farseer-go/fs/batchFileWriter"
	"github.com/farseer-go/fs/core/eumLogLevel"
)

// FileProvider 控制台打印
type FileProvider struct {
	config fileConfig // 配置
}

func (r *FileProvider) CreateLogger(categoryName string, formatter IFormatter, logLevel eumLogLevel.Enum) ILoggerPersistent {
	// 刷新文件的时间间隔不能小于1
	if r.config.RefreshInterval < 1 || r.config.FileSizeLimitMb <= 0 {
		r.config.RefreshInterval = 1
	}

	writer := batchFileWriter.NewWriter(r.config.Path, "log", r.config.RollingInterval, r.config.FileSizeLimitMb, r.config.FileCountLimit, time.Second*time.Duration(r.config.RefreshInterval), false)
	persistent := &fileLoggerPersistent{formatter, logLevel, r.config, writer}
	return persistent
}

type fileLoggerPersistent struct {
	formatter  IFormatter
	logLevel   eumLogLevel.Enum
	config     fileConfig // 配置
	FileWriter *batchFileWriter.BatchFileWriter
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
	r.FileWriter.Write(mustCompile.ReplaceAllString(logContent, ""))
}

func (r *fileLoggerPersistent) Close() {
	r.FileWriter.Close()
}
