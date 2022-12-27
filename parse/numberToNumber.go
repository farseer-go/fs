package parse

import "reflect"

func numberToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](source T, defVal any, returnKind reflect.Kind) any {
	switch returnKind {
	case reflect.Int8:
		return int8(source)
	case reflect.Int16:
		return int16(source)
	case reflect.Int32:
		return int32(source)
	case reflect.Int64:
		return int64(source)
	case reflect.Int:
		return int(source)
	case reflect.Uint:
		return uint(source)
	case reflect.Uint8:
		return uint8(source)
	case reflect.Uint16:
		return uint16(source)
	case reflect.Uint32:
		return uint32(source)
	case reflect.Uint64:
		return uint64(source)
	case reflect.Float32:
		return float32(source)
	case reflect.Float64:
		return float64(source)
	}
	return defVal
}
