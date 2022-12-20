package test

import (
	"github.com/farseer-go/fs/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingle(t *testing.T) {
	container.InitContainer()
	// 注册单例
	container.Register(func() IDatabase { return &mysql{} })
	// 别名重复了
	assert.Panics(t, func() {
		container.Register(func() IDatabase { return &sqlserver{} })
	})

	container.Register(func() IDatabase { return &mysql{} }, "mysql")
	container.Register(func() IDatabase { return &sqlserver{} }, "sqlserver")

	// 取一个不存在的别名的实例
	assert.Nil(t, container.Resolve[IDatabase]("test"))

	// 正常取出
	assert.Equal(t, container.Resolve[IDatabase]().GetDbType(), "mysql")
	assert.Equal(t, container.Resolve[IDatabase]("mysql").GetDbType(), "mysql")
	assert.Equal(t, container.Resolve[IDatabase]("sqlserver").GetDbType(), "sqlserver")

	// 测试单例
	container.Resolve[IDatabase]("mysql").SetTableName("user1")
	assert.Equal(t, container.Resolve[IDatabase]("mysql").GetTableName(), "user1")

	db := container.Resolve[IDatabase]("sqlserver")
	db.SetTableName("user2")
	assert.Equal(t, container.Resolve[IDatabase]("sqlserver").GetTableName(), "user2")
}
