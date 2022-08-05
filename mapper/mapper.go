package mapper

import (
	"github.com/devfeel/mapper"
	"github.com/farseernet/farseer.go/core"
)

// Array 数组转换
func Array[T any](fromSlice any) []T {
	var toSlice []T
	_ = mapper.MapperSlice(fromSlice, &toSlice)
	return toSlice
}

// Single 单个转换
func Single[T any](fromObj any) T {
	var toObj T
	_ = mapper.MapperSlice(fromObj, &toObj)
	return toObj
}

// PageList 转换成core.PageList
func PageList[TData any](fromObj any, recordCount int64) core.PageList[TData] {
	lst := Array[TData](fromObj)
	return core.NewPageList(lst, recordCount)
}
