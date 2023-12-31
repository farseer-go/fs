package flog

import (
	"fmt"
	"github.com/farseer-go/fs/core/eumLogLevel"
)

// ConsoleProvider 控制台打印
type ConsoleProvider struct {
}

func (r *ConsoleProvider) CreateLogger(categoryName string, formatter IFormatter, logLevel eumLogLevel.Enum) ILoggerPersistent {
	return &consoleLoggerPersistent{formatter: formatter, logLevel: logLevel}
}

type consoleLoggerPersistent struct {
	formatter IFormatter
	logLevel  eumLogLevel.Enum
}

func (r *consoleLoggerPersistent) IsEnabled(logLevel eumLogLevel.Enum) bool {
	return logLevel >= r.logLevel
}

func (r *consoleLoggerPersistent) Log(LogLevel eumLogLevel.Enum, log *LogData, exception error) {
	if log.newLine {
		fmt.Println(r.formatter.Formatter(log))
	} else {
		fmt.Print(r.formatter.Formatter(log))
	}
}
