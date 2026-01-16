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
	//fs.Initialize[modules.FarseerKernelModule]("unit test")

	// 注册单例
	container.RegisterTransient(func() IDatabase { return &mysql{} })
	// 注册单例
	container.Register(func() IDatabase { return &sqlserver{} }, "sqlserver")

	// 直接取对象的方式
	db := container.Resolve[myDb]()

	assert.Equal(t, db.Db1.GetDbType(), "mysql")
	assert.Equal(t, db.Db2.GetDbType(), "sqlserver")

	// 通过接口方式
	container.Register(func() IMyDb { return &myDb{} })
	iMyDb := container.Resolve[IMyDb]()
	assert.Equal(t, iMyDb.GetAllDbType(), "mysql_sqlserver")

	// 通过实例的方式
	dbIns := container.ResolveIns(&myDb{})
	assert.Equal(t, dbIns.GetAllDbType(), "mysql_sqlserver")
}
