package flog

import (
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/snowflake"
	"github.com/farseer-go/fs/trace"
	"regexp"
)

// var regexStr = "\\\\u001b\\[[\\d;]*m"
var regexStr = "\u001b\\[[\\d;]*m"
var mustCompile = regexp.MustCompile(regexStr)

// LogData 日志结构
type LogData struct {
	CreateAt  dateTime.DateTime
	LogLevel  eumLogLevel.Enum
	Component string // 组件名称
	Content   string
	newLine   bool // 是否需要换行
	// 上传到FOPS时使用
	TraceId string // 上下文ID
	AppId   string // 应用ID
	AppName string // 应用名称
	AppIp   string // 应用IP
	LogId   string // 主键ID
}

func newLogData(logLevel eumLogLevel.Enum, content string, component string) *LogData {
	var traceId string
	if t := trace.CurTraceContext.Get(); t != nil {
		traceId = t.GetTraceId()
	}
	return &LogData{Content: content, CreateAt: dateTime.Now(), LogLevel: logLevel, Component: component, newLine: true, TraceId: traceId, AppId: parse.ToString(core.AppId), AppName: core.AppName, AppIp: core.AppIp, LogId: parse.ToString(snowflake.GenerateId())}
}

//// 清除颜色
//func (receiver *LogData) clearColor() {
//	receiver.Content = mustCompile.ReplaceAllString(receiver.Content, "")
//}
