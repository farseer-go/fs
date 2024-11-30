package flog

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/snc"
	"github.com/farseer-go/fs/sonyflake"
	"github.com/farseer-go/fs/trace"
)

// FopsProvider 上传到FOPS
type FopsProvider struct {
}

func (r *FopsProvider) CreateLogger(categoryName string, formatter IFormatter, logLevel eumLogLevel.Enum) ILoggerPersistent {
	persistent := &fopsLoggerPersistent{formatter: formatter, fopsServer: configure.GetFopsServer(), queue: make(chan *LogData, 10000)}
	// 异步开启上传
	go persistent.enableUpload()
	return persistent
}

type fopsLoggerPersistent struct {
	formatter  IFormatter
	fopsServer string        // fops服务端
	queue      chan *LogData // 待上传的列表
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
		r.queue <- log
	}
}

// 开启上传
func (r *fopsLoggerPersistent) enableUpload() {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		var lst []*LogData
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
			fmt.Println(r.formatter.Formatter(&LogData{CreateAt: dateTime.Now(), LogLevel: eumLogLevel.Warning, Component: "", Content: err.Error(), newLine: true}))
		}
	}
}

type UploadRequest struct {
	List []*LogData
}

func (r *fopsLoggerPersistent) upload(lstLog []*LogData) error {
	bodyByte, _ := snc.Marshal(UploadRequest{List: lstLog})
	url := r.fopsServer + "flog/upload"
	newRequest, _ := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
	newRequest.Header.Set("Content-Type", "application/json")
	// 链路追踪
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 不验证 HTTPS 证书
			},
		},
	}
	rsp, err := client.Do(newRequest)
	if err != nil {
		return fmt.Errorf("上传日志到FOPS失败：%s", err.Error())
	}

	apiRsp := core.NewApiResponseByReader[any](rsp.Body)
	if apiRsp.StatusCode != 200 {
		return fmt.Errorf("上传日志到FOPS失败（%v）：%s", rsp.StatusCode, apiRsp.StatusMessage)
	}

	return err
}
