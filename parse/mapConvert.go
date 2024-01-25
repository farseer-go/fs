package parse

import (
	"reflect"
	"strconv"
)

func EqualTo1(source any, sourceKind reflect.Kind) bool {
	switch sourceKind {
	case reflect.Int:
		return source.(int) == 1
	case reflect.Int8:
		return source.(int8) == int8(1)
	case reflect.Int16:
		return source.(int16) == int16(1)
	case reflect.Int32:
		return source.(int32) == int32(1)
	case reflect.Int64:
		return source.(int64) == int64(1)
	case reflect.Uint:
		return source.(uint) == uint(1)
	case reflect.Uint8:
		return source.(uint8) == uint8(1)
	case reflect.Uint16:
		return source.(uint16) == uint16(1)
	case reflect.Uint32:
		return source.(uint32) == uint32(1)
	case reflect.Uint64:
		return source.(uint64) == uint64(1)
	case reflect.Float32:
		return source.(float32) == float32(1)
	case reflect.Float64:
		return source.(float64) == float64(1)
	}
	return false
}

func NumberToString(source any, sourceKind reflect.Kind) string {
	switch sourceKind {
	case reflect.Int8:
		return strconv.Itoa(int(source.(int8)))
	case reflect.Int16:
		return strconv.Itoa(int(source.(int16)))
	case reflect.Int32:
		return strconv.Itoa(int(source.(int32)))
	case reflect.Int64:
		return strconv.FormatInt(source.(int64), 10)
	case reflect.Int:
		return strconv.Itoa(source.(int))
	case reflect.Uint:
		return strconv.FormatUint(uint64(source.(uint)), 10)
	case reflect.Uint8:
		return strconv.FormatUint(uint64(source.(uint8)), 10)
	case reflect.Uint16:
		return strconv.FormatUint(uint64(source.(uint16)), 10)
	case reflect.Uint32:
		return strconv.FormatUint(uint64(source.(uint32)), 10)
	case reflect.Uint64:
		return strconv.FormatUint(source.(uint64), 10)
	case reflect.Float32:
		return strconv.FormatFloat(float64(source.(float32)), 'g', 6, 64)
	case reflect.Float64:
		return strconv.FormatFloat(source.(float64), 'g', 6, 64)
	}
	return ""
}
