package test

import (
	"github.com/farseer-go/fs/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

type IMyDb interface {
	GetAllDbType() string
}
type myDb struct {
	Db1 IDatabase
	Db2 IDatabase `inject:"sqlserver"`
}

func (m myDb) GetAllDbType() string {
	return m.Db1.GetDbType() + "_" + m.Db2.GetDbType()
}

// 测试注入
func TestInject(t *testing.T) {
	container.InitContainer()
	// 注册单例
	container.RegisterTransient(func() IDatabase { return &mysql{} })
	// 注册单例
	container.Register(func() IDatabase { return &sqlserver{} }, "sqlserver")

	// 测试直接取对象的方式
	db := container.Resolve[myDb]()

	assert.Equal(t, db.Db1.GetDbType(), "mysql")
	assert.Equal(t, db.Db2.GetDbType(), "sqlserver")

	// 测试通过接口方式
	container.Register(func() IMyDb { return &myDb{} })
	myDb := container.Resolve[IMyDb]()
	assert.Equal(t, myDb.GetAllDbType(), "mysql_sqlserver")
}
