package types

import (
	"reflect"
	"strings"
)

// IsSlice 是否为切片类型
func IsSlice(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, realType.Kind() == reflect.Slice
}

// IsMap 是否为Map类型
func IsMap(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, realType.Kind() == reflect.Map
}

// IsList 判断类型是否为List
func IsList(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), "collections.List[")
}

// IsDictionary 判断类型是否为Dictionary
func IsDictionary(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), "collections.Dictionary[")
}

// IsPageList 判断类型是否为PageList
func IsPageList(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), "collections.PageList[")
}

// IsCollections 是否为集合
func IsCollections(ty reflect.Type) bool {
	return strings.HasPrefix(ty.String(), "collections.")
}

// IsStruct 是否为Struct
func IsStruct(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, realType.Kind() == reflect.Struct
}

// IsEsIndexSet 判断类型是否为ES的IndexSet类型
func IsEsIndexSet(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), "elasticSearch.IndexSet[")
}

// IsDataTableSet 判断类型是否为Data的TableSet类型
func IsDataTableSet(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), "data.TableSet[")
}
