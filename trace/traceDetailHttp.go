package trace

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/parse"
)

type TraceDetailHttp struct {
	HttpMethod          string                                 // post/get/put/delete
	HttpUrl             string                                 // url
	HttpHeaders         collections.Dictionary[string, string] // 头部
	HttpRequestBody     string                                 // 入参
	HttpResponseBody    string                                 // 出参
	HttpResponseHeaders collections.Dictionary[string, string] // 响应头部
	HttpStatusCode      int                                    // 状态码
}

func (receiver *TraceDetailHttp) SetHttpRequest(url string, reqHead map[string]any, rspHead map[string]string, requestBody string, responseBody string, statusCode int) {
	receiver.HttpUrl = url
	receiver.HttpHeaders = collections.NewDictionary[string, string]()
	receiver.HttpRequestBody = requestBody
	receiver.HttpResponseBody = responseBody
	receiver.HttpStatusCode = statusCode
	for k, v := range reqHead {
		receiver.HttpHeaders.Add(k, parse.ToString(v))
	}

	if rspHead != nil {
		receiver.HttpResponseHeaders = collections.NewDictionaryFromMap(rspHead)
	}
}
