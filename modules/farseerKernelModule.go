package modules

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"math/rand"
	"time"
)

type FarseerKernelModule struct {
}

func (module FarseerKernelModule) DependsModule() []FarseerModule {
	return nil
}

func (module FarseerKernelModule) PreInitialize() {
	rand.Seed(int64(time.Now().Nanosecond()))
	container.InitContainer()
	configure.InitConfigure()
}

func (module FarseerKernelModule) Initialize() {
}

func (module FarseerKernelModule) PostInitialize() {
}

func (module FarseerKernelModule) Shutdown() {
}
