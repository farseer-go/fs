package test

import (
	"github.com/farseer-go/fs/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

type IDatabaseFactory interface {
	CreateDatabase() IDatabase
}

type databaseFactory struct {
}

func (d *databaseFactory) CreateDatabase() IDatabase {
	return &sqlserver{}
}

func TestContainerMethod(t *testing.T) {
	//fs.Initialize[modules.FarseerKernelModule]("unit test")

	container.Remove[IDatabase]()
	// 注册获取IDatabase接口的方法
	container.Register(func(factory IDatabaseFactory) IDatabase {
		return factory.CreateDatabase()
	})

	assert.Panics(t, func() { assert.Equal(t, container.Resolve[IDatabase]().GetDbType(), "sqlserver") })

	// 注册IDatabaseFactory接口实例
	container.Register(func() IDatabaseFactory { return &databaseFactory{} })
	assert.Equal(t, container.Resolve[IDatabase]().GetDbType(), "sqlserver")
}
