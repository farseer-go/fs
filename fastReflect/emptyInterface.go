package fastReflect

import "unsafe"

// SliceHeader 切片的底层结构
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// StringHeader is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
//
// In new code, use unsafe.String or unsafe.StringData instead.
type StringHeader struct {
	Data uintptr
	Len  int
}

// EmptyInterface any底层结构，参考reflect.Value的emptyInterface设计
type EmptyInterface struct {
	Typ   *rtype
	Value unsafe.Pointer
}
