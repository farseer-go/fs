package data

import "github.com/farseernet/farseer.go/modules"

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return nil
}

func (module Module) PreInitialize() {
}

func (module Module) Initialize() {
}

func (module Module) PostInitialize() {
}

func (module Module) Shutdown() {
}
