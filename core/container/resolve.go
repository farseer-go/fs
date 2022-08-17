package container

import "github.com/farseer-go/fs/exception"

// Resolve 从容器中获取实例
func Resolve[T any]() (t T) {
	if container == nil {
		exception.ThrowRefuseException("请先调用fs.Initialize[Module]()初始化模块")
	}
	_ = container.Resolve(&t)
	return
}
