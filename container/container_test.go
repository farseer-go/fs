package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type ITest interface {
	GetCount() int
	AddCount()
	ShowTime() string
}

type testStruct struct {
	count    int
	createAt time.Time
}

func (r *testStruct) GetCount() int {
	return r.count
}

func (r *testStruct) AddCount() {
	r.count++
}

func (r *testStruct) ShowTime() string {
	return r.createAt.String()
}

func TestRegister(t *testing.T) {
	InitContainer()
	// 注册单例
	Register(func() ITest { return &testStruct{createAt: time.Now()} })

	// 取一个不存在的别名的实例
	assert.Nil(t, Resolve[ITest]("test"))

	// 正常取出
	iocInstance := Resolve[ITest]()
	assert.NotNil(t, iocInstance)
	assert.Equal(t, iocInstance.GetCount(), 0)

	// 测试单例
	iocInstance.AddCount()
	iocInstance2 := Resolve[ITest]()
	assert.Equal(t, iocInstance2.GetCount(), 1)
	assert.Equal(t, iocInstance.ShowTime(), iocInstance2.ShowTime())

	// 注册临时对象
	RegisterTransient(func() ITest { return &testStruct{createAt: time.Now()} }, "test2")
	iocInstance = Resolve[ITest]("test2")
	iocInstance.AddCount()

	iocInstance2 = Resolve[ITest]("test2")
	assert.Equal(t, iocInstance2.GetCount(), 0)
	assert.NotEqual(t, iocInstance.ShowTime(), iocInstance2.ShowTime())
}
