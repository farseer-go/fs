package flog

import (
	"fmt"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/dateTime"
)

// Trace 打印Trace日志
func Trace(contents ...any) {
	Log(eumLogLevel.Trace, contents...)
}

// Tracef 打印Trace日志
func Tracef(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	Log(eumLogLevel.Trace, content)
}

// Debug 打印Debug日志
func Debug(contents ...any) {
	Log(eumLogLevel.Debug, contents...)
}

// Debugf 打印Debug日志
func Debugf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	Log(eumLogLevel.Debug, content)
}

// Info 打印Info日志
func Info(contents ...any) {
	Log(eumLogLevel.Information, contents...)
}

// Infof 打印Info日志
func Infof(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	Log(eumLogLevel.Information, content)
}

// Warning 打印Warning日志
func Warning(content ...any) {
	Log(eumLogLevel.Warning, content...)
}

// Warningf 打印Warning日志
func Warningf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	Log(eumLogLevel.Warning, content)
}

// Error 打印Error日志
func Error(contents ...any) {
	Log(eumLogLevel.Error, contents...)
}

// Errorf 打印Error日志
func Errorf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	Log(eumLogLevel.Error, content)
}

// Critical 打印Critical日志
func Critical(contents ...any) {
	Log(eumLogLevel.Critical, contents...)
}

// Criticalf 打印Critical日志
func Criticalf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	Log(eumLogLevel.Critical, content)
}

// Log 打印日志
func Log(logLevel eumLogLevel.Enum, contents ...any) {
	content := fmt.Sprint(contents...)
	fmt.Printf("%s %s %s\r\n", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss"), Colors[logLevel]("["+logLevel.ToString()+"]"), content)
}

// Print 打印日志
func Print(contents ...any) {
	content := fmt.Sprint(contents...)
	fmt.Printf("%s %s", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss"), content)
}

// Println 打印日志
func Println(a ...any) {
	content := fmt.Sprintln(a...)
	fmt.Printf("%s %s", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss"), content)
}

// Printf 打印日志
func Printf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss"), content)
}

// ComponentInfo 打印应用日志
func ComponentInfo(appName string, contents ...any) {
	if configure.GetBool("Log.Component." + appName) {
		content := fmt.Sprint(contents...)
		fmt.Printf("%s %s %s\r\n", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss"), Colors[0]("["+appName+"]"), content)
	}
}

// ComponentInfof 打印应用日志
func ComponentInfof(appName string, format string, a ...any) {
	if configure.GetBool("Log.Component." + appName) {
		content := fmt.Sprintf(format, a...)
		fmt.Printf("%s %s %s\r\n", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss"), Colors[0]("["+appName+"]"), content)
	}
}
