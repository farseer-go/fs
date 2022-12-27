package parse

import (
	"reflect"
	"strconv"
)

func numberToString(source any, defValue any, sourceKind reflect.Kind) any {
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
	return defValue
}
