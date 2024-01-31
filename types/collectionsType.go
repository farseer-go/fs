package types

import "reflect"

// GetPageList 获取collections.PageList的元素
func GetPageList(pageList any) (any, int64) {

	pageListValueOf := reflect.ValueOf(pageList)
	if _, isExists := cache["pageList.List"]; !isExists {
		t := pageListValueOf.Type()
		method, _ := t.FieldByName("List")
		cache["pageList.List"] = method.Index

		method, _ = t.FieldByName("RecordCount")
		cache["pageList.RecordCount"] = method.Index
	}

	if _, success := IsPageList(pageListValueOf); !success {
		panic("ToPageList的入参必须是collections.PageList类型")
	}
	var listValueOf reflect.Value
	if len(cache["pageList.List"]) == 1 {
		listValueOf = pageListValueOf.Field(cache["pageList.List"][0])
	} else {
		listValueOf = pageListValueOf.FieldByIndex(cache["pageList.List"])
	}

	var recordCountValueOf reflect.Value
	if len(cache["pageList.RecordCount"]) == 1 {
		recordCountValueOf = pageListValueOf.Field(cache["pageList.RecordCount"][0])
	} else {
		recordCountValueOf = pageListValueOf.FieldByIndex(cache["pageList.RecordCount"])
	}

	return listValueOf.Interface(), recordCountValueOf.Interface().(int64)
}
