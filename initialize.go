package fs

import (
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
var StartupAt time.Time

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
	StartupAt = time.Now()
	AppId = snowflake.GenerateId()
	AppIp = net.GetIp()

	flog.Println("应用名称：", AppName)
	flog.Println("主机名称：", HostName)
	flog.Println("系统时间：", StartupAt)
	flog.Println("进程ID：", ProcessId)
	flog.Println("应用ID：", AppId)
	flog.Println("应用IP：", AppIp)
	flog.Println("---------------------------------------")

	var startupModule TModule
	flog.Println("加载模块...")
	dependModules = modules.Distinct(modules.GetDependModule(startupModule))
	flog.Println("加载完毕，共加载 " + strconv.Itoa(len(dependModules)) + " 个模块")
	flog.Println("---------------------------------------")

	modules.StartModules(dependModules)
	flog.Println("初始化完毕，共耗时" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms")
	flog.Println("---------------------------------------")

	if len(callbackFnList) > 0 {
		for index, fn := range callbackFnList {
			sw.Restart()
			fn()
			flog.Println("运行" + strconv.Itoa(index+1) + "：" + reflect.TypeOf(fn).String() + "，共耗时" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms")
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
