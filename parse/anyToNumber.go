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
