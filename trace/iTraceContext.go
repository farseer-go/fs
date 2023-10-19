package trace

import "github.com/farseer-go/collections"

type ITraceContext interface {
	// End 结束
	End()
	// SetBody 设置webapi的响应报文
	SetBody(requestBody string, statusCode int, responseBody string)
	// GetTraceId 获取traceId
	GetTraceId() int64
	// GetStartTs 获取链路开启时间
	GetStartTs() int64
	// GetList 获取链路明细
	GetList() collections.List[ITraceDetail]
	// AddDetail 添加链路明细
	AddDetail(detail ITraceDetail)
}
