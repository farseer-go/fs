package parse

import (
	"reflect"
	"strconv"
)

func numberToString(source any, defValue any, sourceKind reflect.Kind) any {
	switch sourceKind {
	case reflect.Int8:
		defValue = strconv.Itoa(int(source.(int8)))
	case reflect.Int16:
		defValue = strconv.Itoa(int(source.(int16)))
	case reflect.Int32:
		defValue = strconv.Itoa(int(source.(int32)))
	case reflect.Int64:
		defValue = strconv.FormatInt(source.(int64), 10)
	case reflect.Int:
		defValue = strconv.Itoa(source.(int))
	case reflect.Uint:
		defValue = strconv.FormatUint(uint64(source.(uint)), 10)
	case reflect.Uint8:
		defValue = strconv.FormatUint(uint64(source.(uint8)), 10)
	case reflect.Uint16:
		defValue = strconv.FormatUint(uint64(source.(uint16)), 10)
	case reflect.Uint32:
		defValue = strconv.FormatUint(uint64(source.(uint32)), 10)
	case reflect.Uint64:
		defValue = strconv.FormatUint(source.(uint64), 10)
	case reflect.Float32:
		defValue = strconv.FormatFloat(float64(source.(float32)), 'g', 6, 64)
	case reflect.Float64:
		defValue = strconv.FormatFloat(source.(float64), 'g', 6, 64)
	default:
		return defValue
	}
	return defValue
}
