package configure

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestGetArray(t *testing.T) {
	InitConfigure()
	array := viper.GetStringSlice("Log.CloseComponent")

	exists1 := viper.InConfig("Log.CloseComponent")
	str := viper.Get("Log.CloseComponent")
	exists2 := viper.InConfig("Log.CloseComponent.httpInvoke")
	fmt.Print(exists1)
	fmt.Print(exists2)
	fmt.Print(str)
	fmt.Println(len(array))
}
