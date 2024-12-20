package modules

import (
	"fmt"
	"reflect"

	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
)

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
	for _, farseerModule := range module {
		dependsModules := farseerModule.DependsModule()
		if dependsModules != nil {
			modules = append(modules, GetDependModule(dependsModules...)...)
		}

		moduleName := reflect.TypeOf(farseerModule).String()
		flog.LogBuffer <- fmt.Sprint("Loading Module：" + flog.Colors[5](moduleName) + "")
		modules = append(modules, farseerModule)
		moduleMapLocker.Lock()
		moduleMap[moduleName] = 0
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
	for i := 0; i < len(lst); i++ {
		if reflect.ValueOf(lst[i]).String() == reflect.ValueOf(module).String() {
			return true
		}
	}
	return false
}

// StartModules 启动模块
func StartModules(farseerModules []FarseerModule) {
	for _, farseerModule := range farseerModules {
		if module, ok := farseerModule.(FarseerPreInitializeModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(farseerModule).String()
			module.PreInitialize()
			flog.LogBuffer <- fmt.Sprint("Elapsed time：" + sw.GetText() + " " + moduleName + flog.Yellow(".PreInitialize()"))
			moduleMap[moduleName] += sw.ElapsedDuration()
		}
	}

	for _, farseerModule := range farseerModules {
		if module, ok := farseerModule.(FarseerInitializeModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(farseerModule).String()
			module.Initialize()
			flog.LogBuffer <- fmt.Sprint("Elapsed time：" + sw.GetText() + " " + moduleName + flog.Blue(".Initialize()"))
			moduleMap[moduleName] += sw.ElapsedDuration()
		}
	}

	for _, farseerModule := range farseerModules {
		if module, ok := farseerModule.(FarseerPostInitializeModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(farseerModule).String()
			module.PostInitialize()
			flog.LogBuffer <- fmt.Sprint("Elapsed time：" + sw.GetText() + " " + moduleName + flog.Green(".PostInitialize()"))
			moduleMap[moduleName] += sw.ElapsedDuration()
		}
	}
}

// ShutdownModules 关闭模块
func ShutdownModules(farseerModules []FarseerModule) {
	flog.Println("Modules close...")
	for _, farseerModule := range farseerModules {
		if module, ok := farseerModule.(FarseerShutdownModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(farseerModule).String()
			module.Shutdown()
			flog.Println("Elapsed time：" + sw.GetMillisecondsText() + " " + moduleName + flog.Red(".Shutdown()"))
		}
	}
}
