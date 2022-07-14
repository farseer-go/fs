package linq

import (
	"reflect"
)

type linqForm[T any] struct {
	// source array
	source *[]T
	// array type
	arrayType reflect.Type
	// element type
	elementType reflect.Type
	// value type
	value reflect.Value
}

func From[T any](source []T) linqForm[T] {
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

// RemoveAll 移除条件=true的元素
func (receiver linqForm[T]) RemoveAll(fn whereFunc[T]) []T {
	var lst []T
	for _, item := range *receiver.source {
		if !fn(item) {
			lst = append(lst, item)
		}
	}
	receiver.source = &lst
	return lst
}
