package flog

import (
	"encoding/json"
	"fmt"
	"github.com/farseer-go/fs/core/eumLogLevel"
)

// IFormatter 日志格式
type IFormatter interface {
	Formatter(log *logData) string
}

// JsonFormatter json格式输出
type JsonFormatter struct {
}

func (r *JsonFormatter) Formatter(log *logData) string {
	marshal, _ := json.Marshal(log)
	return string(marshal)
}

// TextFormatter 文本格式输出
type TextFormatter struct {
}

func (r *TextFormatter) Formatter(log *logData) string {
	var logLevelString string
	if log.logLevel != eumLogLevel.NoneLevel {
		logLevelString = Colors[log.logLevel]("[" + log.logLevel.ToString() + "]")
	} else if log.component != "" {
		logLevelString = Colors[0]("[" + log.component + "]")
	}

	if logLevelString != "" {
		return fmt.Sprintf("%s %s %s", log.createAt.ToString("yyyy-MM-dd hh:mm:ss.ffffff"), logLevelString, log.content)
	}

	return fmt.Sprintf("%s %s", log.createAt.ToString("yyyy-MM-dd hh:mm:ss.ffffff"), log.content)
}
