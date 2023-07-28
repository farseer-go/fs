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
	// 如果配置文件没有设置日志级别，则默认为Trace级别
	if logConfig.Default.LogLevel == "" {
		defaultLevel = eumLogLevel.Trace
	}
	var defaultFormat IFormatter = &TextFormatter{}
	if strings.ToLower(logConfig.Default.Format) == "json" {
		defaultFormat = &JsonFormatter{}
	}

	// 读取控制台打印配置
	if !logConfig.Console.Disable {
		formatter, logLevel := loadLevelFormat(logConfig.Console, defaultLevel, defaultFormat)
		factory.AddProviderFormatter(&ConsoleProvider{}, formatter, logLevel)
	}

	// 读取文件写入配置
	if !logConfig.File.Disable || logConfig.File.Path != "" {
		if logConfig.File.Path == "" {
			logConfig.File.Path = "./log"
		}
		formatter, logLevel := loadLevelFormat(logConfig.File.levelFormat, defaultLevel, defaultFormat)
		factory.AddProviderFormatter(&FileProvider{config: logConfig.File}, formatter, logLevel)
	}

	log = factory.CreateLogger("")
	return log
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
	File      fileConfig
}

// 组件日志
type componentConfig struct {
	Task        bool
	CacheManage bool
}

type levelFormat struct {
	LogLevel string // 只记录的等级
	Format   string // 记录格式
	Disable  bool   // 停用
}

type fileConfig struct {
	levelFormat
	Path            string // 日志存放的目录位置
	FileName        string // 日志文件名称
	RollingInterval string // 日志滚动间隔
	FileSizeLimitMb int64  // 日志文件大小限制（MB）
	FileCountLimit  int    // 日志文件数量限制
	RefreshInterval int    // 写入到文件的时间间隔，最少为1（秒）
}
