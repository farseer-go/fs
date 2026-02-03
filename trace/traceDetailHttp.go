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

func (receiver *TraceDetailHttp) SetRequest(url string, reqHead map[string]any, requestBody string) {
	receiver.HttpUrl = url
	receiver.HttpHeaders = make(map[string]string)
	receiver.HttpRequestBody = requestBody
	for k, v := range reqHead {
		receiver.HttpHeaders[k] = parse.ToString(v)
	}
}

func (receiver *TraceDetailHttp) SetResponse(rspHead map[string]string, responseBody string, statusCode int) {
	receiver.HttpResponseBody = responseBody
	receiver.HttpStatusCode = statusCode

	if rspHead != nil {
		receiver.HttpResponseHeaders = rspHead
	}
}
