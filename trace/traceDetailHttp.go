package trace

import (
	"github.com/farseer-go/fs/parse"
)

type TraceDetailHttp struct {
	HttpMethod          string            // post/get/put/delete
	HttpUrl             string            // url
	HttpHeaders         map[string]string // 头部
	HttpRequestBody     string            // 入参
	HttpResponseBody    string            // 出参
	HttpResponseHeaders map[string]string // 响应头部
	HttpStatusCode      int               // 状态码
}

func (receiver *TraceDetailHttp) SetHttpRequest(url string, reqHead map[string]any, rspHead map[string]string, requestBody string, responseBody string, statusCode int) {
	receiver.HttpUrl = url
	receiver.HttpHeaders = make(map[string]string)
	receiver.HttpRequestBody = requestBody
	receiver.HttpResponseBody = responseBody
	receiver.HttpStatusCode = statusCode
	for k, v := range reqHead {
		receiver.HttpHeaders[k] = parse.ToString(v)
	}

	if rspHead != nil {
		receiver.HttpResponseHeaders = rspHead
	}
}
