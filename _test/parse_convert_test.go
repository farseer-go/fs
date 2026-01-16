package test

import (
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/parse"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, 123, parse.ConvertValue("123", reflect.TypeOf(1)).(int))
	assert.Equal(t, 1, parse.Convert(1, 0))
	assert.Equal(t, int64(1), parse.Convert(1, int64(0)))
	assert.Equal(t, "1", parse.Convert(1, ""))
	assert.False(t, parse.Convert(0, false))
	assert.True(t, parse.Convert(int8(1), false))
	assert.True(t, parse.Convert(int16(1), false))
	assert.True(t, parse.Convert(int32(1), false))
	assert.True(t, parse.Convert(int64(1), false))
	assert.True(t, parse.Convert(uint(1), false))
	assert.True(t, parse.Convert(uint8(1), false))
	assert.True(t, parse.Convert(uint16(1), false))
	assert.True(t, parse.Convert(uint32(1), false))
	assert.True(t, parse.Convert(uint64(1), false))
	assert.True(t, parse.Convert(float32(1), false))
	assert.True(t, parse.Convert(float64(1), false))
	assert.False(t, parse.Convert(2, false))
	assert.True(t, parse.Convert(1, false))
	assert.Equal(t, int8(8), parse.Convert(8, int8(1)))
	assert.Equal(t, int16(8), parse.Convert(8, int16(1)))
	assert.Equal(t, int32(8), parse.Convert(8, int32(1)))
	assert.Equal(t, int64(8), parse.Convert(8, int64(1)))
	assert.Equal(t, uint(8), parse.Convert(8, uint(1)))
	assert.Equal(t, uint8(8), parse.Convert(8, uint8(1)))
	assert.Equal(t, uint16(8), parse.Convert(8, uint16(1)))
	assert.Equal(t, uint32(8), parse.Convert(8, uint32(1)))
	assert.Equal(t, uint64(8), parse.Convert(8, uint64(1)))
	assert.Equal(t, float32(8), parse.Convert(8, float32(1)))
	assert.Equal(t, float64(8), parse.Convert(8, float64(1)))
	assert.Equal(t, 8, parse.Convert(int8(8), 1))
	assert.Equal(t, 8, parse.Convert(int16(8), 1))
	assert.Equal(t, 8, parse.Convert(int32(8), 1))
	assert.Equal(t, 8, parse.Convert(int64(8), 1))
	assert.Equal(t, 8, parse.Convert(uint(8), 1))
	assert.Equal(t, 8, parse.Convert(uint8(8), 1))
	assert.Equal(t, 8, parse.Convert(uint16(8), 1))
	assert.Equal(t, 8, parse.Convert(uint32(8), 1))
	assert.Equal(t, 8, parse.Convert(uint64(8), 1))
	assert.Equal(t, 8, parse.Convert(float32(8), 1))
	assert.Equal(t, 8, parse.Convert(float64(8), 1))
	assert.Equal(t, "8", parse.Convert(int8(8), ""))
	assert.Equal(t, "8", parse.Convert(int16(8), ""))
	assert.Equal(t, "8", parse.Convert(int32(8), ""))
	assert.Equal(t, "8", parse.Convert(int64(8), ""))
	assert.Equal(t, "8", parse.Convert(uint(8), ""))
	assert.Equal(t, "8", parse.Convert(uint8(8), ""))
	assert.Equal(t, "8", parse.Convert(uint16(8), ""))
	assert.Equal(t, "8", parse.Convert(uint32(8), ""))
	assert.Equal(t, "8", parse.Convert(uint64(8), ""))
	assert.Equal(t, "8.1", parse.Convert(float32(8.1), ""))
	assert.Equal(t, "8.12", parse.Convert(float64(8.12), ""))

	assert.Equal(t, int8(8), parse.Convert("8", int8(1)))
	assert.Equal(t, int16(8), parse.Convert("8", int16(1)))
	assert.Equal(t, int32(8), parse.Convert("8", int32(1)))
	assert.Equal(t, int64(8), parse.Convert("8", int64(1)))
	assert.Equal(t, uint(8), parse.Convert("8", uint(1)))
	assert.Equal(t, uint8(8), parse.Convert("8", uint8(1)))
	assert.Equal(t, uint16(8), parse.Convert("8", uint16(1)))
	assert.Equal(t, uint32(8), parse.Convert("8", uint32(1)))
	assert.Equal(t, uint64(8), parse.Convert("8", uint64(1)))
	assert.Equal(t, float32(8), parse.Convert("8", float32(1)))
	assert.Equal(t, float64(8), parse.Convert("8", float64(1)))
	assert.True(t, parse.Convert("true", false))
	assert.True(t, parse.Convert("True", false))
	assert.False(t, parse.Convert("false", false))
	assert.False(t, parse.Convert("False", false))
	assert.Equal(t, "123", parse.Convert("123", ""))
	assert.Equal(t, 123, parse.Convert("123", 0))
	assert.Equal(t, 3, parse.Convert("123f", 3))
	assert.Equal(t, uint8(3), parse.Convert("123f", uint8(3)))
	assert.Equal(t, float32(3), parse.Convert("123f", float32(3)))
	assert.Equal(t, []int{1, 2, 3}, parse.Convert("1,2,3", []int{}))
	assert.Equal(t, []string{"1", "2", "3"}, parse.Convert("1,2,3", []string{}))

	assert.Equal(t, struct{}{}, parse.Convert("123", struct{}{}))

	assert.True(t, parse.Convert(true, false))
	assert.False(t, parse.Convert(false, false))
	assert.Equal(t, struct{}{}, parse.Convert(false, struct{}{}))
	assert.Equal(t, 1, parse.Convert(true, 0))
	assert.Equal(t, 0, parse.Convert(false, 0))
	assert.Equal(t, "true", parse.Convert(true, ""))
	assert.Equal(t, "false", parse.Convert(false, ""))

	assert.Equal(t, []float64{1, 2, 3}, parse.Convert([]int{1, 2, 3}, []float64{}))

	t.Run("time.Time转time.Time", func(t *testing.T) {
		time1 := time.Now()
		time2 := parse.Convert(time1, time.UnixMilli(0))
		assert.Equal(t, time1.Year(), time2.Year())
		assert.Equal(t, time1.Month(), time2.Month())
		assert.Equal(t, time1.Day(), time2.Day())
		assert.Equal(t, time1.Hour(), time2.Hour())
		assert.Equal(t, time1.Minute(), time2.Minute())
		assert.Equal(t, time1.Second(), time2.Second())
		assert.Equal(t, time1.UnixMilli(), time2.UnixMilli())
		assert.Equal(t, time1.UnixMicro(), time2.UnixMicro())
		assert.Equal(t, time1.UnixNano(), time2.UnixNano())
	})

	t.Run("time.Time转DateTime", func(t *testing.T) {
		time1 := time.Now()
		dt := parse.Convert(time1, dateTime.DateTime{})
		assert.Equal(t, time1.Year(), dt.Year())
		assert.Equal(t, int(time1.Month()), dt.Month())
		assert.Equal(t, time1.Day(), dt.Day())
		assert.Equal(t, time1.Hour(), dt.Hour())
		assert.Equal(t, time1.Minute(), dt.Minute())
		assert.Equal(t, time1.Second(), dt.Second())
		assert.Equal(t, time1.UnixMilli(), dt.UnixMilli())
		assert.Equal(t, time1.UnixMicro(), dt.UnixMicro())
		assert.Equal(t, time1.UnixNano(), dt.UnixNano())
	})

	t.Run("DateTime转time.Time", func(t *testing.T) {
		dt := dateTime.Now()
		time1 := parse.Convert(dt, time.UnixMilli(0))
		assert.Equal(t, dt.Year(), time1.Year())
		assert.Equal(t, dt.Month(), int(time1.Month()))
		assert.Equal(t, dt.Day(), time1.Day())
		assert.Equal(t, dt.Hour(), time1.Hour())
		assert.Equal(t, dt.Minute(), time1.Minute())
		assert.Equal(t, dt.Second(), time1.Second())
		assert.Equal(t, dt.UnixMilli(), time1.UnixMilli())
		assert.Equal(t, dt.UnixMicro(), time1.UnixMicro())
		assert.Equal(t, dt.UnixNano(), time1.UnixNano())
	})

	t.Run("DateTime转DateTime", func(t *testing.T) {
		dt := dateTime.Now()
		dt2 := parse.Convert(dt, dateTime.New(time.UnixMilli(0)))
		assert.Equal(t, dt.Year(), dt2.Year())
		assert.Equal(t, dt.Month(), dt2.Month())
		assert.Equal(t, dt.Day(), dt2.Day())
		assert.Equal(t, dt.Hour(), dt2.Hour())
		assert.Equal(t, dt.Minute(), dt2.Minute())
		assert.Equal(t, dt.Second(), dt2.Second())
		assert.Equal(t, dt.UnixMilli(), dt2.UnixMilli())
		assert.Equal(t, dt.UnixMicro(), dt2.UnixMicro())
		assert.Equal(t, dt.UnixNano(), dt2.UnixNano())
	})

	t.Run("string转time.Time", func(t *testing.T) {
		dt := parse.Convert("2023-09-15", time.Time{})
		assert.Equal(t, 2023, dt.Year())
		assert.Equal(t, 9, int(dt.Month()))
		assert.Equal(t, 15, dt.Day())
	})

	t.Run("string转DateTime", func(t *testing.T) {
		dt := parse.Convert("2023-09-15", dateTime.DateTime{})
		assert.Equal(t, 2023, dt.Year())
		assert.Equal(t, 9, dt.Month())
		assert.Equal(t, 15, dt.Day())
	})

	t.Run("数字转enum", func(t *testing.T) {
		type EnumUint8 uint8
		const (
			All   EnumUint8 = 66 // 全部
			Other EnumUint8 = 33 // 全部
		)
		assert.Equal(t, All, parse.Convert(int8(66), Other))
		assert.Equal(t, All, parse.Convert(uint8(66), Other))
		assert.Equal(t, All, parse.Convert(int64(66), Other))
		assert.Equal(t, All, parse.Convert(int32(66), Other))
		assert.Equal(t, All, parse.Convert(66, Other))

		type EnumInt8 int8
		const (
			All2   EnumInt8 = 66 // 全部
			Other2 EnumInt8 = 33 // 全部
		)
		assert.Equal(t, All2, parse.Convert(int8(66), Other2))
		assert.Equal(t, All2, parse.Convert(uint8(66), Other2))
		assert.Equal(t, All2, parse.Convert(int64(66), Other2))
		assert.Equal(t, All2, parse.Convert(int32(66), Other2))
		assert.Equal(t, All2, parse.Convert(66, Other2))
	})

	t.Run("字符串转enum", func(t *testing.T) {
		type EnumUint8 uint8
		const (
			All   EnumUint8 = 66 // 全部
			Other EnumUint8 = 33 // 全部
		)
		assert.Equal(t, Other, parse.Convert("33", All))
	})
}

func TestEnum_uint8(t *testing.T) {
	type Enum uint8
	const (
		All Enum = 66 // 全部
	)
	assert.Equal(t, 66, parse.ToInt(All))
	assert.Equal(t, int8(66), parse.ToInt8(All))
	assert.Equal(t, int16(66), parse.ToInt16(All))
	assert.Equal(t, int32(66), parse.ToInt32(All))
	assert.Equal(t, int64(66), parse.ToInt64(All))
}
