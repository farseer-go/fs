package flog

import "github.com/farseer-go/fs/core/eumLogLevel"

type ILoggerProvider interface {
	// CreateLogger 创建对应的Logger对象，创建Logger
	CreateLogger(categoryName string, formatter IFormatter, logLevel eumLogLevel.Enum) ILoggerPersistent
}
