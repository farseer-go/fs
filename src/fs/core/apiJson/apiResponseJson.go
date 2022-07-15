package apiJson

type ApiResponseJson[TData any] struct {
	// 操作是否成功
	Status bool
	// 返回状态代码
	StatusCode int
	// 返回消息内容
	StatusMessage string
	// 不同接口返回的值
	Data TData
}

// SetData 设置Data字段的值
func (receiver *ApiResponseJson[TData]) SetData(data TData) {
	receiver.Data = data
}

// Success 接口调用成功后返回的Json
func Success[TData any](statusMessage string, data TData) ApiResponseJson[TData] {
	return ApiResponseJson[TData]{
		Status:        true,
		StatusMessage: statusMessage,
		StatusCode:    200,
		Data:          data,
	}
}

// Error 接口调用失时返回的Json
func Error(statusMessage string, statusCode int) ApiResponseJson[string] {
	return ApiResponseJson[string]{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    statusCode,
	}
}

// Error403 接口调用失时返回的Json
func Error403(statusMessage string) ApiResponseJson[string] {
	return ApiResponseJson[string]{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    403,
	}
}
