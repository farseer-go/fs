package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Get http get，支持请求超时设置，单位：ms
func Get(url string, body map[string]string, contentType string, requestTimeout int) string {
	client := http.Client{
		Timeout: time.Duration(requestTimeout) * time.Millisecond,
	}
	bytesData, _ := json.Marshal(body)
	request, _ := http.NewRequest("GET", url, bytes.NewReader(bytesData))
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
	rspBody, _ := ioutil.ReadAll(resp.Body)
	return string(rspBody)
}

// GetForm http get，application/x-www-form-urlencoded
func GetForm(url string, body map[string]string, requestTimeout int) string {
	return Post(url, body, "application/x-www-form-urlencoded", requestTimeout)
}

// GetFormWithoutBody http get，application/x-www-form-urlencoded，
func GetFormWithoutBody(url string, body map[string]string, requestTimeout int) string {
	return Post(url, body, "application/x-www-form-urlencoded", requestTimeout)
}

// GetJson Post方式将结果反序列化成TReturn
func GetJson[TReturn any](url string, body map[string]string, requestTimeout int) TReturn {
	rspJson := Post(url, body, "application/json", requestTimeout)
	var val TReturn
	json.Unmarshal([]byte(rspJson), &val)
	return val
}
