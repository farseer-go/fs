package flog

import (
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/dateTime"
)

// 日志结构
type logData struct {
	CreateAt  dateTime.DateTime
	LogLevel  eumLogLevel.Enum
	Component string // 组件名称
	Content   string
	newLine   bool // 是否需要换行
}

func newLogData(logLevel eumLogLevel.Enum, content string, component string) *logData {
	return &logData{Content: content, CreateAt: dateTime.Now(), LogLevel: logLevel, Component: component, newLine: true}
}
