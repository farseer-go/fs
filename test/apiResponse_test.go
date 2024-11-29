package test

import (
	"bytes"
	"testing"

	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/snc"
	"github.com/stretchr/testify/assert"
)

func TestApiResponse(t *testing.T) {
	api := core.Success("成功", "nice")
	apiByte, _ := snc.Marshal(api)
	assert.Equal(t, string(apiByte), api.ToJson())
	assert.Equal(t, apiByte, api.ToBytes())

	api.SetData("very nice")
	assert.Equal(t, "very nice", api.Data)

	byByte := core.NewApiResponseByReader[string](bytes.NewReader(api.ToBytes()))
	assert.Equal(t, api.ToJson(), byByte.ToJson())

	api2 := core.Error403[int]("错误")
	assert.Equal(t, 403, api2.StatusCode)
	assert.Equal(t, "错误", api2.StatusMessage)
	assert.Equal(t, 0, api2.Data)

	api3 := core.Error[bool]("成功", 500)
	assert.Equal(t, 500, api3.StatusCode)
	assert.Equal(t, "成功", api3.StatusMessage)
	assert.Equal(t, false, api3.Data)

	api4 := core.ApiResponse[int](core.ApiResponseIntSuccess("成功", 2))
	assert.Equal(t, 200, api4.StatusCode)
	assert.Equal(t, "成功", api4.StatusMessage)
	assert.Equal(t, 2, api4.Data)

	api5 := core.ApiResponse[int](core.ApiResponseIntError("错误", 100))
	assert.Equal(t, 100, api5.StatusCode)
	assert.Equal(t, "错误", api5.StatusMessage)
	assert.Equal(t, 0, api5.Data)

	api6 := core.ApiResponse[int](core.ApiResponseIntError403("错误"))
	assert.Equal(t, 403, api6.StatusCode)
	assert.Equal(t, "错误", api6.StatusMessage)
	assert.Equal(t, 0, api6.Data)

	api7 := core.ApiResponse[int64](core.ApiResponseLongSuccess("成功", 2))
	assert.Equal(t, 200, api7.StatusCode)
	assert.Equal(t, "成功", api7.StatusMessage)
	assert.Equal(t, int64(2), api7.Data)

	api8 := core.ApiResponse[int64](core.ApiResponseLongError("错误", 100))
	assert.Equal(t, 100, api8.StatusCode)
	assert.Equal(t, "错误", api8.StatusMessage)
	assert.Equal(t, int64(0), api8.Data)

	api9 := core.ApiResponse[int64](core.ApiResponseLongError403("错误"))
	assert.Equal(t, 403, api9.StatusCode)
	assert.Equal(t, "错误", api9.StatusMessage)
	assert.Equal(t, int64(0), api9.Data)

	api10 := core.ApiResponse[string](core.ApiResponseStringSuccess("成功", "steden"))
	assert.Equal(t, 200, api10.StatusCode)
	assert.Equal(t, "成功", api10.StatusMessage)
	assert.Equal(t, "steden", api10.Data)

	api11 := core.ApiResponse[string](core.ApiResponseStringError("错误", 100))
	assert.Equal(t, 100, api11.StatusCode)
	assert.Equal(t, "错误", api11.StatusMessage)
	assert.Equal(t, "", api11.Data)

	api12 := core.ApiResponse[string](core.ApiResponseStringError403("错误"))
	assert.Equal(t, 403, api12.StatusCode)
	assert.Equal(t, "错误", api12.StatusMessage)
	assert.Equal(t, "", api12.Data)

}
