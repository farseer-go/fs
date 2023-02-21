package flog

import "github.com/farseer-go/fs/configure"

var logConfig Config

func Init() {
	logConfig = configure.ParseConfig[Config]("Log")
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
