package linq

import (
	"fmt"
	"reflect"
)

func addition(val1 any, val2 any) any {
	if val1 == nil {
		return val2
	}
	kind := reflect.TypeOf(val1).Kind()
	switch kind {
	case reflect.Int8:
		return val1.(int8) + val2.(int8)
	case reflect.Int16:
		return val1.(int16) + val2.(int16)
	case reflect.Int32:
		return val1.(int32) + val2.(int32)
	case reflect.Int64:
		return val1.(int64) + val2.(int64)
	case reflect.Int:
		return val1.(int) + val2.(int)
	case reflect.Uint:
		return val1.(uint) + val2.(uint)
	case reflect.Uint8:
		return val1.(uint8) + val2.(uint8)
	case reflect.Uint16:
		return val1.(uint16) + val2.(uint16)
	case reflect.Uint32:
		return val1.(uint32) + val2.(uint32)
	case reflect.Uint64:
		return val1.(uint64) + val2.(uint64)
	case reflect.Float32:
		return val1.(float32) + val2.(float32)
	case reflect.Float64:
		return val1.(float64) + val2.(float64)
	}
	panic(fmt.Errorf("该类型无法比较：%s", kind))
}
