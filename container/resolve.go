package container

import (
	"github.com/farseer-go/fs/flog"
	"reflect"
)

// Resolve 从容器中获取实例
// iocName = 别名
func Resolve[TInterface any](iocName ...string) TInterface {
	name := getIocName(iocName...)
	//var t TInterface
	interfaceType := reflect.TypeOf((*TInterface)(nil)).Elem()
	ins, err := defContainer.resolve(interfaceType, name)
	if err != nil {
		_ = flog.Error(err)
		var nilResult TInterface
		return nilResult
	}
	return ins.(TInterface)
}

// ResolveType 从容器中获取实例
// interfaceType = interface type
// iocName = 别名
func ResolveType(interfaceType reflect.Type, iocName ...string) (any, error) {
	name := getIocName(iocName...)
	return defContainer.resolve(interfaceType, name)
}

// ResolveIns 将现有中实例内的字段做注入操作
func ResolveIns[TIns any](ins TIns) TIns {
	return defContainer.inject(ins).(TIns)
}

// ResolveAll 从容器中获取所有实例
func ResolveAll[TInterface any]() []TInterface {
	//var t TInterface
	interfaceType := reflect.TypeOf((*TInterface)(nil)).Elem()
	arrAny := defContainer.resolveAll(interfaceType)
	var arrIns []TInterface
	for _, ins := range arrAny {
		arrIns = append(arrIns, ins.(TInterface))
	}
	return arrIns
}
