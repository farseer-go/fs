package core

import (
	"io"

	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/snc"
)

// ApiResponse 标准的API Response结构
type ApiResponse[TData any] struct {
	// 操作是否成功
	Status bool
	// 返回状态代码
	StatusCode int
	// 返回消息内容
	StatusMessage string
	// 接口返回时间
	ServerAt dateTime.DateTime
	// 耗时
	ElapsedMilliseconds int64
	// 链路追踪ID
	TraceId string
	// 不同接口返回的值
	Data TData
}

// SetData 设置Data字段的值
func (receiver *ApiResponse[TData]) SetData(data TData) {
	receiver.Data = data
	receiver.ServerAt = dateTime.Now()
}

// ToJson 转成Json
func (receiver *ApiResponse[TData]) ToJson() string {
	bytes, _ := snc.Marshal(receiver)
	return string(bytes)
}

// ToBytes 转成Json字节
func (receiver *ApiResponse[TData]) ToBytes() []byte {
	bytes, _ := snc.Marshal(receiver)
	return bytes
}

// Success 接口调用成功后返回的Json
func Success[TData any](statusMessage string, data TData) ApiResponse[TData] {
	return ApiResponse[TData]{
		Status:        true,
		StatusMessage: statusMessage,
		StatusCode:    200,
		Data:          data,
		ServerAt:      dateTime.Now(),
	}
}

// Error 接口调用失时返回的Json
func Error[TData any](statusMessage string, statusCode int) ApiResponse[TData] {
	return ApiResponse[TData]{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    statusCode,
		ServerAt:      dateTime.Now(),
	}
}

// Error403 接口调用失时返回的Json
func Error403[TData any](statusMessage string) ApiResponse[TData] {
	return ApiResponse[TData]{
		Status:        false,
		StatusMessage: statusMessage,
		StatusCode:    403,
		ServerAt:      dateTime.Now(),
	}
}

// NewApiResponseByReader 创建实例
func NewApiResponseByReader[TData any](reader io.Reader) ApiResponse[TData] {
	body, _ := io.ReadAll(reader)
	return NewApiResponseByByte[TData](body)
}

// NewApiResponseByByte 创建实例
func NewApiResponseByByte[TData any](body []byte) ApiResponse[TData] {
	var apiResponse ApiResponse[TData]
	_ = snc.Unmarshal(body, &apiResponse)
	return apiResponse
}
