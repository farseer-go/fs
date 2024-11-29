package modules

import (
	"time"

	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/timingWheel"
	"github.com/farseer-go/fs/trace"
)

type FarseerKernelModule struct {
}

func (module FarseerKernelModule) DependsModule() []FarseerModule {
	return nil
}

func (module FarseerKernelModule) PreInitialize() {
	// 1、初始化配置
	configure.InitConfig()

	// 2、初始化日志
	log := flog.InitLog()

	// 4、配置日志
	if !container.IsRegister[core.ILog]() {
		container.RegisterInstance(log)
	}

	// 清空日志缓冲区
	flog.ClearLogBuffer(log)
	go flog.LoadLogBuffer(log)

	// 初始化时间轮
	timingWheel.NewDefault(100*time.Millisecond, 60)

	// 注册空的链路实现
	if !container.IsRegister[trace.IManager]() {
		container.Register(func() trace.IManager { return &trace.EmptyManager{} })
	}
}
