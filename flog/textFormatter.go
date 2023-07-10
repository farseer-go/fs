package flog

import (
	"fmt"
	"github.com/farseer-go/fs/core/eumLogLevel"
)

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
