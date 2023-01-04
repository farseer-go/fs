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
func IsStruct(ty reflect.Type) bool {
	realType := GetRealType2(ty)
	if realType.Kind() == reflect.Struct {
		if IsCollections(realType) || realType.String() == "time.Time" {
			return false
		}
		return true
	}
	return false
}

// IsGoBasicType 是否为Go内置基础类型
func IsGoBasicType(ty reflect.Type) bool {
	realType := GetRealType2(ty)
	if realType != nil {
		switch realType.Kind() {
		case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
			return true
		}
		switch realType.String() {
		case "time.Time":
			return true
		}
	}
	return false
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

// IsDtoModel 当只有一个参数，且非集合类型，又是结构类型时，判断为DTO
func IsDtoModel(lst []reflect.Type) bool {
	if len(lst) != 1 {
		return false
	}

	return !IsCollections(lst[0]) && !IsGoBasicType(lst[0]) && lst[0].Kind() == reflect.Struct
}

// IsTime 是否为time.Time类型
func IsTime(ty reflect.Type) bool {
	return ty.String() == "time.Time"
}

// IsDateTime 是否为DateTime类型
func IsDateTime(ty reflect.Type) bool {
	return ty.String() == "dateTime.DateTime"
}
