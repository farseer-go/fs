package modules

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"

	"github.com/farseer-go/fs/color"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
)

func getFullModuleName(m FarseerModule) string {
	t := reflect.TypeOf(m)
	// 如果未来可能传入指针，建议加上：if t.Kind() == reflect.Ptr { t = t.Elem() }
	return t.PkgPath() + "." + t.Name()
}

func GetDependModules(startupModule FarseerModule) []FarseerModule {
	dependModules := Distinct(GetDependModule(startupModule))

	// 加载核心模块
	if !exists(dependModules, FarseerKernelModule{}) {
		return append([]FarseerModule{FarseerKernelModule{}}, dependModules...)
	}

	return dependModules
}

// GetDependModule 查找模块的依赖
func GetDependModule(module ...FarseerModule) []FarseerModule {
	var modules []FarseerModule
	for _, FarseerModule := range module {
		dependsModules := FarseerModule.DependsModule()
		if dependsModules != nil {
			modules = append(modules, GetDependModule(dependsModules...)...)
		}

		fullModuleName := getFullModuleName(FarseerModule)
		modules = append(modules, FarseerModule)
		moduleMapLocker.Lock()
		moduleMap[fullModuleName] = 0
		moduleMapLocker.Unlock()
	}

	return modules
}

// Distinct 模块去重
func Distinct(modules []FarseerModule) []FarseerModule {
	var lst []FarseerModule
	for _, module := range modules {
		if !exists(lst, module) {
			lst = append(lst, module)
		}
	}
	return lst
}

// 判断模块是否存在于数组中
func exists(lst []FarseerModule, module FarseerModule) bool {
	curID := getFullModuleName(module)

	for _, item := range lst {
		if getFullModuleName(item) == curID {
			return true
		}
	}
	return false
}

// StartModules 启动模块
func StartModules(LibModules []FarseerModule) {
	for _, FarseerModule := range LibModules {
		if module, ok := FarseerModule.(FarseerPreInitializeModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(FarseerModule).String()
			module.PreInitialize()
			flog.LogBuffer <- fmt.Sprint("Elapsed time: " + sw.GetText() + " " + moduleName + color.Yellow(".PreInitialize()"))
			moduleMap[moduleName] += sw.ElapsedDuration()
		}
	}

	for _, FarseerModule := range LibModules {
		if module, ok := FarseerModule.(FarseerInitializeModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(FarseerModule).String()
			module.Initialize()
			flog.LogBuffer <- fmt.Sprint("Elapsed time: " + sw.GetText() + " " + moduleName + color.Blue(".Initialize()"))
			moduleMap[moduleName] += sw.ElapsedDuration()
		}
	}

	for _, FarseerModule := range LibModules {
		if module, ok := FarseerModule.(FarseerPostInitializeModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(FarseerModule).String()
			module.PostInitialize()
			flog.LogBuffer <- fmt.Sprint("Elapsed time: " + sw.GetText() + " " + moduleName + color.Green(".PostInitialize()"))
			moduleMap[moduleName] += sw.ElapsedDuration()
		}
	}
}

var onceExit sync.Once

// ShutdownModules 关闭模块
func ShutdownModules(LibModules []FarseerModule) {
	onceExit.Do(func() {
		flog.Println("Modules close...")
		for _, FarseerModule := range LibModules {
			if module, ok := FarseerModule.(FarseerShutdownModule); ok {
				sw := stopwatch.StartNew()
				moduleName := reflect.TypeOf(FarseerModule).String()
				module.Shutdown()
				flog.Println("Elapsed time: " + sw.GetMillisecondsText() + " " + moduleName + color.Red(".Shutdown()"))
			}
		}

		if len(core.CallbackExitList) > 0 {
			sw := stopwatch.StartNew()
			for index, fn := range core.CallbackExitList {
				sw.Restart()
				fn.F()
				flog.Println("Run " + strconv.Itoa(index+1) + ": " + fn.Name + ", Use: " + sw.GetText())
			}
			flog.Println("---------------------------------------")
		}
	})
}
