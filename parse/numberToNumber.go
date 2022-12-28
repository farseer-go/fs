package parse

import "reflect"

func numberToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](source T, defVal any, returnKind reflect.Kind) any {
	switch returnKind {
	case reflect.Int8:
		defVal = int8(source)
	case reflect.Int16:
		defVal = int16(source)
	case reflect.Int32:
		defVal = int32(source)
	case reflect.Int64:
		defVal = int64(source)
	case reflect.Int:
		defVal = int(source)
	case reflect.Uint:
		defVal = uint(source)
	case reflect.Uint8:
		defVal = uint8(source)
	case reflect.Uint16:
		defVal = uint16(source)
	case reflect.Uint32:
		defVal = uint32(source)
	case reflect.Uint64:
		defVal = uint64(source)
	case reflect.Float32:
		defVal = float32(source)
	case reflect.Float64:
		defVal = float64(source)
	}
	return defVal
}
