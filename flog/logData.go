package flog

import (
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/dateTime"
	"regexp"
)

// var regexStr = "\\\\u001b\\[[\\d;]*m"
var regexStr = "\u001b\\[[\\d;]*m"
var mustCompile = regexp.MustCompile(regexStr)

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

//// 清除颜色
//func (receiver *logData) clearColor() {
//	receiver.Content = mustCompile.ReplaceAllString(receiver.Content, "")
//}
