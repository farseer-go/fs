package modules

import (
	"github.com/farseernet/farseer.go/utils/stopwatch"
	"log"
	"reflect"
	"strconv"
)

// 查找模块的依赖
func getDependModule(module ...FarseerModule) []FarseerModule {
	var modules []FarseerModule
	for _, farseerModule := range module {
		dependsModules := farseerModule.DependsModule()
		if dependsModules != nil {
			modules = append(modules, getDependModule(dependsModules...)...)
		}

		log.Println("加载模块:" + reflect.TypeOf(farseerModule).String() + "")
		modules = append(modules, farseerModule)
	}
	return modules
}

// StartModules 启动模块
func StartModules(module ...FarseerModule) {
	log.Println("加载模块...")
	farseerModules := getDependModule(module...)
	log.Println("加载完毕，共加载 " + strconv.Itoa(len(farseerModules)) + " 个模块")
	log.Println("---------------------------------------")

	log.Println("Modules模块初始化...")
	sw := stopwatch.StartNew()
	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.PreInitialize()
		log.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".PreInitialize()")
	}

	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.Initialize()
		log.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".Initialize()")
	}

	for _, farseerModule := range farseerModules {
		sw.Restart()
		farseerModule.PostInitialize()
		log.Println("耗时：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms " + reflect.TypeOf(farseerModule).String() + ".PostInitialize()")
	}
	log.Println("基础组件初始化完成")
}
