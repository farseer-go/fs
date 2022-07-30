package http

import (
	"encoding/json"
)

// Post http post，支持请求超时设置，单位：ms
func Post(url string, body any, contentType string, requestTimeout int) string {
	return httpRequest("POST", url, body, contentType, requestTimeout)
}

// PostForm http post，application/x-www-form-urlencoded
func PostForm(url string, body any, requestTimeout int) string {
	return httpRequest("POST", url, body, "application/x-www-form-urlencoded", requestTimeout)
}

// PostFormWithoutBody http post，application/x-www-form-urlencoded
func PostFormWithoutBody(url string, requestTimeout int) string {
	return httpRequest("POST", url, nil, "application/x-www-form-urlencoded", requestTimeout)
}

// PostJson Post方式将结果反序列化成TReturn
func PostJson[TReturn any](url string, body any, requestTimeout int) TReturn {
	rspJson := httpRequest("POST", url, nil, "application/json", requestTimeout)
	var val TReturn
	json.Unmarshal([]byte(rspJson), &val)
	return val
}
