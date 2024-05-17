package flog

import (
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"time"
)

// LogBuffer 日志缓冲区
var LogBuffer = make(chan string, 1000)

// LoadLogBuffer 从日志缓冲区读取日志并打印
func LoadLogBuffer(logIns core.ILog) {
	for log := range LogBuffer {
		logIns.Log(eumLogLevel.NoneLevel, log, "", true)
	}
}

// ClearLogBuffer 清空缓冲区的日志
func ClearLogBuffer(logIns core.ILog) {
	for len(LogBuffer) > 0 {
		logIns.Log(eumLogLevel.NoneLevel, <-LogBuffer, "", true)
	}
}

func CloseBuffer() {
	close(LogBuffer)
	time.Sleep(200 * time.Microsecond)
}
