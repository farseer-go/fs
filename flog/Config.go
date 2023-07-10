package flog

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/core"
)

var logConfig Config

func InitLog() core.ILog {
	logConfig = configure.ParseConfig[Config]("Log")

	factory := DefaultLoggerFactory{}
	factory.AddProvider(&ConsoleProvider{})
	return factory.CreateLogger("")
}

type Config struct {
	LogLevel  string
	Component componentConfig
}

type componentConfig struct {
	Task        bool
	CacheManage bool
}
