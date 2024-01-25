package parse

import (
	"encoding/json"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/fastReflect"
	"github.com/farseer-go/fs/types"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

var layouts = []string{"2006-01-02 15:04:05", "2006-01-02", "2006-01-02T15:04:05Z07:00"}

// ConvertValue 通用的类型转换
func ConvertValue(source any, defValType reflect.Type) reflect.Value {
	// any类型，则直接返回
	if defValType.Kind() == reflect.Interface {
		return reflect.ValueOf(source)
	}

	defVal := reflect.New(defValType).Elem().Interface()
	val := Convert(source, defVal)
	return reflect.ValueOf(val)
}

// Convert 通用的类型转换
func Convert[T any](source any, defVal T) T {
	if source == nil {
		return defVal
	}
	sourceMeta := fastReflect.ValueOf(source)
	defValMeta := fastReflect.ValueOf(defVal)

	// time不支持直接转换，因为是结构体
	if sourceMeta.HashCode == defValMeta.HashCode { // sourceMeta.Kind != reflect.Struct &&
		return source.(T)
	}

	// 数字转...
	if sourceMeta.IsNumber {
		switch defValMeta.TypeIdentity {
		// 转数字
		case "number":
			result := anyToNumber(source, sourceMeta.Kind, defVal, defValMeta.Kind)
			return result.(T)
		// 枚举类型
		case "enum":
			return toEnum[T](defValMeta.ReflectType, source)
		// 转字符串
		case "string":
			var str = NumberToString(source, sourceMeta.Kind)
			return *(*T)(unsafe.Pointer(&str))
		// 转bool
		case "bool":
			result := EqualTo1(source, sourceMeta.Kind)
			return *(*T)(unsafe.Pointer(&result))
		}
	}

	// 字符串转...
	if sourceMeta.IsString {
		var strSource string
		if sourceMeta.ReflectTypeString == "json.Number" {
			jsonNumber := source.(json.Number)
			strSource = string(jsonNumber)
		} else {
			strSource = source.(string)
		}

		switch defValMeta.TypeIdentity {
		// 转枚举类型
		case "enum":
			return toEnum[T](defValMeta.ReflectType, strSource)
		// 转数字
		case "number":
			result := StringToNumber(strSource, defVal, defValMeta.Kind)
			return result.(T)
		// 转数组
		case "sliceOrArray":
			arr := strings.Split(strSource, ",")
			// 字符串数组，则直接转
			if defValMeta.ItemMeta.ReflectType.Kind() == reflect.String {
				return *(*T)(unsafe.Pointer(&arr))
			}
			// 创建数组（耗时65ns）
			slice := reflect.MakeSlice(defValMeta.ReflectType, len(arr), len(arr))
			slicePtr := slice.Pointer()
			for i := 0; i < len(arr); i++ {
				// 找到当前索引位置的内存地址。起始位置 + 每个元素占用的字节大小 ，得到第N个索引的内存起始位置
				itemPtr := unsafe.Pointer(slicePtr + uintptr(i)*defValMeta.ItemMeta.Size)
				// 3条数据的情况下，耗时228ns
				newItemVal := Convert(arr[i], defValMeta.ItemMeta.ZeroValue)
				// 设置值
				fastReflect.SetValue(itemPtr, newItemVal, defValMeta.ItemMeta)
			}
			return slice.Interface().(T)
		case "bool":
			var result any = strings.EqualFold(strSource, "true")
			return result.(T)
		// 转dateTime 转time.Time
		case "time", "dateTime":
			switch len(strSource) {
			case 19:
				if parse, err := time.ParseInLocation(layouts[0], strSource, time.Local); err == nil {
					return *(*T)(unsafe.Pointer(&parse))
				}
			case 10:
				if parse, err := time.ParseInLocation(layouts[1], strSource, time.Local); err == nil {
					return *(*T)(unsafe.Pointer(&parse))
				}
			case 25:
				if parse, err := time.ParseInLocation(layouts[2], strSource, time.Local); err == nil {
					return *(*T)(unsafe.Pointer(&parse))
				}
			}

			for _, layout := range layouts {
				if parse, err := time.ParseInLocation(layout, strSource, time.Local); err == nil {
					return *(*T)(unsafe.Pointer(&parse))
				}
			}
		// list类型
		case "list":
			lstReflectValue := types.ListNew(defValMeta.ReflectType)
			arr := strings.Split(strSource, ",")
			for i := 0; i < len(arr); i++ {
				types.ListAdd(lstReflectValue, ConvertValue(arr[i], defValMeta.ItemMeta.ReflectType).Interface())
			}
			val := lstReflectValue.Elem().Interface()
			return *(*T)(unsafe.Pointer(&val))
		}
	}

	// bool转...
	if sourceMeta.IsBool {
		boolSource := source.(bool)
		switch defValMeta.TypeIdentity {
		// 转数字
		case "number":
			if boolSource {
				return any(1).(T)
			}
			return any(0).(T)
		case "string":
			if boolSource {
				return any("true").(T)
			} else {
				return any("false").(T)
			}
		}
		return defVal
	}

	// time.Time转...
	if sourceMeta.IsTime {
		switch defValMeta.TypeIdentity {
		// 转time.Time
		case "time":
			return *(*T)(sourceMeta.PointerValue)
		// 转DateTime
		case "dateTime":
			var dt = dateTime.New(source.(time.Time))
			return *(*T)(unsafe.Pointer(&dt))
		// 转string
		case "string":
			var str = source.(time.Time).Format("2006-01-02 15:04:05")
			return *(*T)(unsafe.Pointer(&str))
		}
	}

	// DateTime转...
	if sourceMeta.IsDateTime {
		switch defValMeta.TypeIdentity {
		// 转time.Time
		case "time":
			var t = source.(dateTime.DateTime).ToTime()
			return *(*T)(unsafe.Pointer(&t))
		// 转DateTime
		case "dateTime":
			return source.(T)
		// 转string
		case "string":
			var str = source.(dateTime.DateTime).ToString("yyyy-MM-dd HH:mm:ss")
			return *(*T)(unsafe.Pointer(&str))
		}
	}

	// 切片转切片
	if sourceMeta.Type == fastReflect.Slice && defValMeta.Type == fastReflect.Slice {
		arr := reflect.MakeSlice(defValMeta.ReflectType, 0, 0)
		arrSource := reflect.ValueOf(source)
		for i := 0; i < arrSource.Len(); i++ {
			item := arrSource.Index(i)
			destVal := ConvertValue(item.Interface(), defValMeta.ItemMeta.ReflectType)
			arr = reflect.Append(arr, destVal)
		}
		return arr.Interface().(T)
	}

	return defVal
}

// 转枚举
func toEnum[T any](tType reflect.Type, result any) T {
	returnTypeNew := reflect.New(tType).Elem()
	switch tType.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		returnTypeNew.SetInt(ToInt64(result))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		returnTypeNew.SetUint(ToUInt64(result))
	case reflect.Float32, reflect.Float64:
		returnTypeNew.SetFloat(ToFloat64(result))
	default:
		panic("不支持的类型转枚举")
	}
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
	default:
		return false
	}
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
