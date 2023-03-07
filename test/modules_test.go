package test

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"testing"
)

type StartupModule struct {
}

func (module StartupModule) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{modules.FarseerKernelModule{}, modules.FarseerKernelModule{}}
}

func (module StartupModule) Shutdown() {
}

func TestModules(t *testing.T) {
	assert.New(t).Panics(func() {
		modules.ThrowIfNotLoad(StartupModule{})
	})

	fs.Initialize[StartupModule]("test module")

	assert.True(t, modules.IsLoad(StartupModule{}))
}
