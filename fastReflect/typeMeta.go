package fastReflect

import (
	"github.com/farseer-go/fs/types"
	"reflect"
	"strings"
)

// TypeMeta 类型元数据
type TypeMeta struct {
	//Name              string              // 字段名称
	ReflectType          reflect.Type          // 字段类型
	ReflectTypeString    string                // 类型
	ReflectTypeBytes     []byte                // 类型
	TypeIdentity         string                // 类型标识
	Type                 FieldType             // 集合类型
	IsAddr               bool                  // 原类型是否带指针
	StructField          []reflect.StructField // 结构体的字段
	ExportedField        []int                 // 结构体的字段，允许导出的字段索引
	MapType              reflect.Type          // Dic的底层map类型
	keyHashCode          uint32                // map key type
	itemHashCode         uint32                // Item元素的Type or map value type
	SliceType            reflect.Type          // ItemType转成切片类型
	ZeroValue            any                   // 零值时的值
	ZeroReflectValue     reflect.Value         // 零值时的值
	ZeroReflectValueElem reflect.Value         // 零值时的值
	Kind                 reflect.Kind          // 类型
	IsNumber             bool                  // 是否为数字
	IsEmum               bool                  // 是否枚举
	IsString             bool                  // 是否为字符串
	IsBool               bool                  // 是否bool
	IsTime               bool                  // 是否time.Time
	IsDateTime           bool                  // 是否dateTime.DateTime
	IsSliceOrArray       bool                  // 是否切片或数组类型
	IsStruct             bool                  // 是否结构体
	IsMap                bool                  // 是否字典
	HashCode             uint32                // 每个类型的HashCode都是唯一的
	Size                 uintptr               // 内存占用大小
}

func typeOf(reflectType reflect.Type, inf *EmptyInterface) *TypeMeta {
	kind := reflectType.Kind()

	tm := &TypeMeta{
		ReflectType: reflectType,
		IsAddr:      kind == reflect.Pointer,
		Kind:        kind,
	}

	// 作为字段时才能获取
	if inf != nil {
		tm.HashCode = inf.Typ.hash
		tm.Size = reflectType.Size()
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

	// 设置零值
	switch receiver.Kind {
	case reflect.Slice:
		receiver.ZeroReflectValue = reflect.MakeSlice(receiver.ReflectType, 0, 0)
		receiver.ZeroReflectValueElem = receiver.ZeroReflectValue
	case reflect.Map:
		receiver.ZeroReflectValue = reflect.MakeMap(receiver.ReflectType)
		receiver.ZeroReflectValueElem = receiver.ZeroReflectValue
	default:
		receiver.ZeroReflectValue = reflect.New(receiver.ReflectType)
		receiver.ZeroReflectValueElem = reflect.New(receiver.ReflectType).Elem()
	}

	receiver.ZeroValue = receiver.ZeroReflectValueElem.Interface()

	//receiver.Name = receiver.ReflectType.Name()
	receiver.ReflectTypeString = receiver.ReflectType.String()
	receiver.ReflectTypeBytes = []byte(receiver.ReflectTypeString)

	switch receiver.Kind {
	case reflect.Slice:
		receiver.IsSliceOrArray = true
		receiver.Type = Slice
		receiver.SliceType = receiver.ReflectType

		receiver.setItemHashCode(receiver.ReflectType.Elem())
	case reflect.Array:
		receiver.IsSliceOrArray = true
		receiver.Type = Array
		receiver.setItemHashCode(receiver.ReflectType.Elem())
	case reflect.Map:
		receiver.IsMap = true
		receiver.Type = Map
		receiver.MapType = receiver.ReflectType
		// key type
		//keyType := receiver.MapType.Key()
		//keyVal := reflect.New(keyType).Elem()
		//receiver.keyHashCode = PointerOfValue(keyVal).HashCode

		// value type
		receiver.setItemHashCode(receiver.MapType.Elem())
		break
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
			receiver.setItemHashCode(itemType)
			receiver.SliceType = reflect.SliceOf(receiver.GetItemMeta().ReflectType)
			break
		}

		// Dictionary类型
		if isTrue := types.IsDictionaryByType(receiver.ReflectType); isTrue {
			receiver.IsMap = true
			receiver.Type = Dic
			// 得到底层的map类型
			receiver.MapType = types.GetDictionaryMapType(receiver.ReflectType)
			// key type
			keyType := receiver.MapType.Key()
			keyVal := reflect.New(keyType).Elem()
			receiver.keyHashCode = PointerOfValue(keyVal).HashCode

			// value type
			receiver.setItemHashCode(receiver.MapType.Elem())
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
			receiver.setItemHashCode(itemType)
			receiver.SliceType = reflect.SliceOf(receiver.GetItemMeta().ReflectType)
			break
		}

		// 结构体
		if types.IsStruct(receiver.ReflectType) {
			receiver.Type = Struct
			receiver.IsStruct = true
			// 遍历结构体的字段
			for i := 0; i < numField; i++ {
				// 只加载允许导出的类型
				curStructField := receiver.ReflectType.Field(i)
				receiver.StructField = append(receiver.StructField, curStructField)
				// 加载允许导出的索引
				if curStructField.IsExported() {
					receiver.ExportedField = append(receiver.ExportedField, i)
				}
			}
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

func (receiver *TypeMeta) setItemHashCode(itemType reflect.Type) {
	if itemType.Kind() != reflect.Interface {
		itemVal := reflect.New(itemType).Elem()
		receiver.itemHashCode = PointerOfValue(itemVal).HashCode
	} else {
		receiver.itemHashCode = anyNil.HashCode
	}
}

func (receiver *TypeMeta) GetItemMeta() *TypeMeta {
	value, ok := cacheTyp.Load(receiver.itemHashCode)
	if ok {
		return value.(*TypeMeta)
	}
	return nil
}
