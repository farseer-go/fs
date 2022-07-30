package http

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"time"
)

// Get http get，支持请求超时设置，单位：ms
func httpRequest(methodName string, url string, body any, contentType string, requestTimeout int) string {
	client := fasthttp.Client{}

	// request
	request := fasthttp.AcquireRequest()

	// url
	request.SetRequestURI(url)

	// request.body
	bytesData, _ := json.Marshal(body)
	request.SetBody(bytesData)

	// request.contentType
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	// Method
	request.Header.SetMethod(methodName)

	response := fasthttp.AcquireResponse()
	timeout := time.Duration(requestTimeout) * time.Millisecond
	err := client.DoTimeout(request, response, timeout)

	if err != nil {
		panic(err.Error())
	}
	return string(response.Body())
	/*
		// 原生
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
	*/
}
