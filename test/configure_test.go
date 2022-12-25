package test

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetArray(t *testing.T) {
	fs.Initialize[modules.FarseerKernelModule]("unit test")
	assert.Equal(t, "", configure.GetString("a.b.c"))
	configure.SetDefault("a.b.c", "123")
	assert.Equal(t, "123", configure.GetString("a.b.c"))

	assert.Equal(t, "DataType=mysql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:steden@123@tcp(192.168.1.8:3306)/fss_demo?charset=utf8&parseTime=True&loc=Local", configure.GetString("Database.default"))

	arr := configure.GetSlice("A")
	assert.Len(t, arr, 3)
	assert.Equal(t, "a1", arr[0])
}

func TestConfigureGet(t *testing.T) {
	fs.Initialize[modules.FarseerKernelModule]("unit test")
	arr := configure.GetStrings("Database.default")
	assert.Equal(t, "DataType=mysql", arr[0])
	assert.Equal(t, "PoolMaxSize=50", arr[1])
	assert.Equal(t, "PoolMinSize=1", arr[2])

	assert.Equal(t, 20, configure.GetInt("FSS.ReservedTaskCount"))
	assert.Equal(t, int64(20), configure.GetInt64("FSS.ReservedTaskCount"))
	assert.Equal(t, true, configure.GetBool("Log.Component.httpInvoke"))

	assert.Len(t, configure.GetSubNodes("A.B"), 0)
	assert.Len(t, configure.GetSlice("A.B"), 0)

	assert.Equal(t, "DataType=mysql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:steden@123@tcp(192.168.1.8:3306)/fss_demo?charset=utf8&parseTime=True&loc=Local", configure.GetString("Database.default"))
	os.Setenv("Database_default", "aaa")
	assert.Equal(t, "aaa", configure.GetString("Database.default"))
}
