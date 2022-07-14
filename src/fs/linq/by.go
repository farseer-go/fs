package linq

import (
	"reflect"
)

type linqBy[T comparable] struct {
	// source array
	source *[]T
	// array type
	arrayType reflect.Type
	// element type
	elementType reflect.Type
	// value type
	value reflect.Value
}

func By[T comparable](source []T) linqBy[T] {
	return linqBy[T]{
		source: &source,
	}
}

// Contains 查找数组是否包含某元素
func (receiver linqBy[T]) Contains(t T) bool {
	for _, item := range *receiver.source {
		if item == t {
			return true
		}
	}
	return false
}

// Remove 移除指定值的元素
func (receiver linqBy[T]) Remove(val T) []T {
	var lst []T
	for _, item := range *receiver.source {
		if item != val {
			lst = append(lst, item)
		}
	}
	receiver.source = &lst
	return lst
}
