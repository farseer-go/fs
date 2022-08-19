package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ITest interface {
	TestCall() string
}

type testStruct struct {
}

func (r testStruct) TestCall() string {
	return "hello world"
}

func Test_ioc_UseInstance(t *testing.T) {
	InitContainer()
	test := testStruct{}
	Use[ITest](test).Name("test1").Transient().Register()

	iocInstance := ResolveName[ITest]("test")
	assert.Nil(t, iocInstance)
	iocInstance = ResolveName[ITest]("test1")
	assert.NotNil(t, iocInstance)
	assert.Equal(t, iocInstance.TestCall(), "hello world")
}

func Test_ioc_UseFunc(t *testing.T) {
	InitContainer()
	Use[ITest](func() ITest { return testStruct{} }).Name("test1").Register()

	iocInstance := ResolveName[ITest]("test")
	assert.Nil(t, iocInstance)
	iocInstance = ResolveName[ITest]("test1")
	assert.NotNil(t, iocInstance)
	assert.Equal(t, iocInstance.TestCall(), "hello world")
}

func TestRegister(t *testing.T) {
	InitContainer()
	Register(func() ITest { return testStruct{} })

	iocInstance := ResolveName[ITest]("test")
	assert.Nil(t, iocInstance)
	iocInstance = Resolve[ITest]()
	assert.NotNil(t, iocInstance)
	assert.Equal(t, iocInstance.TestCall(), "hello world")
}
