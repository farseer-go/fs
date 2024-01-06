package parse

import "reflect"

// IsEqual 比较两个值是否相等
func IsEqual[TValue any](val1, val2 TValue) bool {
	valType := reflect.TypeOf(val1)
	switch valType.Kind() {
	case reflect.String:
		return any(val1).(string) == any(val2).(string)
	case reflect.Bool:
		return any(val1).(bool) == any(val2).(bool)
	case reflect.Int:
		return any(val1).(int) == any(val2).(int)
	case reflect.Int8:
		return any(val1).(int8) == any(val2).(int8)
	case reflect.Int16:
		return any(val1).(int16) == any(val2).(int16)
	case reflect.Int32:
		return any(val1).(int32) == any(val2).(int32)
	case reflect.Int64:
		return any(val1).(int64) == any(val2).(int64)
	case reflect.Uint:
		return any(val1).(uint) == any(val2).(uint)
	case reflect.Uint8:
		return any(val1).(uint8) == any(val2).(uint8)
	case reflect.Uint16:
		return any(val1).(uint16) == any(val2).(uint16)
	case reflect.Uint32:
		return any(val1).(uint32) == any(val2).(uint32)
	case reflect.Uint64:
		return any(val1).(uint64) == any(val2).(uint64)
	case reflect.Float64:
		return any(val1).(float64) == any(val2).(float64)
	case reflect.Float32:
		return any(val1).(float32) == any(val2).(float32)
	default:
		return false
	}
}
