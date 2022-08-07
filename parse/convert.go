package parse

import (
	"reflect"
	"strings"
)

// Convert 通用的类型转换
func Convert[T any](source any, defVal T) T {
	sourceKind := reflect.TypeOf(source).Kind()
	returnKind := reflect.TypeOf(defVal).Kind()

	if sourceKind == returnKind {
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

	// bool转数字
	if isBool(sourceKind) {
		var result any = 0
		boolSource := source.(bool)
		if boolSource {
			result = 1
		}
		return result.(T)
	}

	// 字符串转...
	if isString(sourceKind) {
		strSource := source.(string)

		if isBool(returnKind) { // 字符串转bool
			var result any = strings.EqualFold(strSource, "true")
			return result.(T)
		}

		// 字符串转数字
		if isNumber(returnKind) {
			return stringToNumber(strSource, defVal, returnKind).(T)
		}
	}
	return defVal
}

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
