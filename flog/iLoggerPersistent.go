package flog

import "github.com/farseer-go/fs/core/eumLogLevel"

// ILoggerPersistent 实际的日志记录器
type ILoggerPersistent interface {
	Log(LogLevel eumLogLevel.Enum, log *logData, exception error)
}
