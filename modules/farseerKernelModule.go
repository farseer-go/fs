package modules

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/core/container"
	"github.com/farseer-go/fs/net"
)

type FarseerKernelModule struct {
}

func (module FarseerKernelModule) DependsModule() []FarseerModule {
	return nil
}

func (module FarseerKernelModule) PreInitialize() {
	container.InitContainer()
	configure.InitConfigure()
	net.InitNet()
}

func (module FarseerKernelModule) Initialize() {
}

func (module FarseerKernelModule) PostInitialize() {
}

func (module FarseerKernelModule) Shutdown() {
}
