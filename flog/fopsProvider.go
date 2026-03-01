package flog

import (
	"strconv"
	"time"

	"github.com/farseer-go/fs/batchFileWriter"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/sonyflake"
	"github.com/farseer-go/fs/trace"
)

// FopsProvider 上传到FOPS
type FopsProvider struct {
}

func (r *FopsProvider) CreateLogger(categoryName string, formatter IFormatter, logLevel eumLogLevel.Enum) ILoggerPersistent {
	writer := batchFileWriter.NewWriter("/var/log/flog/"+core.AppName+"/", "log", "hour", 10, 0, time.Second*5, true)
	persistent := &fopsLoggerPersistent{formatter, writer}
	return persistent
}

type fopsLoggerPersistent struct {
	formatter  IFormatter
	FileWriter *batchFileWriter.BatchFileWriter
}

func (r *fopsLoggerPersistent) IsEnabled(logLevel eumLogLevel.Enum) bool {
	return true
}

func (r *fopsLoggerPersistent) Log(LogLevel eumLogLevel.Enum, log *LogData, exception error) {
	if LogLevel != eumLogLevel.NoneLevel {
		// 上传到FOPS时需要
		if t := trace.CurTraceContext.Get(); t != nil {
			log.TraceId = t.TraceId
		}
		log.Content = mustCompile.ReplaceAllString(log.Content, "")
		log.AppId = strconv.FormatInt(core.AppId, 10)
		log.AppName = core.AppName
		log.AppIp = core.AppIp
		log.LogId = strconv.FormatInt(sonyflake.GenerateId(), 10)

		r.FileWriter.Write(log)
	}
}

func (r *fopsLoggerPersistent) Close() {
	r.FileWriter.Close()
}
