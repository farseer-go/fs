package container

import (
	"reflect"

	"github.com/farseer-go/fs/container/eumLifecycle"
	"github.com/farseer-go/fs/exception"
)

// Container 容器操作
var defContainer = NewContainer()

func InitContainer() {
	defContainer = NewContainer()
}

// Register 注册实例，默认使用单例
func Register(constructor any, iocName ...string) {
	if defContainer == nil {
		exception.ThrowException("Please call fs.Initialize[Module]() to initialize the module first")
	}
	name := getIocName(iocName...)
	defContainer.registerConstructor(constructor, name, eumLifecycle.Single)
}

// RegisterTransient 注册实例，临时生命周期
func RegisterTransient(constructor any, iocName ...string) {
	if defContainer == nil {
		exception.ThrowException("Please call fs.Initialize[Module]() to initialize the module first")
	}
	name := getIocName(iocName...)
	defContainer.registerConstructor(constructor, name, eumLifecycle.Transient)
}

// RegisterInstance 注册实例，单例
func RegisterInstance[TInterface any](ins TInterface, iocName ...string) {
	if defContainer == nil {
		exception.ThrowException("Please call fs.Initialize[Module]() to initialize the module first")
	}
	name := getIocName(iocName...)
	defContainer.registerInstance((*TInterface)(nil), ins, name, eumLifecycle.Single)
}

// IsRegister 是否注册过
func IsRegister[TInterface any](iocName ...string) bool {
	name := getIocName(iocName...)
	interfaceType := reflect.TypeOf((*TInterface)(nil)).Elem()
	return defContainer.isRegister(interfaceType, name)
}

// IsRegisterType 判断类型是否注册过
func IsRegisterType(interfaceType reflect.Type, iocName ...string) bool {
	name := getIocName(iocName...)
	return defContainer.isRegister(interfaceType, name)
}

func getIocName(iocName ...string) string {
	if len(iocName) > 0 {
		return iocName[0]
	}
	return ""
}
