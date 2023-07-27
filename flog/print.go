package flog

import (
	"fmt"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
)

// log 日志
var log core.ILog

// Trace 打印Trace日志
func Trace(contents ...any) {
	log.Trace(contents...)
}

// Tracef 打印Trace日志
func Tracef(format string, a ...any) {
	log.Tracef(format, a...)
}

// Debug 打印Debug日志
func Debug(contents ...any) {
	log.Debug(contents...)
}

// Debugf 打印Debug日志
func Debugf(format string, a ...any) {
	log.Debugf(format, a...)
}

// Info 打印Info日志
func Info(contents ...any) {
	log.Info(contents...)
}

// Infof 打印Info日志
func Infof(format string, a ...any) {
	log.Infof(format, a...)
}

// Warning 打印Warning日志
func Warning(contents ...any) {
	log.Warning(contents...)
}

// Warningf 打印Warning日志
func Warningf(format string, a ...any) {
	log.Warningf(format, a...)
}

// Error 打印Error日志
func Error(contents ...any) error {
	return log.Error(contents...)
}

// ErrorIfExists 打印Error日志（假如存在）
func ErrorIfExists(err error) {
	if err == nil {
		return
	}
	_ = log.Error(err)
}

// Errorf 打印Error日志
func Errorf(format string, a ...any) error {
	return log.Errorf(format, a...)
}

// Critical 打印Critical日志
func Critical(contents ...any) {
	log.Critical(contents...)
}

// Criticalf 打印Critical日志
func Criticalf(format string, a ...any) {
	log.Criticalf(format, a...)
}

// Panic 打印Error日志并panic
func Panic(contents ...any) {
	if len(contents) > 0 && contents[0] != nil {
		_ = log.Error(contents...)
		panic(fmt.Sprint(contents...))
	}
}

// Panicf 打印Error日志并panic
func Panicf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	_ = log.Error(content)
	panic(content)
}

// Print 直接记录日志，没有日志等级
func Print(contents ...any) {
	log.Log(eumLogLevel.NoneLevel, fmt.Sprint(contents...), "", false)
}

// Println 直接记录日志，没有日志等级
func Println(contents ...any) {
	log.Log(eumLogLevel.NoneLevel, fmt.Sprint(contents...), "", true)
}

// Printf 直接记录日志，没有日志等级
func Printf(format string, a ...any) {
	log.Log(eumLogLevel.NoneLevel, fmt.Sprintf(format, a...), "", false)
}

// ComponentInfo 组件日志
func ComponentInfo(appName string, contents ...any) {
	log.Log(eumLogLevel.NoneLevel, fmt.Sprint(contents...), "appName", true)
}

// ComponentInfof 组件日志
func ComponentInfof(appName string, format string, a ...any) {
	log.Log(eumLogLevel.NoneLevel, fmt.Sprintf(format, a...), "appName", true)
}
