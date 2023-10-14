package flog

import (
	"fmt"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"runtime"
	"strconv"
	"strings"
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
	r.log(newLogData(eumLogLevel.Trace, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Tracef(format string, a ...any) {
	r.log(newLogData(eumLogLevel.Trace, fmt.Sprintf(format, a...), ""))
}

func (r *CompositionLogger) Debug(contents ...any) {
	r.log(newLogData(eumLogLevel.Debug, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Debugf(format string, a ...any) {
	r.log(newLogData(eumLogLevel.Debug, fmt.Sprintf(format, a...), ""))
}

func (r *CompositionLogger) Info(contents ...any) {
	r.log(newLogData(eumLogLevel.Information, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Infof(format string, a ...any) {
	r.log(newLogData(eumLogLevel.Information, fmt.Sprintf(format, a...), ""))
}

func (r *CompositionLogger) Warning(contents ...any) {
	r.log(newLogData(eumLogLevel.Warning, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Warningf(format string, a ...any) {
	r.log(newLogData(eumLogLevel.Warning, fmt.Sprintf(format, a...), ""))
}

func (r *CompositionLogger) Error(contents ...any) error {
	r.log(newLogData(eumLogLevel.NoneLevel, Blue(r.fileWithLineNum()), ""))

	log := newLogData(eumLogLevel.Error, fmt.Sprint(contents...), "")
	r.log(log)
	return fmt.Errorf(TextFormatter{}.Formatter(log))
}

func (r *CompositionLogger) Errorf(format string, a ...any) error {
	r.log(newLogData(eumLogLevel.NoneLevel, Blue(r.fileWithLineNum()), ""))

	log := newLogData(eumLogLevel.Error, fmt.Sprintf(format, a...), "")
	r.log(log)
	return fmt.Errorf(TextFormatter{}.Formatter(log))
}

func (r *CompositionLogger) Critical(contents ...any) {
	r.log(newLogData(eumLogLevel.Critical, fmt.Sprint(contents...), ""))
}

func (r *CompositionLogger) Criticalf(format string, a ...any) {
	r.log(newLogData(eumLogLevel.Critical, fmt.Sprintf(format, a...), ""))
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

func (r *CompositionLogger) fileWithLineNum() string {
	// the second caller usually from internal, so set i start from 1
	var fileLineNum string
	for i := 3; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && !strings.HasSuffix(file, "_test.go") && (!r.isSysCom(file) || strings.HasSuffix(file, "healthCheck.go")) { // !strings.HasPrefix(file, gormSourceDir) ||
			fileLineNum = file + ":" + strconv.FormatInt(int64(line), 10)
			break
		}
	}
	return fileLineNum
}

var comNames = []string{"/farseer-go/async/", "/farseer-go/cache/", "/farseer-go/cacheMemory/", "/farseer-go/collections/", "/farseer-go/data/", "/farseer-go/elasticSearch/", "/farseer-go/etcd/", "/farseer-go/eventBus/", "/farseer-go/fs/", "/farseer-go/linkTrace/", "/farseer-go/mapper/", "/farseer-go/queue/", "/farseer-go/rabbit/", "/farseer-go/redis/", "/farseer-go/redisStream/", "/farseer-go/tasks/", "/farseer-go/utils/", "/farseer-go/webapi/"}

func (r *CompositionLogger) isSysCom(file string) bool {
	for _, comName := range comNames {
		if strings.Contains(file, comName) {
			return true
		}
	}
	return false
}
