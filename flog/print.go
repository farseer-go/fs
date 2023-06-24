package flog

import (
	"fmt"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/dateTime"
	"runtime"
	"strconv"
	"strings"
)

// Trace 打印Trace日志
func Trace(contents ...any) {
	log(eumLogLevel.Trace, contents...)
}

// Tracef 打印Trace日志
func Tracef(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	log(eumLogLevel.Trace, content)
}

// Debug 打印Debug日志
func Debug(contents ...any) {
	log(eumLogLevel.Debug, contents...)
}

// Debugf 打印Debug日志
func Debugf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	log(eumLogLevel.Debug, content)
}

// Info 打印Info日志
func Info(contents ...any) {
	log(eumLogLevel.Information, contents...)
}

// Infof 打印Info日志
func Infof(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	log(eumLogLevel.Information, content)
}

// Warning 打印Warning日志
func Warning(content ...any) {
	log(eumLogLevel.Warning, content...)
}

// Warningf 打印Warning日志
func Warningf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	log(eumLogLevel.Warning, content)
}

// Error 打印Error日志
func Error(contents ...any) error {
	Println(Blue(fileWithLineNum()))
	return fmt.Errorf(log(eumLogLevel.Error, contents...))
}

// Errorf 打印Error日志
func Errorf(format string, a ...any) error {
	Println(Blue(fileWithLineNum()))
	content := fmt.Sprintf(format, a...)
	return fmt.Errorf(log(eumLogLevel.Error, content))
}

// Panic 打印Error日志并panic
func Panic(contents ...any) {
	if len(contents) > 0 && contents[0] != nil {
		log(eumLogLevel.Error, contents...)
		panic(fmt.Sprint(contents...))
	}
}

// Panicf 打印Error日志并panic
func Panicf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	log(eumLogLevel.Error, content)

	panic(content)
}

// Critical 打印Critical日志
func Critical(contents ...any) {
	log(eumLogLevel.Critical, contents...)
}

// Criticalf 打印Critical日志
func Criticalf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	log(eumLogLevel.Critical, content)
}

// log 打印日志
func log(logLevel eumLogLevel.Enum, contents ...any) string {
	content := fmt.Sprint(contents...)
	msg := fmt.Sprintf("%s %s %s\r\n", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss.ffffff"), Colors[logLevel]("["+logLevel.ToString()+"]"), content)

	switch strings.ToLower(logConfig.LogLevel) {
	case "debug":
		if logLevel < 1 {
			return msg
		}
	case "information", "info":
		if logLevel < 2 {
			return msg
		}
	case "warning", "warn":
		if logLevel < 3 {
			return msg
		}
	case "error":
		if logLevel < 4 {
			return msg
		}
	case "critical":
		if logLevel < 5 {
			return msg
		}
	}

	fmt.Print(msg)
	return msg
}

// Print 打印日志
func Print(contents ...any) {
	content := fmt.Sprint(contents...)
	fmt.Printf("%s %s", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss.ffffff"), content)
}

// Println 打印日志
func Println(a ...any) {
	content := fmt.Sprintln(a...)
	fmt.Printf("%s %s", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss.ffffff"), content)
}

// Printf 打印日志
func Printf(format string, a ...any) {
	content := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss.ffffff"), content)
}

// ComponentInfo 打印应用日志
func ComponentInfo(appName string, contents ...any) {
	if configure.GetBool("Log.Component." + appName) {
		content := fmt.Sprint(contents...)
		fmt.Printf("%s %s %s\r\n", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss.ffffff"), Colors[0]("["+appName+"]"), content)
	}
}

// ComponentInfof 打印应用日志
func ComponentInfof(appName string, format string, a ...any) {
	if configure.GetBool("Log.Component." + appName) {
		content := fmt.Sprintf(format, a...)
		fmt.Printf("%s %s %s\r\n", dateTime.Now().ToString("yyyy-MM-dd hh:mm:ss.ffffff"), Colors[0]("["+appName+"]"), content)
	}
}

func fileWithLineNum() string {
	// the second caller usually from internal, so set i start from 1
	var fileLineNum string
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.HasSuffix(file, "_test.go")) { // !strings.HasPrefix(file, gormSourceDir) ||
			fileLineNum = file + ":" + strconv.FormatInt(int64(line), 10)
			break
		}
	}
	return fileLineNum
}
