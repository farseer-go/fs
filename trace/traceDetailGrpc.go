package trace

import (
	"github.com/farseer-go/fs/parse"
)

type TraceDetailGrpc struct {
	GrpcMethod          string            `json:",omitempty"` // post/get/put/delete
	GrpcUrl             string            `json:",omitempty"` // url
	GrpcHeaders         map[string]string `json:",omitempty"` // 头部
	GrpcRequestBody     string            `json:",omitempty"` // 入参
	GrpcResponseBody    string            `json:",omitempty"` // 出参
	GrpcResponseHeaders map[string]string `json:",omitempty"` // 响应头部
	GrpcStatusCode      int               `json:",omitempty"` // 状态码
}

func (receiver *TraceDetailGrpc) SetHttpRequest(url string, reqHead map[string]any, rspHead map[string]string, requestBody string, responseBody string, statusCode int) {
	receiver.GrpcUrl = url
	receiver.GrpcHeaders = make(map[string]string)
	receiver.GrpcRequestBody = requestBody
	receiver.GrpcResponseBody = responseBody
	receiver.GrpcStatusCode = statusCode
	for k, v := range reqHead {
		receiver.GrpcHeaders[k] = parse.ToString(v)
	}
	if rspHead != nil {
		receiver.GrpcResponseHeaders = make(map[string]string)
	}
}
