package core

// ApiResponseLong 标准的API Response结构（默认int值）
type ApiResponseLong ApiResponse[int64]

// ApiResponseLongSuccess 接口调用成功后返回的Json
func ApiResponseLongSuccess(statusMessage string, data int64) ApiResponseLong {
	return ApiResponseLong{
		Status:        true,
		StatusMessage: statusMessage,
		StatusCode:    200,
		Data:          data,
	}
}

// ApiResponseLongError 接口调用失时返回的Json
func ApiResponseLongError(statusMessage string, statusCode int) ApiResponseLong {
	return ApiResponseLong{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    statusCode,
	}
}

// ApiResponseLongError403 接口调用失时返回的Json
func ApiResponseLongError403(statusMessage string) ApiResponseLong {
	return ApiResponseLong{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    403,
	}
}
