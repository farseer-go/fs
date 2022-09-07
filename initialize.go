package fs

import (
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/net"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/snowflake"
	"github.com/farseer-go/fs/stopwatch"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

// StartupAt 应用启动时间
var StartupAt dateTime.DateTime

// AppName 应用名称
var AppName string

// HostName 主机名称
var HostName string

// AppId 应用ID
var AppId int64

// AppIp 应用IP
var AppIp string

// ProcessId 进程Id
var ProcessId int

// 依赖的模块
var dependModules []modules.FarseerModule

var callbackFnList []func()

// Initialize 初始化框架
func Initialize[TModule modules.FarseerModule](appName string) {
	sw := stopwatch.StartNew()

	AppName = appName
	ProcessId = os.Getppid()
	HostName, _ = os.Hostname()
	rand.Seed(time.Now().UnixNano())
	snowflake.Init(parse.HashCode64(HostName), rand.Int63n(32))
	StartupAt = dateTime.Now()
	AppId = snowflake.GenerateId()
	AppIp = net.GetIp()

	flog.Println("应用名称：", flog.Colors[2](AppName))
	flog.Println("主机名称：", flog.Colors[2](HostName))
	flog.Println("系统时间：", flog.Colors[2](StartupAt.ToString("yyyy-MM-dd hh:mm:ss")))
	flog.Println("进程ID：", flog.Colors[2](ProcessId))
	flog.Println("应用ID：", flog.Colors[2](AppId))
	flog.Println("应用IP：", flog.Colors[2](AppIp))
	flog.Println("---------------------------------------")

	var startupModule TModule
	flog.Println("加载模块...")
	dependModules = modules.Distinct(modules.GetDependModule(startupModule))
	flog.Println("加载完毕，共加载 " + strconv.Itoa(len(dependModules)) + " 个模块")
	flog.Println("---------------------------------------")

	modules.StartModules(dependModules)
	flog.Println("初始化完毕，共耗时：" + sw.GetMillisecondsText())
	flog.Println("---------------------------------------")

	if len(callbackFnList) > 0 {
		for index, fn := range callbackFnList {
			sw.Restart()
			fn()
			flog.Println("运行" + strconv.Itoa(index+1) + "：" + reflect.TypeOf(fn).String() + "，共耗时：" + sw.GetMillisecondsText())
			flog.Println("---------------------------------------")
		}
	}
}

// Exit 应用退出
func Exit() {
	modules.ShutdownModules(dependModules)
}

// AddInitCallback 添加框架启动完后执行的函数
func AddInitCallback(fn func()) {
	callbackFnList = append(callbackFnList, fn)
}
