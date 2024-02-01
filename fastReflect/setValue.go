package fastReflect

import (
	"reflect"
	"unsafe"
)

func SetValue(ptr unsafe.Pointer, val any, valType *TypeMeta) {
	switch valType.Kind {
	case reflect.Int:
		*(*int)(ptr) = val.(int)
	case reflect.Int8:
		*(*int8)(ptr) = val.(int8)
	case reflect.Int16:
		*(*int16)(ptr) = val.(int16)
	case reflect.Int32:
		*(*int32)(ptr) = val.(int32)
	case reflect.Int64:
		*(*int64)(ptr) = val.(int64)
	case reflect.Uint:
		*(*uint)(ptr) = val.(uint)
	case reflect.Uint8:
		*(*uint8)(ptr) = val.(uint8)
	case reflect.Uint16:
		*(*uint16)(ptr) = val.(uint16)
	case reflect.Uint32:
		*(*uint32)(ptr) = val.(uint32)
	case reflect.Uint64:
		*(*uint64)(ptr) = val.(uint64)
	case reflect.Float32:
		*(*float32)(ptr) = val.(float32)
	case reflect.Float64:
		*(*float64)(ptr) = val.(float64)
	case reflect.String:
		*(*string)(ptr) = val.(string)
	case reflect.Bool:
		*(*bool)(ptr) = val.(bool)
	}
}
