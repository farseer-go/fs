package container

import (
	"reflect"
)

// Resolve 从容器中获取实例
// iocName = 别名
func Resolve[TInterface any](iocName ...string) TInterface {
	name := ""
	if len(iocName) > 0 {
		name = iocName[0]
	}
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
	name := ""
	if len(iocName) > 0 {
		name = iocName[0]
	}

	return defContainer.resolve(interfaceType, name)
}
