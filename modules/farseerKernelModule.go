package modules

import (
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
	container.InitContainer()
	flog.Init()
	timingWheel.NewDefault(100*time.Millisecond, 60)
}

func (module FarseerKernelModule) Initialize() {
}

func (module FarseerKernelModule) PostInitialize() {
}

func (module FarseerKernelModule) Shutdown() {
}
