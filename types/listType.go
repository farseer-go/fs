package types

import "reflect"

// string key=自定义标识
// value int=field or method对应的索引
var Cache = make(map[string][]int)

// ListNew 动态创建一个新的List
func ListNew(lstType reflect.Type) reflect.Value {
	key := lstType.String() + ".New"
	if _, isExists := Cache[key]; !isExists {
		method, _ := reflect.New(lstType).Type().MethodByName("New")
		Cache[key] = []int{method.Index}
	}

	lstValue := reflect.New(lstType)
	lstValue.Method(Cache[key][0]).Call(nil)
	return lstValue
}

// ListAdd 动态添加元素
func ListAdd(lstValue reflect.Value, item any) {
	key := lstValue.String() + ".Add"

	if _, isExists := Cache[key]; !isExists {
		method, _ := lstValue.Type().MethodByName("Add")
		Cache[key] = []int{method.Index}
	}

	itemValue := reflect.ValueOf(item)
	if itemValue.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
	}
	if itemValue.Kind() == reflect.Slice {
		lstValue.Method(Cache[key][0]).CallSlice([]reflect.Value{itemValue})
	} else {
		lstValue.Method(Cache[key][0]).Call([]reflect.Value{itemValue})
	}
}

// GetListItemArrayType 获取List的原始数组类型
func GetListItemArrayType(lstType reflect.Type) reflect.Type {
	key := lstType.String() + ".source"
	if _, isExists := Cache[key]; !isExists {
		method, _ := lstType.FieldByName("source")
		Cache[key] = method.Index
	}
	if len(Cache[key]) == 1 {
		return lstType.Field(Cache[key][0]).Type.Elem()
	}
	return lstType.FieldByIndex(Cache[key]).Type.Elem()
}

// GetListItemType 获取List的元素Type
func GetListItemType(lstType reflect.Type) reflect.Type {
	key := lstType.String() + ".source"

	if _, isExists := Cache[key]; !isExists {
		method, _ := lstType.FieldByName("source")
		Cache[key] = method.Index
	}

	var field reflect.Type
	if len(Cache[key]) == 1 {
		field = lstType.Field(Cache[key][0]).Type
	} else {
		field = lstType.FieldByIndex(Cache[key]).Type
	}

	return field.Elem().Elem()
}

// GetListToArray 在集合中获取数据
func GetListToArray(lstValue reflect.Value) []any {
	key := lstValue.String() + ".ToArray"

	if _, isExists := Cache[key]; !isExists {
		method, _ := lstValue.Type().MethodByName("ToArray")
		Cache[key] = []int{method.Index}
	}

	arrValue := lstValue.Method(Cache[key][0]).Call(nil)[0]

	var items []any
	for i := 0; i < arrValue.Len(); i++ {
		item := arrValue.Index(i).Interface()
		items = append(items, item)
	}
	return items
}

// GetListToArrayValue 在集合中获取数据
func GetListToArrayValue(lstValue reflect.Value) reflect.Value {
	key := lstValue.String() + ".ToArray"
	if _, isExists := Cache[key]; !isExists {
		method, _ := lstValue.Type().MethodByName("ToArray")
		Cache[key] = []int{method.Index}
	}

	arrValue := lstValue.Method(Cache[key][0]).Call(nil)[0]
	return arrValue
}
