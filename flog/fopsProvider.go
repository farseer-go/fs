package flog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/trace"
	"net/http"
	"time"
)

// FopsProvider 上传到FOPS
type FopsProvider struct {
}

func (r *FopsProvider) CreateLogger(categoryName string, formatter IFormatter, logLevel eumLogLevel.Enum) ILoggerPersistent {
	persistent := &fopsLoggerPersistent{formatter: formatter, fopsServer: configure.GetFopsServer(), queue: make(chan *logData, 10000)}
	// 异步开启上传
	go persistent.enableUpload()
	return persistent
}

type fopsLoggerPersistent struct {
	formatter  IFormatter
	fopsServer string        // fops服务端
	queue      chan *logData // 待上传的列表
}

func (r *fopsLoggerPersistent) IsEnabled(logLevel eumLogLevel.Enum) bool {
	return true
}

func (r *fopsLoggerPersistent) Log(LogLevel eumLogLevel.Enum, log *logData, exception error) {
	r.queue <- log
}

// 开启上传
func (r *fopsLoggerPersistent) enableUpload() {
	for range time.NewTicker(3 * time.Second).C {
		var lst []*logData
		// 当队列中有数据 且 取出的数量<1000时，则继续取出
		for len(r.queue) > 0 && len(lst) < 1000 {
			lst = append(lst, <-r.queue)
		}

		// 没有取到数据
		if len(lst) == 0 {
			continue
		}

		// 上传
		if err := r.upload(lst); err != nil {
			// 重新放回队列
			for i := 0; i < len(lst); i++ {
				r.queue <- lst[i]
			}
			// 不能使用flog.Error，如果此处执行了，会一直产生无用的错误信息
			fmt.Println(r.formatter.Formatter(&logData{CreateAt: dateTime.Now(), LogLevel: eumLogLevel.Warning, Component: "", Content: err.Error(), newLine: true}))
		}
	}
}

type UploadRequest struct {
	List []*logData
}

func (r *fopsLoggerPersistent) upload(lstLog []*logData) error {
	bodyByte, _ := json.Marshal(UploadRequest{List: lstLog})
	url := r.fopsServer + "flog/upload"
	newRequest, _ := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
	newRequest.Header.Set("Content-Type", "application/json")
	// 链路追踪
	if traceContext := container.Resolve[trace.IManager]().GetCurTrace(); traceContext != nil {
		newRequest.Header.Set("Trace-Id", parse.ToString(traceContext.GetTraceId()))
		newRequest.Header.Set("Trace-App-Name", fs.AppName)
	}
	client := &http.Client{}
	rsp, err := client.Do(newRequest)
	if err != nil {
		return err
	}

	apiRsp := core.NewApiResponseByReader[any](rsp.Body)
	if apiRsp.StatusCode != 200 {
		return fmt.Errorf("上传日志到%s失败（%v）：%s", url, rsp.StatusCode, apiRsp.StatusMessage)
	}

	return err
}
