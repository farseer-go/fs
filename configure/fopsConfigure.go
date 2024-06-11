package configure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/farseer-go/fs/core"
	"net/http"
)

// DomainObject 配置中心
type fopsConfigureVO struct {
	AppName string // 应用名称
	Key     string // 配置KEY
	Ver     int    // 版本
	Value   string // 配置VALUE
}

func getFopsConfigure() ([]fopsConfigureVO, error) {
	bodyByte, _ := json.Marshal(map[string]string{"AppName": core.AppName})
	url := fopsServer + "configure/list"
	newRequest, _ := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
	newRequest.Header.Set("Content-Type", "application/json")
	// 读取配置
	client := &http.Client{}
	var lst []fopsConfigureVO
	rsp, err := client.Do(newRequest)
	if err != nil {
		return lst, fmt.Errorf("读取配置中心时失败：%s", err.Error())
	}

	apiRsp := core.NewApiResponseByReader[[]fopsConfigureVO](rsp.Body)
	if apiRsp.StatusCode != 200 {
		return lst, fmt.Errorf("读取配置中心时失败（%v）：%s", rsp.StatusCode, apiRsp.StatusMessage)
	}
	return apiRsp.Data, err
}
