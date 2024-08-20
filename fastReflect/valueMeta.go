package fastReflect

import (
	"reflect"
	"sync"
	"unsafe"
)

type FieldType int

var lock sync.RWMutex

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
var anyValueMap []any
var anyNil = TypeMeta{
	ReflectType:       reflect.TypeOf(anyValueMap).Elem(),
	ReflectTypeString: "interface {}",
	ReflectTypeBytes:  []byte("interface {}"),
	Type:              Interface,
	Kind:              reflect.Interface,
	HashCode:          252279353,
	Size:              16,
}

var cacheTyp sync.Map

func init() {
	cacheTyp.Store(uint32(252279353), &anyNil)
}

// PointerOfValue 传入任意变量类型的值，得出该值对应的类型
// //go:linkname nanotime1 reflect.nanotime1
func PointerOfValue(val reflect.Value) PointerMeta {
	inf := (*EmptyInterface)(unsafe.Pointer(&val))
	valueMeta := PointerMeta{PointerValue: inf.Value, HashCode: inf.Typ.hash}

	if typeMeta, exists := cacheTyp.Load(valueMeta.HashCode); exists {
		valueMeta.TypeMeta = typeMeta.(*TypeMeta)
		return valueMeta
	}
	valueMeta.TypeMeta = typeOf(val.Type(), inf)
	cacheTyp.Store(valueMeta.HashCode, valueMeta.TypeMeta)
	return valueMeta
}

// PointerOf 传入任意变量类型的值，得出该值对应的类型
func PointerOf(val any) PointerMeta {
	inf := (*EmptyInterface)(unsafe.Pointer(&val))
	valueMeta := PointerMeta{PointerValue: inf.Value, HashCode: inf.Typ.hash}

	if typeMeta, exists := cacheTyp.Load(valueMeta.HashCode); exists {
		valueMeta.TypeMeta = typeMeta.(*TypeMeta)
		return valueMeta
	}

	valueMeta.TypeMeta = typeOf(reflect.TypeOf(val), inf)
	cacheTyp.Store(valueMeta.HashCode, valueMeta.TypeMeta)
	return valueMeta
}

func Test(val any) PointerMeta {
	inf := (*EmptyInterface)(unsafe.Pointer(&val))
	if typeMeta, exists := cacheTyp.Load(inf.Typ.hash); exists {
		return PointerMeta{
			TypeMeta:     typeMeta.(*TypeMeta),
			PointerValue: inf.Value,
		}
	}

	reflectType := reflect.TypeOf(val)
	c := typeOf(reflectType, inf)
	cacheTyp.Store(inf.Typ.hash, c)
	return PointerMeta{
		TypeMeta:     c,
		PointerValue: inf.Value,
	}
}
