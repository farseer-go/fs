package core

import "github.com/farseer-go/fs/core/eumLogLevel"

// ILog 日志记录器的外壳，由ILoggerFactory创建
type ILog interface {
	Trace(contents ...any)                                                         // Trace Trace日志
	Tracef(format string, a ...any)                                                // Tracef Trace日志
	Debug(contents ...any)                                                         // Debug Debug日志
	Debugf(format string, a ...any)                                                // Debugf Debug日志
	Info(contents ...any)                                                          // Info Info日志
	Infof(format string, a ...any)                                                 // Infof Info日志
	Warning(content ...any)                                                        // Warning Warning日志
	Warningf(format string, a ...any)                                              // Warningf Warning日志
	Error(contents ...any) error                                                   // Error Error日志
	Errorf(format string, a ...any) error                                          // Errorf Error日志
	Critical(contents ...any)                                                      // Critical Critical日志
	Criticalf(format string, a ...any)                                             // Criticalf Critical日志
	Log(logLevel eumLogLevel.Enum, content string, component string, newLine bool) // 通用的日志记录
}
