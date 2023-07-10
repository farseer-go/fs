package flog

import "github.com/farseer-go/fs/core/eumLogLevel"

type ILoggerProvider interface {
	// CreateLogger 创建对应的Logger对象，创建Logger
	CreateLogger(categoryName string, formatter IFormatter, logLevel eumLogLevel.Enum) ILoggerPersistent
}

// ILoggerPersistent 实际的日志记录器
type ILoggerPersistent interface {
	Log(LogLevel eumLogLevel.Enum, log *logData, exception error)
}
