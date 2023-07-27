package flog

import (
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
)

// LogBuffer 日志缓冲区
var LogBuffer = make(chan string, 1000)

// LoadLogBuffer 从日志缓冲区读取日志并打印
func LoadLogBuffer(logIns core.ILog) {
	for log := range LogBuffer {
		logIns.Log(eumLogLevel.NoneLevel, log, "", true)
		//Println(log)
	}
}

// ClearLogBuffer 清空缓冲区的日志
func ClearLogBuffer(logIns core.ILog) {
	for len(LogBuffer) > 0 {
		logIns.Log(eumLogLevel.NoneLevel, <-LogBuffer, "", true)
		//Println(<-LogBuffer)
	}
}
