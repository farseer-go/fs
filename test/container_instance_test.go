package test

import (
	"github.com/farseer-go/fs/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstance(t *testing.T) {
	test := &mysql{}

	container.InitContainer()
	// 根据实例注册
	container.RegisterInstance[IDatabase](test)
	test.SetTableName("user4")
	assert.Equal(t, test.GetTableName(), "user4")

	// 正常取出
	iocInstance := container.Resolve[IDatabase]()
	assert.Equal(t, iocInstance.GetTableName(), "user4")
	iocInstance.SetTableName("user5")

	assert.Equal(t, container.Resolve[IDatabase]().GetTableName(), "user5")
}
