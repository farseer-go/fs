package flog

import (
	"fmt"

	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/snc"
)

// IFormatter 日志格式
type IFormatter interface {
	Formatter(log *LogData) string
}

// JsonFormatter json格式输出
type JsonFormatter struct {
}

func (r JsonFormatter) Formatter(log *LogData) string {
	marshal, _ := snc.Marshal(LogData{
		CreateAt:  log.CreateAt,
		LogLevel:  log.LogLevel,
		Component: log.Component,
		Content:   mustCompile.ReplaceAllString(log.Content, ""),
		newLine:   log.newLine,
	})
	return string(marshal)
}

// TextFormatter 文本格式输出
type TextFormatter struct {
	config levelFormat
}

func (r TextFormatter) Formatter(log *LogData) string {
	var logLevelString string
	if log.LogLevel != eumLogLevel.NoneLevel {
		logLevelString = Colors[log.LogLevel]("[" + log.LogLevel.ToString() + "]")

	} else if log.Component != "" {
		logLevelString = Colors[0]("[" + log.Component + "]")
	}

	if logLevelString != "" {
		return fmt.Sprintf("%s %s %s", log.CreateAt.ToString(r.config.TimeFormat), logLevelString, log.Content)
	}

	return fmt.Sprintf("%s %s", log.CreateAt.ToString(r.config.TimeFormat), log.Content)
}
