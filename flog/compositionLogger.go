package flog

import (
	"fmt"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/core/eumLogLevel"
)

// CompositionLogger 根据已添加的Provider，创建组合模式的Logger（壳）
type CompositionLogger struct {
	loggerPersistentList []ILoggerPersistent
}

// 调用所有日志记录器的实现
func (r *CompositionLogger) log(log *LogData) {
	for i := 0; i < len(r.loggerPersistentList); i++ {
		if r.loggerPersistentList[i].IsEnabled(log.LogLevel) {
			r.loggerPersistentList[i].Log(log.LogLevel, log, nil)
		}
	}
}

func (r *CompositionLogger) Trace(contents ...any) {
	r.log(newLogData(eumLogLevel.Trace, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Tracef(format string, a ...any) {
	r.Trace(fmt.Sprintf(format, a...))
}

func (r *CompositionLogger) Debug(contents ...any) {
	r.log(newLogData(eumLogLevel.Debug, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Debugf(format string, a ...any) {
	r.Debug(fmt.Sprintf(format, a...))
}

func (r *CompositionLogger) Info(contents ...any) {
	r.log(newLogData(eumLogLevel.Information, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Infof(format string, a ...any) {
	r.Info(fmt.Sprintf(format, a...))
}

func (r *CompositionLogger) Warning(contents ...any) {
	r.log(newLogData(eumLogLevel.Warning, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Warningf(format string, a ...any) {
	r.Warning(fmt.Sprintf(format, a...))
}

func (r *CompositionLogger) Error(contents ...any) error {
	log := newLogData(eumLogLevel.Error, fmt.Sprint(contents...), "")
	r.log(log)
	return fmt.Errorf(TextFormatter{}.Formatter(log))
}

func (r *CompositionLogger) Errorf(format string, a ...any) error {
	return r.Error(fmt.Sprintf(format, a...))
}

func (r *CompositionLogger) Critical(contents ...any) {
	r.log(newLogData(eumLogLevel.Critical, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Criticalf(format string, a ...any) {
	r.Criticalf(fmt.Sprintf(format, a...))
}

func (r *CompositionLogger) Log(logLevel eumLogLevel.Enum, content string, component string, newLine bool) {
	// 日志不需要记录
	if component != "" && !configure.GetBool("Log.Component."+component) {
		return
	}

	log := newLogData(logLevel, content, component)
	log.newLine = newLine
	r.log(log)
}
