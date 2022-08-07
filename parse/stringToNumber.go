package parse

import (
	"reflect"
	"strconv"
)

func stringToNumber(source string, defVal any, returnKind reflect.Kind) any {
	if returnKind == reflect.Int8 || returnKind == reflect.Int16 || returnKind == reflect.Int32 || returnKind == reflect.Int || returnKind == reflect.Int64 {
		result, err := strconv.ParseInt(source, 10, 64)
		if err != nil {
			return defVal
		}
		switch returnKind {
		case reflect.Int8:
			return int8(result)
		case reflect.Int16:
			return int16(result)
		case reflect.Int32:
			return int32(result)
		case reflect.Int64:
			return result
		case reflect.Int:
			return int(result)
		}
	}

	if returnKind == reflect.Uint8 || returnKind == reflect.Uint16 || returnKind == reflect.Uint32 || returnKind == reflect.Uint || returnKind == reflect.Uint64 {
		result, err := strconv.ParseUint(source, 10, 64)
		if err != nil {
			return defVal
		}
		switch returnKind {
		case reflect.Uint8:
			return uint8(result)
		case reflect.Uint16:
			return uint16(result)
		case reflect.Uint32:
			return uint32(result)
		case reflect.Uint64:
			return result
		case reflect.Uint:
			return uint(result)
		}
	}

	if returnKind == reflect.Float32 || returnKind == reflect.Float64 {
		result, err := strconv.ParseFloat(source, 64)
		if err != nil {
			return defVal
		}
		switch returnKind {
		case reflect.Float32:
			return float32(result)
		case reflect.Float64:
			return result
		}
	}
	return defVal
}
