package flog

import "github.com/farseer-go/fs/configure"

var LogConfig Config

func Init() {
	LogConfig = configure.ParseConfig[Config]("Log")
}

type Config struct {
	LogLevel  string
	Component componentConfig
}

type componentConfig struct {
	HttpInvoke  bool
	Task        bool
	CacheManage bool
}
