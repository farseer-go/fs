package container

import (
	"github.com/farseer-go/fs/exception"
	"github.com/studyzy/iocgo"
)

// Container 容器操作
var container *iocgo.Container

func InitContainer() {
	container = iocgo.NewContainer()
}

// Register 注册接口
func Register(constructor any) error {
	if container == nil {
		exception.ThrowRefuseException("请先调用fs.Initialize[Module]()初始化模块")
	}
	return container.Register(constructor)
}
