package types

import (
	"reflect"
	"strings"
)

const (
	ListTypeString              = "collections.List["
	ListAnyTypeString           = "collections.ListAny"
	DictionaryTypeString        = "collections.Dictionary["
	PageListTypeString          = "collections.PageList["
	CollectionsTypeString       = "github.com/farseer-go/collections"
	CollectionsPrefixTypeString = "collections."
	DatetimeString              = "dateTime.DateTime"
	TimeString                  = "time.Time"
	DecimalString               = "decimal.Decimal"
	DomainSetString             = "data.DomainSet["
	TableSetString              = "data.TableSet["
	IndexSetString              = "elasticSearch.IndexSet["
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
	realTypeString := realType.String()
	return realType, strings.HasPrefix(realTypeString, ListTypeString) || realTypeString == ListAnyTypeString
}

// IsListByType 判断类型是否为List
func IsListByType(realType reflect.Type) (reflect.Type, bool) {
	realTypeString := realType.String()
	return realType, strings.HasPrefix(realTypeString, ListTypeString) || realTypeString == ListAnyTypeString
}

// IsDictionary 判断类型是否为Dictionary
func IsDictionary(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), DictionaryTypeString)
}

// IsDictionaryByType 判断类型是否为Dictionary
func IsDictionaryByType(realType reflect.Type) bool {
	return strings.HasPrefix(realType.String(), DictionaryTypeString)
}

// IsPageList 判断类型是否为PageList
func IsPageList(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), PageListTypeString)
}

// IsPageListByType 判断类型是否为PageList
func IsPageListByType(realType reflect.Type) bool {
	return strings.HasPrefix(realType.String(), PageListTypeString)
}

// IsCollections 是否为集合
func IsCollections(ty reflect.Type) bool {
	return strings.HasPrefix(ty.String(), CollectionsPrefixTypeString)
}

// IsStruct 是否为Struct
func IsStruct(ty reflect.Type) bool {
	realType := GetRealType2(ty)
	if realType.Kind() == reflect.Struct {
		if IsCollections(realType) || realType.String() == TimeString || realType.String() == DatetimeString || realType.String() == DecimalString {
			return false
		}
		return true
	}
	return false
}

// IsGoBasicType 是否为Go内置基础类型
func IsGoBasicType(ty reflect.Type) bool {
	realType := GetRealType2(ty)
	if realType == nil {
		return false
	}

	switch realType.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		return true
	default:
		switch realType.String() {
		case TimeString, DecimalString, DatetimeString:
			return true
		}
	}
	return false
}

// IsEsIndexSet 判断类型是否为ES的IndexSet类型
func IsEsIndexSet(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), IndexSetString)
}

// IsDataTableSet 判断类型是否为Data的TableSet类型
func IsDataTableSet(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), TableSetString)
}

// IsDataDomainSet 判断类型是否为Data的DomainSet类型
func IsDataDomainSet(val reflect.Value) (reflect.Type, bool) {
	realType := GetRealType(val)
	return realType, strings.HasPrefix(realType.String(), DomainSetString)
}

// IsDtoModelIgnoreInterface 当第一个模型为struct，其它类型为interface时，判断为DTO
func IsDtoModelIgnoreInterface(lst []reflect.Type) bool {
	if len(lst) < 1 {
		return false
	}

	// 第一个参数必须为struct类型
	isDto := !IsGoBasicType(lst[0]) && (lst[0].Kind() == reflect.Struct || lst[0].Kind() == reflect.Array || lst[0].Kind() == reflect.Slice)
	if !isDto {
		return false
	}

	// 其它类型必须为interface
	for i := 1; i < len(lst); i++ {
		if lst[i].Kind() != reflect.Interface {
			return false
		}
	}
	return true
}

// IsDtoModel 当只有一个参数，且非集合类型，又是结构类型时，判断为DTO
func IsDtoModel(lst []reflect.Type) bool {
	if len(lst) != 1 {
		return false
	}

	return !IsCollections(lst[0]) && !IsGoBasicType(lst[0]) && lst[0].Kind() == reflect.Struct
}

// IsNil 判断值是否为nil
func IsNil(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return val.IsNil()
	default:
		return false
	}
}
