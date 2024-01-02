package fops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/trace"
	"net/http"
)

type UploadTraceRequest struct {
	List any
}

// UploadTrace 上传链路记录
func UploadTrace(lstTraceContext any) error {
	bodyByte, _ := json.Marshal(UploadTraceRequest{List: lstTraceContext})
	url := configure.GetFopsServer() + "linkTrace/upload"
	newRequest, _ := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
	newRequest.Header.Set("Content-Type", "application/json")
	// 链路追踪
	if traceContext := container.Resolve[trace.IManager]().GetCurTrace(); traceContext != nil {
		newRequest.Header.Set("Trace-Id", parse.ToString(traceContext.GetTraceId()))
		newRequest.Header.Set("Trace-App-Name", core.AppName)
	}
	client := &http.Client{}
	rsp, err := client.Do(newRequest)
	if err != nil {
		return fmt.Errorf("上传链路记录到FOPS失败：%s", err.Error())
	}

	apiRsp := core.NewApiResponseByReader[any](rsp.Body)
	if apiRsp.StatusCode != 200 {
		return fmt.Errorf("上传链路记录到FOPS失败（%v）：%s", rsp.StatusCode, apiRsp.StatusMessage)
	}

	return err
}
