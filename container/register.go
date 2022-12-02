package container

import (
	"github.com/farseer-go/fs/exception"
	"github.com/studyzy/iocgo"
	"reflect"
)

// Container 容器操作
var container *iocgo.Container

func InitContainer() {
	container = iocgo.NewContainer()
}

type ioc[TInterface any] struct {
	isTransient bool   // 是否注册临时实例（默认为单例）
	name        string // 别名
	inter       *TInterface
	constructor any
}

// Register 注册实例，默认使用单例
func Register(constructor any) {
	if container == nil {
		exception.ThrowRefuseException("请先调用fs.Initialize[Module]()初始化模块")
	}
	_ = container.Register(constructor)
}

// Use 使用已存在的实例或函数注册
func Use[TInterface any](constructorOrInstance any) *ioc[TInterface] {
	if container == nil {
		exception.ThrowRefuseException("请先调用fs.Initialize[Module]()初始化模块")
	}
	return &ioc[TInterface]{
		constructor: constructorOrInstance,
	}
}

// Transient 临时模式（默认为单例模式）
func (receiver *ioc[TInterface]) Transient() *ioc[TInterface] {
	receiver.isTransient = true
	return receiver
}

// Name Ioc别名
func (receiver *ioc[TInterface]) Name(iocName string) *ioc[TInterface] {
	receiver.name = iocName
	return receiver
}

func (receiver *ioc[TInterface]) Register() {
	options := []iocgo.Option{iocgo.Lifestyle(receiver.isTransient)}
	if receiver.name != "" {
		options = append(options, iocgo.Name(receiver.name))
	}

	switch reflect.TypeOf(receiver.constructor).Kind() {
	case reflect.Struct:
		_ = container.RegisterInstance(receiver.inter, receiver.constructor, options...)
	case reflect.Func:
		_ = container.Register(receiver.constructor, options...)
	}
}
