package test

import (
	"fmt"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"testing"
)

type healthCheck1 struct {
	name string
}

func (c *healthCheck1) Check() (string, error) {
	return "healthCheck1", nil
}

type healthCheck2 struct {
	name string
}

func (c *healthCheck2) Check() (string, error) {
	return "healthCheck2", fmt.Errorf("test error")
}

type healthCheck1Module struct{}

func (t healthCheck1Module) DependsModule() []modules.FarseerModule {
	return nil
}
func (t healthCheck1Module) PreInitialize() {}
func (t healthCheck1Module) Initialize()    {}
func (t healthCheck1Module) PostInitialize() {
	container.Register(func() core.IHealthCheck {
		return &healthCheck1{}
	}, "healthCheck1")
	container.Register(func() core.IHealthCheck {
		return &healthCheck2{}
	}, "healthCheck2")
}
func (t healthCheck1Module) Shutdown() {}

func TestHealthCheck(t *testing.T) {
	assert.Panics(t, func() {
		fs.Initialize[healthCheck1Module]("unit test")
	})
}
