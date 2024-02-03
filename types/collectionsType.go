package types

import "reflect"

// GetPageList 获取collections.PageList的元素
func GetPageList(pageList any) (any, int64) {

	pageListValueOf := reflect.ValueOf(pageList)
	if _, isExists := Cache["pageList.List"]; !isExists {
		t := pageListValueOf.Type()
		method, _ := t.FieldByName("List")
		Cache["pageList.List"] = method.Index

		method, _ = t.FieldByName("RecordCount")
		Cache["pageList.RecordCount"] = method.Index
	}

	if _, success := IsPageList(pageListValueOf); !success {
		panic("ToPageList的入参必须是collections.PageList类型")
	}
	var listValueOf reflect.Value
	if len(Cache["pageList.List"]) == 1 {
		listValueOf = pageListValueOf.Field(Cache["pageList.List"][0])
	} else {
		listValueOf = pageListValueOf.FieldByIndex(Cache["pageList.List"])
	}

	var recordCountValueOf reflect.Value
	if len(Cache["pageList.RecordCount"]) == 1 {
		recordCountValueOf = pageListValueOf.Field(Cache["pageList.RecordCount"][0])
	} else {
		recordCountValueOf = pageListValueOf.FieldByIndex(Cache["pageList.RecordCount"])
	}

	return listValueOf.Interface(), recordCountValueOf.Interface().(int64)
}
