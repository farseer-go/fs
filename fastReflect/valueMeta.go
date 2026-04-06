package fastReflect

import (
	"reflect"
	"sync"
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

	// 先占位写入缓存，防止自引用结构体（如 *VO 含 Top *VO）导致 parseType 无限递归
	placeholder := &TypeMeta{ReflectType: val.Type(), HashCode: inf.Typ.hash}
	cacheTyp.Store(valueMeta.HashCode, placeholder)

	tm := typeOf(val.Type(), inf)
	// 原地更新 placeholder 的字段，而非替换指针。
	// 这样所有已通过缓存命中拿到 placeholder 指针的地方（如 FieldTypeMetas）
	// 也会自动看到完整的 TypeMeta，避免自引用类型的字段被识别为零值 Type（List）。
	*placeholder = *tm
	valueMeta.TypeMeta = placeholder
	return valueMeta
}

// PointerOfValueWithMeta 已知 TypeMeta 时的快速路径：跳过 sync.Map 查找，只取 PointerValue
// 适用于热路径中类型已经通过 FieldTypeMetas 缓存的场景
func PointerOfValueWithMeta(val reflect.Value, tm *TypeMeta) PointerMeta {
	inf := (*EmptyInterface)(unsafe.Pointer(&val))
	return PointerMeta{
		TypeMeta:     tm,
		PointerValue: inf.Value,
		HashCode:     tm.HashCode,
	}
}

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
