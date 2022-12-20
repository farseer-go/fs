package test

import (
	"github.com/farseer-go/fs/configure"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetArray(t *testing.T) {
	_ = configure.ReadInConfig()
	assert.Equal(t, "", configure.GetString("a.b.c"))
	configure.SetDefault("a.b.c", "123")
	assert.Equal(t, "123", configure.GetString("a.b.c"))

	assert.Equal(t, "DataType=mysql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:steden@123@tcp(192.168.1.8:3306)/fss_demo?charset=utf8&parseTime=True&loc=Local", configure.GetString("Database.default"))

	arr := configure.GetSlice("A")
	assert.Len(t, arr, 3)
	assert.Equal(t, "a1", arr[0])
}

// 测试环境变量
func TestEnv(t *testing.T) {
	_ = configure.ReadInConfig()
	assert.Equal(t, "DataType=mysql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:steden@123@tcp(192.168.1.8:3306)/fss_demo?charset=utf8&parseTime=True&loc=Local", configure.GetString("Database.default"))
	os.Setenv("Database_default", "aaa")
	assert.Equal(t, "aaa", configure.GetString("Database.default"))
}
