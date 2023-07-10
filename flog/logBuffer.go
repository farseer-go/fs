package flog

import (
	"github.com/farseer-go/fs/core"
)

// LogBuffer 日志缓冲区
var LogBuffer = make(chan string, 1000)

// LoadLogBuffer 从日志缓冲区读取日志并打印
func LoadLogBuffer(Log core.ILog) {
	for log := range LogBuffer {
		Log.Println(log)
	}
}

// ClearLogBuffer 清空缓冲区的日志
func ClearLogBuffer(Log core.ILog) {
	for len(LogBuffer) > 0 {
		Log.Println(<-LogBuffer)
	}
}
