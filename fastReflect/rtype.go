package fastReflect

import "unsafe"

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

// arrayType represents a fixed array type.
type arrayType struct {
	rtype
	elem  *rtype // array element type
	slice *rtype // slice type
	len   uintptr
}

// chanType represents a channel type.
type chanType struct {
	rtype
	elem *rtype  // channel element type
	dir  uintptr // channel direction (ChanDir)
}

type funcType struct {
	rtype
	inCount  uint16
	outCount uint16 // top bit is set if last input parameter is ...
}
type interfaceType struct {
	rtype
	pkgPath name      // import path
	methods []imethod // sorted by hash
}

type imethod struct {
	name nameOff // name of method
	typ  typeOff // .(*FuncType) underneath
}

// mapType represents a map type.
type mapType struct {
	rtype
	key    *rtype // map key type
	elem   *rtype // map element (value) type
	bucket *rtype // internal bucket structure
	// function for hashing keys (ptr to key, seed) -> hash
	hasher     func(unsafe.Pointer, uintptr) uintptr
	keysize    uint8  // size of key slot
	valuesize  uint8  // size of value slot
	bucketsize uint16 // size of bucket
	flags      uint32
}

// ptrType represents a pointer type.
type ptrType struct {
	rtype
	elem *rtype // pointer element (pointed at) type
}

// sliceType represents a slice type.
type sliceType struct {
	rtype
	elem *rtype // slice element type
}
type structField struct {
	name   name    // name is always non-empty
	typ    *rtype  // type of field
	offset uintptr // byte offset of field
}

type structType struct {
	rtype
	pkgPath name
	fields  []structField // sorted by offset
}

type name struct {
	bytes *byte
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
