package flog

import (
	"fmt"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"time"
)

// Trace 打印Trace日志
func Trace(content any) {
	Log(eumLogLevel.Trace, content)
}

// Tracef 打印Trace日志
func Tracef(format string, a ...any) {
	content := fmt.Sprintf(format, a)
	Log(eumLogLevel.Trace, content)
}

// Debug 打印Debug日志
func Debug(content any) {
	Log(eumLogLevel.Debug, content)
}

// Debugf 打印Debug日志
func Debugf(format string, a ...any) {
	content := fmt.Sprintf(format, a)
	Log(eumLogLevel.Debug, content)
}

// Info 打印Info日志
func Info(content any) {
	Log(eumLogLevel.Information, content)
}

// Infof 打印Info日志
func Infof(format string, a ...any) {
	content := fmt.Sprintf(format, a)
	Log(eumLogLevel.Information, content)
}

// Warning 打印Warning日志
func Warning(content any) {
	Log(eumLogLevel.Warning, content)
}

// Warningf 打印Warning日志
func Warningf(format string, a ...any) {
	content := fmt.Sprintf(format, a)
	Log(eumLogLevel.Warning, content)
}

// Error 打印Error日志
func Error(content any) {
	Log(eumLogLevel.Error, content)
}

// Errorf 打印Error日志
func Errorf(format string, a ...any) {
	content := fmt.Sprintf(format, a)
	Log(eumLogLevel.Error, content)
}

// Critical 打印Critical日志
func Critical(content any) {
	Log(eumLogLevel.Critical, content)
}

// Criticalf 打印Critical日志
func Criticalf(format string, a ...any) {
	content := fmt.Sprintf(format, a)
	Log(eumLogLevel.Critical, content)
}

// Log 打印日志
func Log(logLevel eumLogLevel.Enum, content any) {
	fmt.Printf("%s [%s] %s\r\n", time.Now().Format("2006-01-02 15:04:05"), logLevel.ToString(), content)
}

// Print 打印日志
func Print(content any) {
	fmt.Printf("%s %s", time.Now().Format("2006-01-02 15:04:05"), content)
}

// Println 打印日志
func Println(a ...any) {
	content := fmt.Sprintln(a...)
	fmt.Printf("%s %s", time.Now().Format("2006-01-02 15:04:05"), content)
}

// Printf 打印日志
func Printf(format string, a ...any) {
	content := fmt.Sprintf(format, a)
	fmt.Printf("%s %s", time.Now().Format("2006-01-02 15:04:05"), content)
}
