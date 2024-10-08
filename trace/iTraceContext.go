package trace

type ITraceContext interface {
	// End 结束
	End(err error)
	// SetBody 设置webapi的响应报文
	SetBody(requestBody string, statusCode int, responseBody string)
	// SetResponseBody 设置webapi的响应报文
	SetResponseBody(responseBody string)
	// GetTraceId 获取traceId
	GetTraceId() string
	// GetStartTs 获取链路开启时间
	GetStartTs() int64
	// GetList 获取链路明细
	GetList() []any
	// AddDetail 添加链路明细
	AddDetail(detail ITraceDetail)
	// Error 异常信息
	Error(err error)
	// Ignore 忽略这次的链路追踪
	Ignore()
	// GetAppInfo 获取应用信息
	GetAppInfo() (string, string, string, string, string)
	// GetTraceLevel 得到当前链路层
	GetTraceLevel() int
}
