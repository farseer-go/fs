package modules

import (
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
	"reflect"
	"strconv"
)

// GetDependModule 查找模块的依赖
func GetDependModule(module ...FarseerModule) []FarseerModule {
	var modules []FarseerModule
	for _, farseerModule := range module {
		dependsModules := farseerModule.DependsModule()
		if dependsModules != nil {
			modules = append(modules, GetDependModule(dependsModules...)...)
		}

		flog.Println("加载模块:" + reflect.TypeOf(farseerModule).String() + "")
		modules = append(modules, farseerModule)
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
		if farseerModule == module {
			return true
		}
	}
	return false
}

// StartModules 启动模块
func StartModules(farseerModules []FarseerModule) {
	flog.Println("Modules模块初始化...")
	sw := stopwatch.StartNew()
	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.PreInitialize()
		flog.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".PreInitialize()")
	}
	flog.Println("---------------------------------------")

	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.Initialize()
		flog.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".Initialize()")
	}
	flog.Println("---------------------------------------")

	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.PostInitialize()
		flog.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".PostInitialize()")
	}
	flog.Println("基础组件初始化完成")
}

// ShutdownModules 关闭模块
func ShutdownModules(farseerModules []FarseerModule) {
	flog.Println("Modules模块关闭...")
	sw := stopwatch.StartNew()
	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.Shutdown()
		flog.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".Shutdown()")
	}
	flog.Println("---------------------------------------")
}
