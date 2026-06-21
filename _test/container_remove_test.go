package test

import (
	"testing"
	"time"

	"github.com/farseer-go/fs/container"
	"github.com/stretchr/testify/assert"
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

func TestContainerRemoveUnused_test(t *testing.T) {
	container.InitContainer()
	// 注册一个默认的（临时生命周期：RemoveUnused 的访问时间淘汰仅对临时实例有意义，
	// 单例创建后常驻、不再刷新访问时间）
	container.RegisterTransient(func() IDatabase { return &mysql{} })
	// 注册一个testName
	container.RegisterTransient(func() IDatabase { return &mysql{} }, "testName")
	assert.Equal(t, 2, len(container.ResolveAll[IDatabase]()))
	assert.Equal(t, true, container.IsRegister[IDatabase]())
	assert.Equal(t, true, container.IsRegister[IDatabase]("testName"))

	// 移除 testName
	container.RemoveUnused[IDatabase](time.Second)
	assert.Equal(t, 2, len(container.ResolveAll[IDatabase]()))
	assert.Equal(t, true, container.IsRegister[IDatabase]())
	assert.Equal(t, true, container.IsRegister[IDatabase]("testName"))

	time.Sleep(10 * time.Millisecond)
	container.ResolveAll[IDatabase]()
	container.RemoveUnused[IDatabase](5 * time.Millisecond)
	assert.Equal(t, 2, len(container.ResolveAll[IDatabase]()))
	assert.Equal(t, true, container.IsRegister[IDatabase]())
	assert.Equal(t, true, container.IsRegister[IDatabase]("testName"))

	time.Sleep(10 * time.Millisecond)
	container.Resolve[IDatabase]()
	container.RemoveUnused[IDatabase](5 * time.Millisecond)
	assert.Equal(t, 1, len(container.ResolveAll[IDatabase]()))
	assert.Equal(t, true, container.IsRegister[IDatabase]())
	assert.Equal(t, false, container.IsRegister[IDatabase]("testName"))

	time.Sleep(10 * time.Millisecond)
	container.RemoveUnused[IDatabase](5 * time.Millisecond)
	assert.Equal(t, 0, len(container.ResolveAll[IDatabase]()))
	assert.Equal(t, false, container.IsRegister[IDatabase]())
	assert.Equal(t, false, container.IsRegister[IDatabase]("testName"))
}
