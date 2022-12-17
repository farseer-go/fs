package container

import (
	"reflect"
)

// Resolve 从容器中获取实例
// iocName = 别名
func Resolve[T any](iocName ...string) T {
	name := ""
	if len(iocName) > 0 {
		name = iocName[0]
	}
	var t *T
	interfaceType := reflect.TypeOf(t).Elem()
	ins := defContainer.resolve(interfaceType, name)
	if ins == nil {
		var nilResult T
		return nilResult
	}
	return ins.(T)
}
