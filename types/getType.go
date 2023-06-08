package types

import (
	"reflect"
)

// GetRealType 获取真实类型
func GetRealType(val reflect.Value) reflect.Type {
	realValue := val
	if realValue.Kind() == reflect.Pointer {
		realValue = realValue.Elem()
	}
	if realValue.Kind() == reflect.Interface && !realValue.IsZero() {
		realValue = realValue.Elem()
	}

	// 无效的值，只能返回原型
	if realValue.Kind() == reflect.Invalid {
		return val.Type()
	}

	return realValue.Type()
}

func GetRealType2(val reflect.Type) reflect.Type {
	if val != nil {
		if val.Kind() == reflect.Pointer {
			val = val.Elem()
		}
	}
	return val
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
