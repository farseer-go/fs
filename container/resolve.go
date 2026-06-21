package container

import (
	"reflect"

	"github.com/farseer-go/fs/flog"
)

// Resolve 从容器中获取实例
// iocName = 别名
func Resolve[TInterface any](iocName ...string) TInterface {
	name := getIocName(iocName...)
	//var t TInterface
	interfaceType := reflect.TypeOf((*TInterface)(nil)).Elem()
	ins, err := defContainer.resolveField(interfaceType, name)
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
	return defContainer.resolveField(interfaceType, name)
}

// IsSingleType 判断指定接口类型+别名是否注册为单例。
// 用于调用方在初始化阶段判断某依赖能否被安全地预解析并长期缓存。
// 仅对 interface 类型有意义；struct 等动态创建类型返回 false。
func IsSingleType(interfaceType reflect.Type, iocName ...string) bool {
	name := getIocName(iocName...)
	return defContainer.isSingle(interfaceType, name)
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
