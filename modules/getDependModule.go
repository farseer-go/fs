package modules

import (
	"github.com/farseer-go/fs/stopwatch"
	"log"
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

		log.Println("加载模块:" + reflect.TypeOf(farseerModule).String() + "")
		modules = append(modules, farseerModule)
	}
	return modules
}

// StartModules 启动模块
func StartModules(farseerModules []FarseerModule) {
	log.Println("Modules模块初始化...")
	sw := stopwatch.StartNew()
	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.PreInitialize()
		log.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".PreInitialize()")
	}
	log.Println("---------------------------------------")

	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.Initialize()
		log.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".Initialize()")
	}
	log.Println("---------------------------------------")

	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.PostInitialize()
		log.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".PostInitialize()")
	}
	log.Println("基础组件初始化完成")
}

// ShutdownModules 关闭模块
func ShutdownModules(farseerModules []FarseerModule) {
	log.Println("Modules模块关闭...")
	sw := stopwatch.StartNew()
	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.Shutdown()
		log.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".Shutdown()")
	}
	log.Println("---------------------------------------")
}
