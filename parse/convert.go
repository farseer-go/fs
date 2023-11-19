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
			result := anyToNumber(source, sourceKind, defVal, returnKind)
			// 这是一个枚举类型
			if strings.Contains(defValType.String(), ".") && returnKind == reflect.Int {
				return toEnum[T](defValType, result.(int))
			}
			return result.(T)
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
			result := stringToNumber(strSource, defVal, returnKind)
			// 这是一个枚举类型
			if strings.Contains(defValType.String(), ".") && returnKind == reflect.Int {
				return toEnum[T](defValType, result.(int))
			}
			return result.(T)
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
				types.ListAdd(lstReflectValue, ConvertValue(arr[i], lstItemType).Interface())
			}
			return lstReflectValue.Elem().Interface().(T)
		}

		// 转time.Time
		layouts := []string{"2006-01-02 15:04:05", "2006-01-02", "2006-01-02T15:04:05Z07:00"}
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

func toEnum[T any](tType reflect.Type, result int) T {
	returnTypeNew := reflect.New(tType).Elem()
	returnTypeNew.SetInt(int64(result))
	return returnTypeNew.Interface().(T)
}

// ToUInt 转换成int64类型
func ToUInt(source any) uint { return Convert(source, uint(0)) }

// ToUInt8 转换成uint8类型
func ToUInt8(source any) uint8 { return Convert(source, uint8(0)) }

// ToUInt16 转换成uint16类型
func ToUInt16(source any) uint16 { return Convert(source, uint16(0)) }

// ToUInt32 转换成uint32类型
func ToUInt32(source any) uint32 { return Convert(source, uint32(0)) }

// ToUInt64 转换成uint64类型
func ToUInt64(source any) uint64 { return Convert(source, uint64(0)) }

// ToInt 转换成int类型
func ToInt(source any) int { return Convert(source, 0) }

// ToInt8 转换成int8类型
func ToInt8(source any) int8 { return Convert(source, int8(0)) }

// ToInt16 转换成int16类型
func ToInt16(source any) int16 { return Convert(source, int16(0)) }

// ToInt32 转换成int32类型
func ToInt32(source any) int32 { return Convert(source, int32(0)) }

// ToInt64 转换成int64类型
func ToInt64(source any) int64 { return Convert(source, int64(0)) }

// ToFloat32 转换成float32类型
func ToFloat32(source any) float32 { return Convert(source, float32(0)) }

// ToFloat64 转换成float64类型
func ToFloat64(source any) float64 { return Convert(source, float64(0)) }

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
