package modules

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/timingWheel"
	"time"
)

type FarseerKernelModule struct {
}

func (module FarseerKernelModule) DependsModule() []FarseerModule {
	return nil
}

func (module FarseerKernelModule) PreInitialize() {
	// 初始化配置
	configure.InitConfig()

	// 初始化日志
	log := flog.InitLog()

	// 初始化容器
	container.InitContainer(log)

	// 配置日志
	container.RegisterInstance(log)

	// 初始化时间轮
	timingWheel.NewDefault(100*time.Millisecond, 60)
}
