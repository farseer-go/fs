package linq

import "reflect"

type linqDictionary[TK comparable, TV any] struct {
	// source array
	source map[TK]TV
	// array type
	arrayType reflect.Type
	// element type
	elementType reflect.Type
	// value type
	value reflect.Value
}

// Dictionary 针对字典的操作
func Dictionary[TK comparable, TV any](source map[TK]TV) linqDictionary[TK, TV] {
	return linqDictionary[TK, TV]{
		source: source,
	}
}

// ExistsKey 是否存在KEY
func (receiver linqDictionary[TK, TV]) ExistsKey(key TK) bool {
	_, exists := receiver.source[key]
	return exists
}
