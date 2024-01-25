package fastReflect

import (
	"reflect"
	"unsafe"
)

type FieldType int

const (
	List        FieldType = iota // 集合类型
	PageList    FieldType = iota // 集合类型
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

// ValueMeta 元数据
type ValueMeta struct {
	*TypeMeta
	//ReflectValue reflect.PointerValue  // 值
	PointerValue unsafe.Pointer // 值指向的内存地址
}

var cacheTyp = make(map[uint32]*TypeMeta)

// ValueOf 传入任意变量类型的值，得出该值对应的类型
func ValueOf(val any) ValueMeta {
	inf := (*emptyInterface)(unsafe.Pointer(&val))

	c, exists := cacheTyp[inf.typ.hash]
	if !exists {
		reflectType := reflect.TypeOf(val)
		c = typeOf(inf, reflectType)
		cacheTyp[inf.typ.hash] = c
	}

	return ValueMeta{
		TypeMeta:     c,
		PointerValue: inf.value,
	}
}

func Test(val any) ValueMeta {
	inf := (*emptyInterface)(unsafe.Pointer(&val))
	c, exists := cacheTyp[inf.typ.hash]
	if !exists {
		reflectType := reflect.TypeOf(val)
		c = typeOf(inf, reflectType)
		cacheTyp[inf.typ.hash] = c
	}
	return ValueMeta{
		TypeMeta:     c,
		PointerValue: inf.value,
	}
}
