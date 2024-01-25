package fastReflect

import (
	"github.com/farseer-go/fs/types"
	"reflect"
	"strings"
)

// TypeMeta 类型元数据
type TypeMeta struct {
	Name               string              // 字段名称
	ReflectType        reflect.Type        // 字段类型
	ReflectTypeString  string              // 类型
	ReflectStructField reflect.StructField // 字段类型
	Type               FieldType           // 集合类型
	IsAnonymous        bool                // 是否为内嵌类型
	IsExported         bool                // 是否为可导出类型
	IsIgnore           bool                // 是否为忽略字段
	IsAddr             bool                // 原类型是否带指针
	MapKey             reflect.Value       // map key
	NumField           int                 // 结构体的字段数量
	ItemMeta           *TypeMeta           // Item元素的Type
	SliceType          reflect.Type        // ItemType转成切片类型
	ZeroValue          any                 // 零值时的值
	Kind               reflect.Kind
	IsNumber           bool    // 是否为数字
	IsEmum             bool    // 是否枚举
	IsString           bool    // 是否为字符串
	IsBool             bool    // 是否bool
	IsTime             bool    // 是否time.Time
	IsDateTime         bool    // 是否dateTime.DateTime
	IsSliceOrArray     bool    // 是否切片或数组类型
	IsStruct           bool    // 是否结构体
	HashCode           uint32  // 每个类型的HashCode都是唯一的
	Size               uintptr // 内存占用大小
	TypeIdentity       string  // 类型标识
}

func typeOf(inf *EmptyInterface, reflectType reflect.Type) *TypeMeta {
	kind := reflect.Kind(inf.Typ.kind)

	tm := &TypeMeta{
		ReflectType: reflectType,
		IsAddr:      kind == reflect.Pointer,
		Kind:        kind,

		// 作为字段时才能获取
		IsAnonymous: false,
		IsExported:  false,
		IsIgnore:    false,
		HashCode:    inf.Typ.hash,
		Size:        inf.Typ.size,

		ZeroValue: reflect.New(reflectType).Elem().Interface(),
	}

	// 解析类型
	tm.parseType()
	return tm
}

func (receiver *TypeMeta) parseType() {
	// 指针类型，需要取出指针指向的类型
	if receiver.Kind == reflect.Pointer {
		receiver.ReflectType = receiver.ReflectType.Elem()
		receiver.Kind = receiver.ReflectType.Kind()
	}

	// 取真实的类型
	if receiver.Kind == reflect.Interface {
		receiver.ReflectType = receiver.ReflectType.Elem()
		receiver.Kind = receiver.ReflectType.Kind()
	}

	receiver.ReflectTypeString = receiver.ReflectType.String()

	switch receiver.Kind {
	case reflect.Slice:
		receiver.Type = Slice
		itemType := receiver.ReflectType.Elem()
		itemVal := reflect.New(itemType).Elem().Interface()
		receiver.ItemMeta = ValueOf(itemVal).TypeMeta
		receiver.SliceType = reflect.SliceOf(receiver.ItemMeta.ReflectType)
		receiver.IsSliceOrArray = true
	case reflect.Array:
		receiver.Type = Array
		itemType := receiver.ReflectType.Elem()
		itemVal := reflect.New(itemType).Elem().Interface()
		receiver.ItemMeta = ValueOf(itemVal).TypeMeta
		receiver.IsSliceOrArray = true
	case reflect.Map:
		receiver.Type = Map
	case reflect.Chan:
		receiver.Type = Chan
	case reflect.Func:
		receiver.Type = Func
	case reflect.Invalid:
		receiver.Type = Invalid
	case reflect.Interface:
		receiver.Type = Interface
	default:
		// 基础类型
		if types.IsGoBasicType(receiver.ReflectType) {
			receiver.Type = GoBasicType

			// 基础类型使用
			receiver.IsNumber = isNumber(receiver.Kind)
			receiver.IsEmum = receiver.IsNumber && strings.Contains(receiver.ReflectTypeString, ".")
			receiver.IsString = receiver.Kind == reflect.String
			receiver.IsBool = receiver.Kind == reflect.Bool
			receiver.IsTime = receiver.ReflectTypeString == TimeString
			receiver.IsDateTime = receiver.ReflectTypeString == DatetimeString
			break
		}

		// List类型
		if _, isTrue := types.IsListByType(receiver.ReflectType); isTrue {
			receiver.Type = List
			itemType := types.GetListItemType(receiver.ReflectType)
			itemVal := reflect.New(itemType).Elem().Interface()
			receiver.ItemMeta = ValueOf(itemVal).TypeMeta
			receiver.SliceType = reflect.SliceOf(receiver.ItemMeta.ReflectType)
			break
		}

		// Dictionary类型
		if isTrue := types.IsDictionaryByType(receiver.ReflectType); isTrue {
			receiver.Type = Dic
			break
		}

		// PageList类型
		if isTrue := types.IsPageListByType(receiver.ReflectType); isTrue {
			receiver.Type = PageList
			break
		}

		// 自定义集合类型
		numField := receiver.ReflectType.NumField()
		if numField > 0 && receiver.ReflectType.Field(0).PkgPath == CollectionsTypeString {
			receiver.Type = CustomList
			itemType := types.GetListItemType(receiver.ReflectType)
			itemVal := reflect.New(itemType).Elem().Interface()
			receiver.ItemMeta = ValueOf(itemVal).TypeMeta
			receiver.SliceType = reflect.SliceOf(receiver.ItemMeta.ReflectType)
			break
		}

		// 结构体
		if types.IsStruct(receiver.ReflectType) {
			receiver.Type = Struct
			receiver.NumField = numField
			receiver.IsStruct = true
			break
		}
		receiver.Type = Unknown
	}

	if receiver.IsEmum {
		receiver.TypeIdentity = "enum"
	} else if receiver.IsNumber {
		receiver.TypeIdentity = "number"
	} else if receiver.IsString {
		receiver.TypeIdentity = "string"
	} else if receiver.IsBool {
		receiver.TypeIdentity = "bool"
	} else if receiver.IsTime {
		receiver.TypeIdentity = "time"
	} else if receiver.IsDateTime {
		receiver.TypeIdentity = "dateTime"
	} else if receiver.IsSliceOrArray {
		receiver.TypeIdentity = "sliceOrArray"
	} else if receiver.Type == List {
		receiver.TypeIdentity = "list"
	}
}
