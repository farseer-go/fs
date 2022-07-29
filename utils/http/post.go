package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Post http post，支持请求超时设置，单位：ms
func Post(url string, body map[string]string, contentType string, requestTimeout int) string {
	client := http.Client{
		Timeout: time.Duration(requestTimeout) * time.Millisecond,
	}
	bytesData, _ := json.Marshal(body)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
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

// PostForm http post，application/x-www-form-urlencoded
func PostForm(url string, body map[string]string, requestTimeout int) string {
	return Post(url, body, "application/x-www-form-urlencoded", requestTimeout)
}

// PostFormWithoutBody http post，application/x-www-form-urlencoded
func PostFormWithoutBody(url string, body map[string]string, requestTimeout int) string {
	return Post(url, body, "application/x-www-form-urlencoded", requestTimeout)
}

// PostJson Post方式将结果反序列化成TReturn
func PostJson[TReturn any](url string, body map[string]string, requestTimeout int) TReturn {
	rspJson := Post(url, body, "application/json", requestTimeout)
	var val TReturn
	json.Unmarshal([]byte(rspJson), &val)
	return val
}
