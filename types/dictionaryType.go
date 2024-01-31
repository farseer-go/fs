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
	if _, isExists := cache["dic.AddMap"]; !isExists {
		method, _ := lstValue.Type().MethodByName("AddMap")
		cache["dic.AddMap"] = []int{method.Index}
	}

	itemValue := reflect.ValueOf(item)
	if itemValue.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
	}
	if itemValue.Kind() == reflect.Slice {
		lstValue.Method(cache["dic.AddMap"][0]).CallSlice([]reflect.Value{itemValue})
	} else {
		lstValue.Method(cache["dic.AddMap"][0]).Call([]reflect.Value{itemValue})
	}
}

// GetDictionaryToMap 获取Dictionary的map元素
func GetDictionaryToMap(lstValue reflect.Value) reflect.Value {
	if _, isExists := cache["dic.ToMap"]; !isExists {
		method, _ := lstValue.Type().MethodByName("ToMap")
		cache["dic.ToMap"] = []int{method.Index}
	}
	return lstValue.Method(cache["dic.ToMap"][0]).Call(nil)[0]
}

// GetDictionaryMapType 获取List的原始数组类型
func GetDictionaryMapType(lstType reflect.Type) reflect.Type {
	innerMapType := lstType.Field(0).Type

	if _, isExists := cache["dic.source"]; !isExists {
		for i := 0; i < innerMapType.NumField(); i++ {
			field, _ := innerMapType.FieldByName("source")
			cache["dic.source"] = field.Index
		}
	}
	if len(cache["dic.source"]) == 1 {
		return innerMapType.Field(cache["dic.source"][0]).Type
	}
	return innerMapType.FieldByIndex(cache["dic.source"]).Type
}
