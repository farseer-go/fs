package core

// ILog 日志接口
type ILog interface {
	Trace(contents ...any)                // Trace 打印Trace日志
	Tracef(format string, a ...any)       // Tracef 打印Trace日志
	Debug(contents ...any)                // Debug 打印Debug日志
	Debugf(format string, a ...any)       // Debugf 打印Debug日志
	Info(contents ...any)                 // Info 打印Info日志
	Infof(format string, a ...any)        // Infof 打印Info日志
	Warning(content ...any)               // Warning 打印Warning日志
	Warningf(format string, a ...any)     // Warningf 打印Warning日志
	Error(contents ...any) error          // Error 打印Error日志
	Errorf(format string, a ...any) error // Errorf 打印Error日志
	Panic(contents ...any)                // Panic 打印Error日志并panic
	Panicf(format string, a ...any)       // Panicf 打印Error日志并panic
	Critical(contents ...any)             // Critical 打印Critical日志
	Criticalf(format string, a ...any)    // Criticalf 打印Critical日志
}
