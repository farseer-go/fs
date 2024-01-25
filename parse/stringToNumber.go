package parse

import (
	"reflect"
	"strconv"
)

func StringToNumber(source string, defVal any, defValKind reflect.Kind) any {
	switch defValKind {
	case reflect.Float32:
		if result, err := strconv.ParseFloat(source, 64); err == nil {
			return float32(result)
		}
	case reflect.Float64:
		if result, err := strconv.ParseFloat(source, 64); err == nil {
			return result
		}
	case reflect.Uint8:
		if result, err := strconv.ParseUint(source, 10, 64); err == nil {
			return uint8(result)
		}
	case reflect.Uint16:
		if result, err := strconv.ParseUint(source, 10, 64); err == nil {
			return uint16(result)
		}
	case reflect.Uint32:
		if result, err := strconv.ParseUint(source, 10, 64); err == nil {
			return uint32(result)
		}
	case reflect.Uint64:
		if result, err := strconv.ParseUint(source, 10, 64); err == nil {
			return result
		}
	case reflect.Uint:
		if result, err := strconv.ParseUint(source, 10, 64); err == nil {
			return uint(result)
		}
	case reflect.Int8:
		if result, err := strconv.ParseInt(source, 10, 64); err == nil {
			return int8(result)
		}
	case reflect.Int16:
		if result, err := strconv.ParseInt(source, 10, 64); err == nil {
			return int16(result)
		}
	case reflect.Int32:
		if result, err := strconv.ParseInt(source, 10, 64); err == nil {
			return int32(result)
		}
	case reflect.Int64:
		if result, err := strconv.ParseInt(source, 10, 64); err == nil {
			return result
		}
	case reflect.Int:
		if result, err := strconv.ParseInt(source, 10, 64); err == nil {
			return int(result)
		}
	}
	return defVal
}
