package core

import (
	"github.com/farseer-go/fs/dateTime"
)

var (
	StartupAt dateTime.DateTime // 应用启动时间
	AppName   string            // 应用名称
	HostName  string            // 主机名称
	AppId     int64             // 应用ID
	AppIp     string            // 应用IP
	ProcessId int               // 进程Id
)
