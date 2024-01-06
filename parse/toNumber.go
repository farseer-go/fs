package parse

import "reflect"

func anyToNumber(source any, sourceKind reflect.Kind, defVal any, returnKind reflect.Kind) any {
	switch s := source.(type) {
	case int8:
		return numberToNumber(s, defVal, returnKind)
	case int16:
		return numberToNumber(s, defVal, returnKind)
	case int32:
		return numberToNumber(s, defVal, returnKind)
	case int64:
		return numberToNumber(s, defVal, returnKind)
	case int:
		return numberToNumber(s, defVal, returnKind)
	case uint:
		return numberToNumber(s, defVal, returnKind)
	case uint8:
		return numberToNumber(s, defVal, returnKind)
	case uint16:
		return numberToNumber(s, defVal, returnKind)
	case uint32:
		return numberToNumber(s, defVal, returnKind)
	case uint64:
		return numberToNumber(s, defVal, returnKind)
	case float32:
		return numberToNumber(s, defVal, returnKind)
	case float64:
		return numberToNumber(s, defVal, returnKind)
	default: // 自定义的数字类型
		switch sourceKind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return numberToNumber(reflect.ValueOf(source).Int(), defVal, returnKind)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return numberToNumber(reflect.ValueOf(source).Uint(), defVal, returnKind)
		case reflect.Float32, reflect.Float64:
			return numberToNumber(reflect.ValueOf(source).Float(), defVal, returnKind)
		default:
			return source
		}
	}
}

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
	default:
		return defVal
	}
	return defVal
}
