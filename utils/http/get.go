package http

import (
	"encoding/json"
	_ "github.com/valyala/fasthttp"
)

// Get http get，支持请求超时设置，单位：ms
func Get(url string, body any, contentType string, requestTimeout int) string {
	return httpRequest("GET", url, body, contentType, requestTimeout)
}

// GetForm http get，application/x-www-form-urlencoded
func GetForm(url string, body any, requestTimeout int) string {
	return httpRequest("GET", url, body, "application/x-www-form-urlencoded", requestTimeout)
}

// GetFormWithoutBody http get，application/x-www-form-urlencoded，
func GetFormWithoutBody(url string, body any, requestTimeout int) string {
	return httpRequest("GET", url, body, "application/x-www-form-urlencoded", requestTimeout)
}

// GetJson Post方式将结果反序列化成TReturn
func GetJson[TReturn any](url string, body any, requestTimeout int) TReturn {
	rspJson := httpRequest("GET", url, body, "application/json", requestTimeout)
	var val TReturn
	json.Unmarshal([]byte(rspJson), &val)
	return val
}
