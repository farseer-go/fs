package modules

import (
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
	"reflect"
)

// GetDependModule 查找模块的依赖
func GetDependModule(module ...FarseerModule) []FarseerModule {
	var modules []FarseerModule
	for _, farseerModule := range module {
		dependsModules := farseerModule.DependsModule()
		if dependsModules != nil {
			modules = append(modules, GetDependModule(dependsModules...)...)
		}

		moduleName := reflect.TypeOf(farseerModule).String()
		flog.Println("Loading Module：" + flog.Colors[5](moduleName) + "")
		modules = append(modules, farseerModule)
		moduleMap[moduleName] = 0
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
	return append([]FarseerModule{FarseerKernelModule{}}, lst...)
}

// 判断模块是否存在于数组中
func exists(lst []FarseerModule, module FarseerModule) bool {
	for _, farseerModule := range lst {
		if reflect.ValueOf(farseerModule).String() == reflect.ValueOf(module).String() {
			return true
		}
	}
	return false
}

// StartModules 启动模块
func StartModules(farseerModules []FarseerModule) {
	//flog.Println("module initialization...")
	for _, farseerModule := range farseerModules {
		if module, ok := farseerModule.(FarseerPreInitializeModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(farseerModule).String()
			module.PreInitialize()
			flog.Println("Elapsed time：" + sw.GetMillisecondsText() + " " + moduleName + flog.Yellow(".PreInitialize()"))
			moduleMap[moduleName] += sw.ElapsedMilliseconds()
		}
	}

	for _, farseerModule := range farseerModules {
		if module, ok := farseerModule.(FarseerInitializeModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(farseerModule).String()
			module.Initialize()
			flog.Println("Elapsed time：" + sw.GetMillisecondsText() + " " + moduleName + flog.Blue(".Initialize()"))
			moduleMap[moduleName] += sw.ElapsedMilliseconds()
		}
	}

	for _, farseerModule := range farseerModules {
		if module, ok := farseerModule.(FarseerPostInitializeModule); ok {
			sw := stopwatch.StartNew()
			moduleName := reflect.TypeOf(farseerModule).String()
			module.PostInitialize()
			flog.Println("Elapsed time：" + sw.GetMillisecondsText() + " " + moduleName + flog.Green(".PostInitialize()"))
			moduleMap[moduleName] += sw.ElapsedMilliseconds()
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
