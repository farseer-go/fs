package core

// ApiResponseInt 标准的API Response结构（默认int值）
type ApiResponseInt ApiResponse[int]

// ApiResponseIntSuccess 接口调用成功后返回的Json
func ApiResponseIntSuccess(statusMessage string, data int) ApiResponseInt {
	return ApiResponseInt{
		Status:        true,
		StatusMessage: statusMessage,
		StatusCode:    200,
		Data:          data,
	}
}

// ApiResponseIntError 接口调用失时返回的Json
func ApiResponseIntError(statusMessage string, statusCode int) ApiResponseInt {
	return ApiResponseInt{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    statusCode,
	}
}

// ApiResponseIntError403 接口调用失时返回的Json
func ApiResponseIntError403(statusMessage string) ApiResponseInt {
	return ApiResponseInt{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    403,
	}
}
