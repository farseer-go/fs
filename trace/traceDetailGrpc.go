package trace

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/parse"
)

type TraceDetailGrpc struct {
	GrpcMethod          string                                 // post/get/put/delete
	GrpcUrl             string                                 // url
	GrpcHeaders         collections.Dictionary[string, string] // 头部
	GrpcRequestBody     string                                 // 入参
	GrpcResponseBody    string                                 // 出参
	GrpcResponseHeaders collections.Dictionary[string, string] // 响应头部
	GrpcStatusCode      int                                    // 状态码
}

func (receiver *TraceDetailGrpc) SetHttpRequest(url string, reqHead map[string]any, rspHead map[string]string, requestBody string, responseBody string, statusCode int) {
	receiver.GrpcUrl = url
	receiver.GrpcHeaders = collections.NewDictionary[string, string]()
	receiver.GrpcRequestBody = requestBody
	receiver.GrpcResponseBody = responseBody
	receiver.GrpcStatusCode = statusCode
	for k, v := range reqHead {
		receiver.GrpcHeaders.Add(k, parse.ToString(v))
	}
	if rspHead != nil {
		receiver.GrpcResponseHeaders = collections.NewDictionaryFromMap(rspHead)
	}
}
