package modules

import (
	"fmt"
	"sync"
	"time"

	"github.com/farseer-go/fs/color"
)

var moduleMap = make(map[string]time.Duration)
var moduleMapLocker sync.RWMutex

// IsLoad 模块是否加载
func IsLoad(module FarseerModule) bool {
	moduleMapLocker.RLock()
	defer moduleMapLocker.RUnlock()

	fullModuleName := getFullModuleName(module)

	_, isExists := moduleMap[fullModuleName]
	return isExists
}

// ThrowIfNotLoad 如果没加载模块时，退出应用
func ThrowIfNotLoad(module FarseerModule) {
	load := IsLoad(module)
	if !load {
		fullModuleName := getFullModuleName(module)
		panic(fmt.Sprintf("When using the %s module, you need to depend on the %s module in the startup module", color.Colors[4](fullModuleName), color.Colors[4](fullModuleName)))
	}
}
