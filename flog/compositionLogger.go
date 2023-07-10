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
func (r *CompositionLogger) log(log *logData) {
	for i := 0; i < len(r.loggerPersistentList); i++ {
		if r.loggerPersistentList[i].IsEnabled(log.LogLevel) {
			r.loggerPersistentList[i].Log(log.LogLevel, log, nil)
		}
	}
}

func (r *CompositionLogger) Trace(contents ...any) {
	log := newLogData(eumLogLevel.Trace, fmt.Sprint(contents...), "")
	r.log(log)
}

func (r *CompositionLogger) Tracef(format string, a ...any) {
	log := newLogData(eumLogLevel.Trace, fmt.Sprintf(format, a...), "")
	r.log(log)
}

func (r *CompositionLogger) Debug(contents ...any) {
	log := newLogData(eumLogLevel.Debug, fmt.Sprint(contents...), "")
	r.log(log)
}

func (r *CompositionLogger) Debugf(format string, a ...any) {
	log := newLogData(eumLogLevel.Debug, fmt.Sprintf(format, a...), "")
	r.log(log)
}

func (r *CompositionLogger) Info(contents ...any) {
	log := newLogData(eumLogLevel.Information, fmt.Sprint(contents...), "")
	r.log(log)
}

func (r *CompositionLogger) Infof(format string, a ...any) {
	log := newLogData(eumLogLevel.Information, fmt.Sprintf(format, a...), "")
	r.log(log)
}

func (r *CompositionLogger) Warning(contents ...any) {
	log := newLogData(eumLogLevel.Warning, fmt.Sprint(contents...), "")
	r.log(log)
}

func (r *CompositionLogger) Warningf(format string, a ...any) {
	log := newLogData(eumLogLevel.Warning, fmt.Sprintf(format, a...), "")
	r.log(log)
}

func (r *CompositionLogger) Error(contents ...any) error {
	log := newLogData(eumLogLevel.Error, fmt.Sprint(contents...), "")
	r.log(log)

	formatter := TextFormatter{}
	return fmt.Errorf(formatter.Formatter(log))
}

func (r *CompositionLogger) Errorf(format string, a ...any) error {
	log := newLogData(eumLogLevel.Error, fmt.Sprintf(format, a...), "")
	r.log(log)

	formatter := TextFormatter{}
	return fmt.Errorf(formatter.Formatter(log))
}

func (r *CompositionLogger) Critical(contents ...any) {
	log := newLogData(eumLogLevel.Critical, fmt.Sprint(contents...), "")
	r.log(log)
}

func (r *CompositionLogger) Criticalf(format string, a ...any) {
	log := newLogData(eumLogLevel.Critical, fmt.Sprintf(format, a...), "")
	r.log(log)
}

func (r *CompositionLogger) Panic(contents ...any) {
	if len(contents) > 0 && contents[0] != nil {
		log := newLogData(eumLogLevel.Error, fmt.Sprint(contents...), "")
		r.log(log)
		panic(fmt.Sprint(contents...))
	}
}

func (r *CompositionLogger) Panicf(format string, a ...any) {
	log := newLogData(eumLogLevel.Error, fmt.Sprintf(format, a...), "")
	r.log(log)
	panic(fmt.Sprintf(format, a...))
}

func (r *CompositionLogger) Print(contents ...any) {
	log := newLogData(eumLogLevel.NoneLevel, fmt.Sprint(contents...), "")
	log.newLine = false
	r.log(log)
}

func (r *CompositionLogger) Println(contents ...any) {
	log := newLogData(eumLogLevel.NoneLevel, fmt.Sprint(contents...), "")
	r.log(log)
}

func (r *CompositionLogger) Printf(format string, a ...any) {
	log := newLogData(eumLogLevel.NoneLevel, fmt.Sprintf(format, a...), "")
	log.newLine = false
	r.log(log)
}

func (r *CompositionLogger) ComponentInfo(appName string, contents ...any) {
	if configure.GetBool("Log.Component." + appName) {
		log := newLogData(eumLogLevel.NoneLevel, fmt.Sprintln(contents...), appName)
		r.log(log)
	}
}

func (r *CompositionLogger) ComponentInfof(appName string, format string, a ...any) {
	if configure.GetBool("Log.Component." + appName) {
		log := newLogData(eumLogLevel.NoneLevel, fmt.Sprintf(format, a...), appName)
		r.log(log)
	}
}
