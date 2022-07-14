package linq

import (
	"reflect"
)

type linqForm[T comparable] struct {
	// source array
	source *[]T
	// array type
	arrayType reflect.Type
	// element type
	elementType reflect.Type
	// value type
	value reflect.Value
}

func From[T comparable](source []T) linqForm[T] {
	return linqForm[T]{
		source: &source,
	}
}

type whereFunc[T any] func(item T) bool

// Where 对数据进行筛选
func (receiver linqForm[T]) Where(fn whereFunc[T]) linqForm[T] {
	var lst []T
	for _, item := range *receiver.source {
		if fn(item) {
			lst = append(lst, item)
		}
	}
	receiver.source = &lst
	return receiver
}

// Find 查找符合条件的元素
func (receiver linqForm[T]) Find(fn whereFunc[T]) T {
	for _, item := range *receiver.source {
		if fn(item) {
			return item
		}
	}
	var t T
	return t
}

// FindAll 查找符合条件的元素列表
func (receiver linqForm[T]) FindAll(fn whereFunc[T]) []T {
	var lst []T
	for _, item := range *receiver.source {
		if fn(item) {
			lst = append(lst, item)
		}
	}
	return lst
}

// Remove 移除为nil的元素
func (receiver linqForm[T]) Remove(val T) []T {
	var lst []T
	for _, item := range *receiver.source {
		if item != val {
			lst = append(lst, item)
		}
	}
	receiver.source = &lst
	return lst
}
