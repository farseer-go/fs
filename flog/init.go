package flog

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"strings"
)

var logConfig Config

func InitLog() core.ILog {
	factory := DefaultLoggerFactory{}
	logConfig = configure.ParseConfig[Config]("Log")
	// 读取默认等级、默认格式
	defaultLevel := eumLogLevel.GetEnum(logConfig.Default.LogLevel)
	var defaultFormat IFormatter = &TextFormatter{}
	if strings.ToLower(logConfig.Default.Format) == "json" {
		defaultFormat = &JsonFormatter{}
	}

	// 读取控制台配置
	consoleFormat, consoleLevel := loadLevelFormat(logConfig.Console, defaultLevel, defaultFormat)
	factory.AddProviderFormatter(&ConsoleProvider{}, consoleFormat, consoleLevel)

	return factory.CreateLogger("")
}

// 使用具体配置还是默认配置
func loadLevelFormat(logLevelFormat levelFormat, defaultLevel eumLogLevel.Enum, defaultFormat IFormatter) (IFormatter, eumLogLevel.Enum) {
	curLevel := defaultLevel
	curFormat := defaultFormat
	if logLevelFormat.LogLevel != "" {
		curLevel = eumLogLevel.GetEnum(logLevelFormat.LogLevel)
	}

	if logLevelFormat.Format != "" {
		if strings.ToLower(logLevelFormat.Format) == "json" {
			curFormat = &JsonFormatter{}
		} else {
			curFormat = &TextFormatter{}
		}
	}

	return curFormat, curLevel
}

type Config struct {
	Component componentConfig
	Default   levelFormat
	Console   levelFormat
	File      levelFormat
}

// 组件日志
type componentConfig struct {
	Task        bool
	CacheManage bool
}

type levelFormat struct {
	LogLevel string
	Format   string
}
