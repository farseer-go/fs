package test

import (
	"bytes"
	"github.com/farseer-go/fs/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApiResponse(t *testing.T) {
	api := core.Success("成功", "nice")
	assert.Equal(t, "{\"Status\":true,\"StatusCode\":200,\"StatusMessage\":\"成功\",\"Data\":\"nice\"}", api.ToJson())
	assert.Equal(t, "{\"Status\":true,\"StatusCode\":200,\"StatusMessage\":\"成功\",\"Data\":\"nice\"}", string(api.ToBytes()))

	api.SetData("very nice")
	assert.Equal(t, "very nice", api.Data)

	byByte := core.NewApiResponseByReader[string](bytes.NewReader(api.ToBytes()))
	assert.Equal(t, api.ToJson(), byByte.ToJson())

	api2 := core.Error403[int]("错误")
	assert.Equal(t, "{\"Status\":false,\"StatusCode\":403,\"StatusMessage\":\"错误\",\"Data\":0}", api2.ToJson())
	api3 := core.Error[bool]("成功", 500)
	assert.Equal(t, "{\"Status\":false,\"StatusCode\":500,\"StatusMessage\":\"成功\",\"Data\":false}", api3.ToJson())

	api4 := core.ApiResponse[int](core.ApiResponseIntSuccess("成功", 2))
	assert.Equal(t, "{\"Status\":true,\"StatusCode\":200,\"StatusMessage\":\"成功\",\"Data\":2}", api4.ToJson())

	api5 := core.ApiResponse[int](core.ApiResponseIntError("错误", 100))
	assert.Equal(t, "{\"Status\":false,\"StatusCode\":100,\"StatusMessage\":\"错误\",\"Data\":0}", api5.ToJson())

	api6 := core.ApiResponse[int](core.ApiResponseIntError403("错误"))
	assert.Equal(t, "{\"Status\":false,\"StatusCode\":403,\"StatusMessage\":\"错误\",\"Data\":0}", api6.ToJson())

	api7 := core.ApiResponse[int64](core.ApiResponseLongSuccess("成功", 2))
	assert.Equal(t, "{\"Status\":true,\"StatusCode\":200,\"StatusMessage\":\"成功\",\"Data\":2}", api7.ToJson())

	api8 := core.ApiResponse[int64](core.ApiResponseLongError("错误", 100))
	assert.Equal(t, "{\"Status\":false,\"StatusCode\":100,\"StatusMessage\":\"错误\",\"Data\":0}", api8.ToJson())

	api9 := core.ApiResponse[int64](core.ApiResponseLongError403("错误"))
	assert.Equal(t, "{\"Status\":false,\"StatusCode\":403,\"StatusMessage\":\"错误\",\"Data\":0}", api9.ToJson())

	api10 := core.ApiResponse[string](core.ApiResponseStringSuccess("成功", "steden"))
	assert.Equal(t, "{\"Status\":true,\"StatusCode\":200,\"StatusMessage\":\"成功\",\"Data\":\"steden\"}", api10.ToJson())

	api11 := core.ApiResponse[string](core.ApiResponseStringError("错误", 100))
	assert.Equal(t, "{\"Status\":false,\"StatusCode\":100,\"StatusMessage\":\"错误\",\"Data\":\"\"}", api11.ToJson())

	api12 := core.ApiResponse[string](core.ApiResponseStringError403("错误"))
	assert.Equal(t, "{\"Status\":false,\"StatusCode\":403,\"StatusMessage\":\"错误\",\"Data\":\"\"}", api12.ToJson())

}
