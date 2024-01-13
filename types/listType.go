package types

import "reflect"

// ListNew 动态创建一个新的List
func ListNew(lstType reflect.Type) reflect.Value {
	lstValue := reflect.New(lstType)
	lstValue.MethodByName("New").Call(nil)
	return lstValue
}

// ListAdd 动态添加元素
func ListAdd(lstValue reflect.Value, item any) {
	itemValue := reflect.ValueOf(item)
	if itemValue.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
	}
	if itemValue.Kind() == reflect.Slice {
		lstValue.MethodByName("Add").CallSlice([]reflect.Value{itemValue})
	} else {
		lstValue.MethodByName("Add").Call([]reflect.Value{itemValue})
	}
}

// GetListItemArrayType 获取List的原始数组类型
func GetListItemArrayType(lstType reflect.Type) reflect.Type {
	sourceField, _ := lstType.FieldByName("source")
	return sourceField.Type.Elem()
}

// GetListItemType 获取List的元素Type
func GetListItemType(lstType reflect.Type) reflect.Type {
	sourceField, _ := lstType.FieldByName("source")
	return sourceField.Type.Elem().Elem()
}

func GetListToArray(lstValue reflect.Value) []any {
	arrValue := lstValue.MethodByName("ToArray").Call(nil)[0]
	var items []any
	for i := 0; i < arrValue.Len(); i++ {
		item := arrValue.Index(i).Interface()
		items = append(items, item)
	}
	return items
}
