package flog

import (
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
)

type ILoggerFactory interface {
	// AddProvider 添加多个日志写入提供者（应用初始化时调用）
	AddProvider(provider ILoggerProvider)
	// CreateLogger 根据已添加的Provider，创建组合模式的Logger（壳）
	CreateLogger(categoryName string) core.ILog
}

// DefaultLoggerFactory 默认的日志创建工厂（不需要应用去实现）
type DefaultLoggerFactory struct {
	// 已添加的日志记录器
	loggerPersistentList []ILoggerPersistent
}

func (r *DefaultLoggerFactory) AddProvider(provider ILoggerProvider) {
	r.AddProviderFormatter(provider, &TextFormatter{}, eumLogLevel.Information)
}

func (r *DefaultLoggerFactory) AddProviderFormatter(provider ILoggerProvider, formatter IFormatter, logLevel eumLogLevel.Enum) {
	if provider != nil {
		r.loggerPersistentList = append(r.loggerPersistentList, provider.CreateLogger("", formatter, logLevel))
	}
}

func (r *DefaultLoggerFactory) CreateLogger(categoryName string) core.ILog {
	return &CompositionLogger{
		loggerPersistentList: r.loggerPersistentList,
	}
}
