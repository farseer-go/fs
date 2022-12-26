package core

// ApiResponseString 标准的API Response结构（默认string值）
type ApiResponseString ApiResponse[string]

// ApiResponseStringSuccess 接口调用成功后返回的Json
func ApiResponseStringSuccess(statusMessage string, data string) ApiResponseString {
	return ApiResponseString{
		Status:        true,
		StatusMessage: statusMessage,
		StatusCode:    200,
		Data:          data,
	}
}

// ApiResponseStringError 接口调用失时返回的Json
func ApiResponseStringError(statusMessage string, statusCode int) ApiResponseString {
	return ApiResponseString{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    statusCode,
	}
}

// ApiResponseStringError403 接口调用失时返回的Json
func ApiResponseStringError403(statusMessage string) ApiResponseString {
	return ApiResponseString{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    403,
	}
}
