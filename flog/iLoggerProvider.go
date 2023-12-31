package flog

import "github.com/farseer-go/fs/core/eumLogLevel"

type ILoggerProvider interface {
	// CreateLogger 创建对应的Logger对象，创建Logger
	CreateLogger(categoryName string, formatter IFormatter, logLevel eumLogLevel.Enum) ILoggerPersistent
}

// ILoggerPersistent 实际的日志记录器
type ILoggerPersistent interface {
	// IsEnabled 根据日志等级确定是否需要记录
	IsEnabled(logLevel eumLogLevel.Enum) bool
	// Log 日志记录
	Log(LogLevel eumLogLevel.Enum, log *LogData, exception error)
}
