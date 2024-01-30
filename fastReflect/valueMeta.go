package fastReflect

import (
	"reflect"
	"unsafe"
)

type FieldType int

const (
	List        FieldType = iota // 集合类型
	PageList                     // 集合类型
	CustomList                   // 自定义List类型
	Slice                        // 结构体类型
	Array                        // 结构体类型
	Map                          // map类型
	Dic                          // 字典类型
	GoBasicType                  // 基础类型
	Struct                       // 结构体类型
	Chan                         // 结构体类型
	Interface                    // 结构体类型
	Func                         // 结构体类型
	Invalid                      // 结构体类型
	Unknown                      // 结构体类型
)

// PointerMeta 元数据
type PointerMeta struct {
	*TypeMeta
	PointerValue unsafe.Pointer // 值指向的内存地址
	HashCode     uint32
}

// 字段为any类型，当为nil时无法获取，需手动设置类型
var anyNil = &TypeMeta{
	Name:              "interface {}",
	ReflectType:       nil,
	ReflectTypeString: "interface {}",
	Type:              Interface,
	IsAddr:            false,
	NumField:          0,
	StructField:       nil,
	ItemMeta:          nil,
	SliceType:         nil,
	ZeroValue:         nil,
	Kind:              reflect.Interface,
	IsNumber:          false,
	IsEmum:            false,
	IsString:          false,
	IsBool:            false,
	IsTime:            false,
	IsDateTime:        false,
	IsSliceOrArray:    false,
	IsStruct:          false,
	HashCode:          252279353,
	Size:              16,
	TypeIdentity:      "",
}
var cacheTyp = map[uint32]*TypeMeta{252279353: anyNil}

// PointerOf 传入任意变量类型的值，得出该值对应的类型
func PointerOf(val any) PointerMeta {
	inf := (*EmptyInterface)(unsafe.Pointer(&val))
	valueMeta := PointerMeta{PointerValue: inf.Value}
	if inf.Typ != nil {
		valueMeta.HashCode = inf.Typ.hash
	} else {
		valueMeta.HashCode = 252279353
	}
	var exists bool
	if valueMeta.TypeMeta, exists = cacheTyp[valueMeta.HashCode]; !exists {
		valueMeta.TypeMeta = typeOf(reflect.TypeOf(val), inf)
		cacheTyp[valueMeta.HashCode] = valueMeta.TypeMeta
		return valueMeta
	}
	return valueMeta
}

func Test(val any) PointerMeta {
	inf := (*EmptyInterface)(unsafe.Pointer(&val))
	c, exists := cacheTyp[inf.Typ.hash]
	if !exists {
		reflectType := reflect.TypeOf(val)
		c = typeOf(reflectType, inf)
		cacheTyp[inf.Typ.hash] = c
	}
	return PointerMeta{
		TypeMeta:     c,
		PointerValue: inf.Value,
	}
}
