package modules

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/flog"
)

type FarseerKernelModule struct {
}

func (module FarseerKernelModule) DependsModule() []FarseerModule {
	return nil
}

func (module FarseerKernelModule) PreInitialize() {
	container.InitContainer()
	err := configure.ReadInConfig()
	if err != nil { // 捕获读取中遇到的error
		flog.Errorf("读取配置文件farseer.yaml时发生错误: %s \n", err)
	} else {
		flog.Println("farseer.yaml配置加载正常")
	}
}

func (module FarseerKernelModule) Initialize() {
}

func (module FarseerKernelModule) PostInitialize() {
}

func (module FarseerKernelModule) Shutdown() {
}
