package fastReflect

import (
	"reflect"
	"unsafe"
)

func MakeSlice(typMeta *TypeMeta, arr []any) any {
	// 创建数组（耗时65ns）
	slice := reflect.MakeSlice(typMeta.ReflectType, len(arr), len(arr))
	slicePtr := slice.Pointer()
	for i := 0; i < len(arr); i++ {
		// 找到当前索引位置的内存地址。起始位置 + 每个元素占用的字节大小 ，得到第N个索引的内存起始位置
		itemPtr := unsafe.Pointer(slicePtr + uintptr(i)*typMeta.ItemMeta.Size)
		switch typMeta.ItemMeta.Kind {
		case reflect.Int:
			*(*int)(itemPtr) = arr[i].(int)
		case reflect.String:
			*(*string)(itemPtr) = arr[i].(string)
		default:
			println(itemPtr)
		}
	}
	return slice.Interface()
}

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
