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
	if _, isExists := getCache("dic.AddMap"); !isExists {
		method, _ := lstValue.Type().MethodByName("AddMap")
		setCache("dic.AddMap", []int{method.Index})
	}

	itemValue := reflect.ValueOf(item)
	if itemValue.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
	}
	if itemValue.Kind() == reflect.Slice {
		lstValue.Method(getCacheVal("dic.AddMap")[0]).CallSlice([]reflect.Value{itemValue})
	} else {
		lstValue.Method(getCacheVal("dic.AddMap")[0]).Call([]reflect.Value{itemValue})
	}
}

// GetDictionaryToMap 获取Dictionary的map元素
func GetDictionaryToMap(lstValue reflect.Value) reflect.Value {
	if _, isExists := getCache("dic.ToMap"); !isExists {
		method, _ := lstValue.Type().MethodByName("ToMap")
		setCache("dic.ToMap", []int{method.Index})
	}
	return lstValue.Method(getCacheVal("dic.ToMap")[0]).Call(nil)[0]
}

// GetDictionaryMapType 获取Dictionary的原始数组类型
func GetDictionaryMapType(lstType reflect.Type) reflect.Type {
	innerMapType := lstType.Field(0).Type
	if _, isExists := getCache("dic.source"); !isExists {
		field, _ := innerMapType.FieldByName("source2")
		setCache("dic.source", field.Index)
	}
	if len(getCacheVal("dic.source")) == 1 {
		return innerMapType.Field(getCacheVal("dic.source")[0]).Type
	}
	return innerMapType.FieldByIndex(getCacheVal("dic.source")).Type
}
