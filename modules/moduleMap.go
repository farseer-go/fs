package modules

import (
	"fmt"
	"github.com/farseer-go/fs/flog"
	"reflect"
	"sync"
)

var moduleMap = make(map[string]int64)
var moduleMapLocker sync.RWMutex

// IsLoad 模块是否加载
func IsLoad(module FarseerModule) bool {
	moduleMapLocker.RLock()
	defer moduleMapLocker.RUnlock()

	moduleName := reflect.TypeOf(module).String()
	_, isExists := moduleMap[moduleName]
	return isExists
}

// ThrowIfNotLoad 如果没加载模块时，退出应用
func ThrowIfNotLoad(module FarseerModule) {
	load := IsLoad(module)
	if !load {
		moduleName := reflect.TypeOf(module).String()
		panic(fmt.Sprintf("When using the %s module, you need to depend on the %s module in the startup module", flog.Colors[4](moduleName), flog.Colors[4](moduleName)))
	}
}
