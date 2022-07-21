package parse

import (
	"strconv"
	"strings"
)

// Convert 通用的类型转换
func Convert[T any](source any, defVal T) T {
	// 根据source值的类型做转换适配
	switch source.(type) {
	case int:
		return convertInt(source.(int), defVal).(T)
	case string:
		return convertString(source.(string), defVal).(T)
	}
	return defVal
}

func convertInt(source int, defVal any) any {
	switch defVal.(type) {
	case bool:
		return any(source == 1)
	case string:
		return strconv.Itoa(source)
	case int64:
		return int64(source)
	}
	return defVal
}

func convertString(source string, defVal any) any {
	switch defVal.(type) {
	case bool:
		return any(strings.ToLower(source) == "true")
	case string:
		return source
	case int, int32:
		if result, isOk := strconv.Atoi(source); isOk == nil {
			return result
		}
	case int64:
		if result, isOk := strconv.ParseInt(source, 10, 64); isOk == nil {
			return result
		}
	}
	return defVal
}

type numberType interface {
	~int | ~int64 | ~int8 | ~int16 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}
