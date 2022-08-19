package container

import (
	"github.com/farseer-go/fs/exception"
	"github.com/studyzy/iocgo"
)

// Resolve 从容器中获取实例
func Resolve[T any]() (t T) {
	if container == nil {
		exception.ThrowRefuseException("请先调用fs.Initialize[Module]()初始化模块")
	}
	_ = container.Resolve(&t)
	return
}

// ResolveName 指定ioc别名从容器中获取实例
func ResolveName[T any](iocName string) (t T) {
	if container == nil {
		exception.ThrowRefuseException("请先调用fs.Initialize[Module]()初始化模块")
	}
	_ = container.Resolve(&t, iocgo.ResolveName(iocName))
	return
}
