package container

import "github.com/studyzy/iocgo"

//Container 容器操作
var container *iocgo.Container

func init() {
	container = iocgo.NewContainer()
}

// Register 注册接口
func Register(constructor any) error {
	return container.Register(constructor)
}
