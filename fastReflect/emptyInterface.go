package fastReflect

import "unsafe"

// SliceHeader 切片的底层结构
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// EmptyInterface any底层结构，参考reflect.Value的emptyInterface设计
type EmptyInterface struct {
	Typ   *rtype
	Value unsafe.Pointer
}

// tflag 用于表示 rtype 值后面的内存中有哪些额外的类型信息。
// rtype 值后面的内存中还有哪些额外的类型信息。
//
// tflag 值必须与 rtype 值的副本保持同步：
// cmd/compile/internal/gc/reflect.go
// cmd/link/internal/ld/decodesym.go
// runtime/type.go
type tflag uint8

type nameOff int32 // 偏移到名称
type typeOff int32 // 偏移到*rtype

// rtype 是大多数值的通用实现。它被嵌入到其他结构类型中。
// 必须与 ../runtime/type.go:/^type._type.同步。
type rtype struct {
	size       uintptr                                   // 变量占用内存字节大小
	ptrdata    uintptr                                   // 类型中可以包含指针的字节数
	hash       uint32                                    // 类型的哈希值
	tflag      tflag                                     // 额外类型信息标志
	align      uint8                                     // 该类型变量的对齐方式
	fieldAlign uint8                                     // 该类型 struct 字段的对齐方式
	kind       uint8                                     // 类型枚举
	equal      func(unsafe.Pointer, unsafe.Pointer) bool // 用于比较此类型对象的函数 (ptr to object A, ptr to object B) -> ==?
	gcdata     *byte                                     // 垃圾收集数据
	str        nameOff                                   // 字符串形式
	ptrToThis  typeOff                                   // 该类型指针的类型，可能为 0
}
