package fops

import (
	"bytes"
	"encoding/json"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/trace"
	"net/http"
	"time"
)

// RegisterApp 定时向FOPS中心注册应用信息
func RegisterApp() {
	// 先通过配置节点读
	fopsServer := configure.GetString("Fops.Server")
	if fopsServer != "" {
		fopsServer = configure.GetFopsServer()
		flog.LogBuffer <- flog.Green("【✓】") + "FOPS Center：" + flog.Blue(fopsServer)
		// 定时向FOPS中心注册应用信息
		go register()
	}
}

type RegisterAppRequest struct {
	AppName     string            // 应用名称
	AppId       int64             // 应用ID
	AppIp       string            // 应用IP
	HostName    string            // 主机名称
	ProcessId   int               // 进程Id
	StartupAt   dateTime.DateTime // 应用启动时间
	CpuUsage    float64           // CPU百分比
	MemoryUsage float64           // 内存百分比
}

// 每隔3秒，上传当前应用信息
func register() {
	for {
		bodyByte, _ := json.Marshal(RegisterAppRequest{StartupAt: core.StartupAt, AppName: core.AppName, HostName: core.HostName, AppId: core.AppId, AppIp: core.AppIp, ProcessId: core.ProcessId})
		url := configure.GetFopsServer() + "apps/register"
		newRequest, _ := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
		newRequest.Header.Set("Content-Type", "application/json")
		// 链路追踪
		if traceContext := container.Resolve[trace.IManager]().GetCurTrace(); traceContext != nil {
			newRequest.Header.Set("Trace-Id", parse.ToString(traceContext.GetTraceId()))
			newRequest.Header.Set("Trace-App-Name", core.AppName)
		}
		client := &http.Client{}
		rsp, err := client.Do(newRequest)

		if err == nil {
			apiRsp := core.NewApiResponseByReader[any](rsp.Body)
			if apiRsp.StatusCode != 200 {
				flog.Warningf("注册应用信息到FOPS失败（%v）%s", rsp.StatusCode, apiRsp.StatusMessage)
				<-time.NewTicker(20 * time.Second).C
				continue
			}
		} else {
			flog.Warningf("注册应用信息到FOPS失败：%s", err.Error())
			<-time.NewTicker(20 * time.Second).C
			continue
		}
		<-time.NewTicker(20 * time.Second).C
	}
}
