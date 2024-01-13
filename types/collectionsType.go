package types

import "reflect"

// GetPageList 获取collections.PageList的元素
func GetPageList(pageList any) (any, int64) {
	pageListValueOf := reflect.ValueOf(pageList)
	if _, success := IsPageList(pageListValueOf); !success {
		panic("ToPageList的入参必须是collections.PageList类型")
	}
	listValueOf := pageListValueOf.FieldByName("List")
	recordCountValueOf := pageListValueOf.FieldByName("RecordCount")
	return listValueOf.Interface(), recordCountValueOf.Interface().(int64)
}
