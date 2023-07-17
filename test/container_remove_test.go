package test

import (
	"github.com/farseer-go/fs/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainerRemove_test(t *testing.T) {
	container.InitContainer()
	// 注册一个默认的
	container.Register(func() IDatabase { return &mysql{} })
	// 注册一个testName
	container.Register(func() IDatabase { return &mysql{} }, "testName")
	assert.Equal(t, 2, len(container.ResolveAll[IDatabase]()))
	assert.Equal(t, true, container.IsRegister[IDatabase]())
	assert.Equal(t, true, container.IsRegister[IDatabase]("testName"))

	// 移除 testName
	container.Remove[IDatabase]("testName")
	assert.Equal(t, 1, len(container.ResolveAll[IDatabase]()))
	assert.Equal(t, true, container.IsRegister[IDatabase]())
	assert.Equal(t, false, container.IsRegister[IDatabase]("testName"))

	// 移除 默认的
	container.Remove[IDatabase]()
	assert.Equal(t, 0, len(container.ResolveAll[IDatabase]()))
	assert.Equal(t, false, container.IsRegister[IDatabase]())
	assert.Equal(t, false, container.IsRegister[IDatabase]("testName"))
}
