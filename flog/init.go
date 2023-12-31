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
	// 默认为yyyy-MM-dd hh:mm:ss.ffffff格式
	if logConfig.Default.TimeFormat == "" {
		logConfig.Default.TimeFormat = "yyyy-MM-dd hh:mm:ss.ffffff"
	}

	// 读取默认等级、默认格式
	defaultLevel := eumLogLevel.GetEnum(logConfig.Default.LogLevel)
	// 如果配置文件没有设置日志级别，则默认为Trace级别
	if logConfig.Default.LogLevel == "" {
		defaultLevel = eumLogLevel.Trace
	}
	var defaultFormat IFormatter = &TextFormatter{config: logConfig.Default}
	if strings.ToLower(logConfig.Default.Format) == "json" {
		defaultFormat = &JsonFormatter{}
	}

	// 读取控制台打印配置
	if !logConfig.Console.Disable {
		formatter, logLevel := loadLevelFormat(logConfig.Console, defaultLevel, defaultFormat, logConfig.Default)
		factory.AddProviderFormatter(&ConsoleProvider{}, formatter, logLevel)
	}

	// 读取文件写入配置
	if !logConfig.File.Disable && logConfig.File.Path != "" {
		if logConfig.File.Path == "" {
			logConfig.File.Path = "./log"
		}
		formatter, logLevel := loadLevelFormat(logConfig.File.levelFormat, defaultLevel, defaultFormat, logConfig.Default)
		factory.AddProviderFormatter(&FileProvider{config: logConfig.File}, formatter, logLevel)
	}

	// 上传到FOPS
	if !logConfig.Fops.Disable && configure.GetString("Fops.Server") != "" {
		formatter, logLevel := loadLevelFormat(logConfig.Console, defaultLevel, defaultFormat, logConfig.Default)
		factory.AddProviderFormatter(&FopsProvider{}, formatter, logLevel)
	}
	log = factory.CreateLogger("")
	return log
}

// 使用具体配置还是默认配置
func loadLevelFormat(logLevelFormat levelFormat, defaultLevel eumLogLevel.Enum, defaultFormatter IFormatter, defaultLevelFormat levelFormat) (IFormatter, eumLogLevel.Enum) {
	curLevel := defaultLevel
	curFormat := defaultFormatter
	// 使用默认的时间格式
	if logLevelFormat.TimeFormat == "" {
		logLevelFormat.TimeFormat = defaultLevelFormat.TimeFormat
	}

	// 默认日志等级
	if logLevelFormat.LogLevel != "" {
		curLevel = eumLogLevel.GetEnum(logLevelFormat.LogLevel)
	}

	// 默认日志内容格式
	if logLevelFormat.Format != "" {
		if strings.ToLower(logLevelFormat.Format) == "json" {
			curFormat = &JsonFormatter{}
		} else {
			curFormat = &TextFormatter{config: logLevelFormat}
		}
	}

	return curFormat, curLevel
}
