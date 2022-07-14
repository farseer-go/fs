package mapper

import (
	"github.com/devfeel/mapper"
)

// MapperArray 转换结构数组
func MapperArray[T any](fromSlice any) T {
	var toSlice T
	_ = mapper.MapperSlice(fromSlice, &toSlice)
	return toSlice
}

// AutoMapper 转换
func AutoMapper[T any](fromObj any) T {
	var toObj T
	_ = mapper.MapperSlice(fromObj, &toObj)
	return toObj
}
