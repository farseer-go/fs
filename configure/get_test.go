package configure

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestGetArray(t *testing.T) {
	ReadInConfig()
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
