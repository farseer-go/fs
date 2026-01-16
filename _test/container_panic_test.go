package test

import (
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainerPanic_test(t *testing.T) {
	container.InitContainer()
	assert.Panics(t, func() {
		container.Register(func(db IDatabase) IDatabase { return &mysql{} })
	})

	assert.Panics(t, func() {
		container.Register(func() (IDatabase, error) { return &mysql{}, nil })
	})

	assert.Panics(t, func() {
		container.Register(func() string { return "" })
	})

	assert.Panics(t, func() {
		container.Register(func(str string) IDatabase { return &mysql{} })
	})

	assert.Panics(t, func() {
		container.RegisterInstance[mysql](mysql{})
	})

	assert.Nil(t, container.Resolve[modules.FarseerModule]())
	assert.Equal(t, "", container.Resolve[string]())

	assert.Panics(t, func() {
		container.Register(&mysql{})
	})

	container.RegisterTransient(func() IDatabase { return nil }, "testNil")

	assert.Panics(t, func() {
		panic("")
	})

	assert.Panics(t, func() {
		panic("")
	})
}
