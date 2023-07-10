package flog

import (
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/dateTime"
)

// 日志结构
type logData struct {
	content   string
	createAt  dateTime.DateTime
	logLevel  eumLogLevel.Enum
	component string
}

func newLogData(logLevel eumLogLevel.Enum, content string, component string) *logData {
	return &logData{content: content, createAt: dateTime.Now(), logLevel: logLevel, component: component}
}
