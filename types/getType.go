package types

import (
	"reflect"
)

// GetRealType 获取真实类型
func GetRealType(val reflect.Value) reflect.Type {
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	return val.Type()
}

// GetInParam 获取方法的入参
func GetInParam(methodType reflect.Type) []reflect.Type {
	var arr []reflect.Type
	for inIndex := 0; inIndex < methodType.NumIn(); inIndex++ {
		arr = append(arr, methodType.In(inIndex))
	}
	return arr
}

// GetOutParam 获取方法的出参
func GetOutParam(methodType reflect.Type) []reflect.Type {
	var arr []reflect.Type
	for inIndex := 0; inIndex < methodType.NumOut(); inIndex++ {
		arr = append(arr, methodType.Out(inIndex))
	}
	return arr
}

// IsDtoModel 当只有一个参数，且非集合类型，又是结构类型时，判断为DTO
func IsDtoModel(lst []reflect.Type) bool {
	if len(lst) != 1 {
		return false
	}

	return !IsCollections(lst[0]) && lst[0].Kind() == reflect.Struct
}
