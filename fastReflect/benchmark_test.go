package fastReflect

import (
	"unsafe"
)

func toT[T any](source any, defval T) T {
	s := ValueOf(source)
	if s.IsString {
		return *(*T)(s.PointerValue)
	}
	//if s.IsString {
	//	str := source.(string)
	//	return *(*T)(unsafe.Pointer(&str))
	//}
	return *(*T)(unsafe.Pointer(&source))
}
