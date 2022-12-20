package test

import (
	"github.com/farseer-go/fs/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransient(t *testing.T) {
	container.InitContainer()

	// 注册临时对象
	container.RegisterTransient(func() IDatabase { return &mysql{} })

	mysql := container.Resolve[IDatabase]()
	mysql.SetTableName("user3")

	assert.Equal(t, container.Resolve[IDatabase]().GetTableName(), "")
}
