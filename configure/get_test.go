package configure

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetArray(t *testing.T) {
	_ = ReadInConfig()
	array := viper.GetStringSlice("Log.Component")
	exists1 := viper.InConfig("Log.Component")
	str := viper.Get("Log.Component")
	exists2 := viper.InConfig("Log.Component.httpInvoke")
	mapstr := viper.GetStringMapString("Log.Component")
	fmt.Print(exists1)
	fmt.Print(exists2)
	fmt.Print(str)
	fmt.Print(mapstr)
	fmt.Println(len(array))
}

// 测试环境变量
func TestEnv(t *testing.T) {
	_ = ReadInConfig()
	fmt.Println(GetString("Database.default"))
	os.Setenv("DATABASE_DEFAULT", "aaa")
	assert.Equal(t, "aaa", GetString("Database.default"))
}
