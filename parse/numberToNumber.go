package parse

import "reflect"

func anyToNumber(source any, sourceKind reflect.Kind, defVal any, returnKind reflect.Kind) any {
	var result = defVal
	switch sourceKind {
	case reflect.Int8:
		result = numberToNumber(source.(int8), defVal, returnKind)
	case reflect.Int16:
		result = numberToNumber(source.(int16), defVal, returnKind)
	case reflect.Int32:
		result = numberToNumber(source.(int32), defVal, returnKind)
	case reflect.Int64:
		result = numberToNumber(source.(int64), defVal, returnKind)
	case reflect.Int:
		result = numberToNumber(source.(int), defVal, returnKind)
	case reflect.Uint:
		result = numberToNumber(source.(uint), defVal, returnKind)
	case reflect.Uint8:
		result = numberToNumber(source.(uint8), defVal, returnKind)
	case reflect.Uint16:
		result = numberToNumber(source.(uint16), defVal, returnKind)
	case reflect.Uint32:
		result = numberToNumber(source.(uint32), defVal, returnKind)
	case reflect.Uint64:
		result = numberToNumber(source.(uint64), defVal, returnKind)
	case reflect.Float32:
		result = numberToNumber(source.(float32), defVal, returnKind)
	case reflect.Float64:
		result = numberToNumber(source.(float64), defVal, returnKind)
	}
	return result
}

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
