package linq

import "reflect"

type linqMap[TK comparable, TV any] struct {
	// source array
	source map[TK]TV
	// array type
	arrayType reflect.Type
	// element type
	elementType reflect.Type
	// value type
	value reflect.Value
}

// Map 针对字典的操作
func Map[TK comparable, TV any](source map[TK]TV) linqMap[TK, TV] {
	return linqMap[TK, TV]{
		source: source,
	}
}

// ExistsKey 是否存在KEY
func (receiver linqMap[TK, TV]) ExistsKey(key TK) bool {
	_, exists := receiver.source[key]
	return exists
}

// ToValue 将map的value转成数组[]value
func (receiver linqMap[TK, TV]) ToValue() []TV {
	var arr []TV
	for _, v := range receiver.source {
		arr = append(arr, v)
	}
	return arr
}

// ToKey 将map的key转成数组[]key
func (receiver linqMap[TK, TV]) ToKey() []TK {
	var arr []TK
	for k := range receiver.source {
		arr = append(arr, k)
	}
	return arr
}
