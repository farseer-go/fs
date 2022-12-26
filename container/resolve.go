package container

import (
	"reflect"
)

// Resolve 从容器中获取实例
// iocName = 别名
func Resolve[TInterface any](iocName ...string) TInterface {
	name := getIocName(iocName...)
	//var t TInterface
	interfaceType := reflect.TypeOf((*TInterface)(nil)).Elem()
	ins := defContainer.resolve(interfaceType, name)
	if ins == nil {
		var nilResult TInterface
		return nilResult
	}
	return ins.(TInterface)
}

// ResolveType 从容器中获取实例
// interfaceType = interface type
// iocName = 别名
func ResolveType(interfaceType reflect.Type, iocName ...string) any {
	name := getIocName(iocName...)
	return defContainer.resolve(interfaceType, name)
}
