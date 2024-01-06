package parse

import (
	"reflect"
	"strconv"
)

func stringToNumber(source string, defVal any, returnKind reflect.Kind) any {
	switch returnKind {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		defVal = int64ToNumber(source, defVal, returnKind)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		defVal = uintToNumber(source, defVal, returnKind)
	case reflect.Float32, reflect.Float64:
		defVal = floatToNumber(source, defVal, returnKind)
	default:
		return defVal
	}
	return defVal
}

func int64ToNumber(source string, defVal any, returnKind reflect.Kind) any {
	result, err := strconv.ParseInt(source, 10, 64)
	if err == nil {
		switch returnKind {
		case reflect.Int8:
			defVal = int8(result)
		case reflect.Int16:
			defVal = int16(result)
		case reflect.Int32:
			defVal = int32(result)
		case reflect.Int64:
			defVal = result
		case reflect.Int:
			defVal = int(result)
		default:
			return defVal
		}
	}
	return defVal
}

func uintToNumber(source string, defVal any, returnKind reflect.Kind) any {
	result, err := strconv.ParseUint(source, 10, 64)
	if err == nil {
		switch returnKind {
		case reflect.Uint8:
			defVal = uint8(result)
		case reflect.Uint16:
			defVal = uint16(result)
		case reflect.Uint32:
			defVal = uint32(result)
		case reflect.Uint64:
			defVal = result
		case reflect.Uint:
			defVal = uint(result)
		default:
			return defVal
		}
	}
	return defVal
}

func floatToNumber(source string, defVal any, returnKind reflect.Kind) any {
	result, err := strconv.ParseFloat(source, 64)
	if err == nil {
		switch returnKind {
		case reflect.Float32:
			defVal = float32(result)
		case reflect.Float64:
			defVal = result
		default:
			return defVal
		}
	}
	return defVal
}
