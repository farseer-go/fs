package parse

import (
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/types"
	"reflect"
	"strings"
	"time"
)

// ConvertValue 通用的类型转换
func ConvertValue(source any, defValType reflect.Type) reflect.Value {
	defVal := reflect.New(defValType).Elem().Interface()
	val := Convert(source, defVal)
	return reflect.ValueOf(val)
}

// Convert 通用的类型转换
func Convert[T any](source any, defVal T) T {
	if source == nil {
		return defVal
	}

	sourceType := reflect.TypeOf(source)
	sourceKind := sourceType.Kind()
	defValType := reflect.TypeOf(defVal)
	returnKind := defValType.Kind()

	if sourceKind != reflect.Struct && sourceKind == returnKind {
		return source.(T)
	}

	// 数字转...
	if isNumber(sourceKind) {
		// 数字转数字
		if isNumber(returnKind) {
			return anyToNumber(source, sourceKind, defVal, returnKind).(T)
		}

		// 数字转bool
		if isBool(returnKind) {
			var result any = equalTo1(source, sourceKind)
			return result.(T)
		}

		// 数字转字符串
		if isString(returnKind) {
			return numberToString(source, defVal, sourceKind).(T)
		}
	}

	// bool转...
	if isBool(sourceKind) {
		boolSource := source.(bool)
		var result any

		// 转数字
		if isNumber(returnKind) {
			result = 0
			if boolSource {
				result = 1
			}
			return result.(T)
		}

		if isString(returnKind) {
			if boolSource {
				result = "true"
			} else {
				result = "false"
			}
			return result.(T)
		}
		return defVal
	}

	// 字符串转...
	if isString(sourceKind) {
		strSource := source.(string)
		// bool
		if isBool(returnKind) {
			var result any = strings.EqualFold(strSource, "true")
			return result.(T)
		}

		// 数字
		if isNumber(returnKind) {
			return stringToNumber(strSource, defVal, returnKind).(T)
		}

		// 数组
		if isArray(returnKind) {
			arr := strings.Split(strSource, ",")
			itemType := defValType.Elem()
			// 字符串数组，则直接转
			if itemType.Kind() == reflect.String {
				return any(arr).(T)
			}

			// 非字符串数组，则要动态
			slice := reflect.MakeSlice(defValType, 0, len(arr))
			for i := 0; i < len(arr); i++ {
				slice = reflect.Append(slice, ConvertValue(arr[i], itemType))
			}
			return slice.Interface().(T)
		}

		// list类型
		if isList(defValType) {
			lstReflectValue := types.ListNew(defValType)
			lstItemType := types.GetListItemType(defValType)
			arr := strings.Split(strSource, ",")
			for i := 0; i < len(arr); i++ {
				types.ListAdd(&lstReflectValue, ConvertValue(arr[i], lstItemType).Interface())
			}
			return lstReflectValue.Elem().Interface().(T)
		}

		// 转time.Time
		layouts := []string{time.DateTime, time.DateOnly, time.RFC3339}
		if types.IsTime(defValType) {
			for _, layout := range layouts {
				parse, err := time.ParseInLocation(layout, source.(string), time.Local)
				if err == nil {
					return any(parse).(T)
				}
			}
		}

		// 转DateTime
		if types.IsDateTime(defValType) {
			for _, layout := range layouts {
				parse, err := time.ParseInLocation(layout, source.(string), time.Local)
				if err == nil {
					return any(dateTime.New(parse)).(T)
				}
			}
		}
	}

	// time.Time转...
	if types.IsTime(sourceType) {
		// 转time.Time
		if types.IsTime(defValType) {
			return source.(T)
		}
		// 转DateTime
		if types.IsDateTime(defValType) {
			var dt any = dateTime.New(source.(time.Time))
			return dt.(T)
		}
	}

	// DateTime转...
	if types.IsDateTime(sourceType) {
		// 转time.Time
		if types.IsTime(defValType) {
			var t any = source.(dateTime.DateTime).ToTime()
			return t.(T)
		}

		// 转DateTime
		if types.IsDateTime(defValType) {
			return source.(T)
		}
	}
	return defVal
}

// ToInt 转换成int类型
func ToInt(source any) int { return Convert(source, 0) }

// ToInt64 转换成int64类型
func ToInt64(source any) int64 { return Convert(source, int64(0)) }

// ToString 转换成string类型
func ToString(source any) string { return Convert(source, "") }

// ToBool 转换成bool类型
func ToBool(source any) bool { return Convert(source, false) }

// ToTime 转换成time.Time类型
func ToTime(source any) time.Time { return Convert(source, time.Time{}) }

// 数字类型
func isNumber(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// 布尔值类型
func isBool(kind reflect.Kind) bool {
	return kind == reflect.Bool
}

// 布尔值类型
func isString(kind reflect.Kind) bool {
	return kind == reflect.String
}

// 数组
func isArray(kind reflect.Kind) bool {
	return kind == reflect.Array || kind == reflect.Slice
}

// isList 判断类型是否为List
func isList(realType reflect.Type) bool {
	return strings.HasPrefix(realType.String(), "collections.List[")
}
