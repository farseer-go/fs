package types

import "reflect"

// GetPageList 获取collections.PageList的元素
func GetPageList(pageList any) (any, int64) {
	pageListValueOf := reflect.ValueOf(pageList)
	if _, isExists := getCache("pageList.List"); !isExists {
		t := pageListValueOf.Type()
		method, _ := t.FieldByName("List")
		setCache("pageList.List", method.Index)

		method, _ = t.FieldByName("RecordCount")
		setCache("pageList.RecordCount", method.Index)
	}

	if _, success := IsPageList(pageListValueOf); !success {
		panic("ToPageList的入参必须是collections.PageList类型")
	}
	var listValueOf reflect.Value
	if len(getCacheVal("pageList.List")) == 1 {
		listValueOf = pageListValueOf.Field(getCacheVal("pageList.List")[0])
	} else {
		listValueOf = pageListValueOf.FieldByIndex(getCacheVal("pageList.List"))
	}

	var recordCountValueOf reflect.Value
	if len(getCacheVal("pageList.RecordCount")) == 1 {
		recordCountValueOf = pageListValueOf.Field(getCacheVal("pageList.RecordCount")[0])
	} else {
		recordCountValueOf = pageListValueOf.FieldByIndex(getCacheVal("pageList.RecordCount"))
	}

	return listValueOf.Interface(), recordCountValueOf.Interface().(int64)
}
