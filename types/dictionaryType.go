package types

import "reflect"

// DictionaryNew 动态创建一个新的Dictionary
func DictionaryNew(lstType reflect.Type) reflect.Value {
	dicValue := reflect.New(lstType)
	dicValue.MethodByName("New").Call(nil)
	return dicValue
}

// DictionaryAddMap 动态添加元素
func DictionaryAddMap(lstValue reflect.Value, item any) {
	itemValue := reflect.ValueOf(item)
	if itemValue.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
	}
	if itemValue.Kind() == reflect.Slice {
		lstValue.MethodByName("AddMap").CallSlice([]reflect.Value{itemValue})
	} else {
		lstValue.MethodByName("AddMap").Call([]reflect.Value{itemValue})
	}
}

// GetDictionaryToMap 获取Dictionary的map元素
func GetDictionaryToMap(lstValue reflect.Value) reflect.Value {
	mapValue := lstValue.MethodByName("ToMap").Call(nil)[0]
	return mapValue
}

// GetDictionaryMapType 获取List的原始数组类型
func GetDictionaryMapType(lstType reflect.Type) reflect.Type {
	sourceField, _ := lstType.FieldByName("source")
	return sourceField.Type
}
