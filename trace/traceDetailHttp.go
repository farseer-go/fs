package trace

import (
	"github.com/farseer-go/fs/parse"
)

type TraceDetailHttp struct {
	HttpMethod          string            `json:",omitempty"` // post/get/put/delete
	HttpUrl             string            `json:",omitempty"` // url
	HttpHeaders         map[string]string `json:",omitempty"` // 头部
	HttpRequestBody     string            `json:",omitempty"` // 入参
	HttpResponseBody    string            `json:",omitempty"` // 出参
	HttpResponseHeaders map[string]string `json:",omitempty"` // 响应头部
	HttpStatusCode      int               `json:",omitempty"` // 状态码
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
