package fs

import (
	"github.com/farseer-go/fs/configure"
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
	"strings"
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
func Initialize[TModule modules.FarseerModule](appName string) error {
	sw := stopwatch.StartNew()

	AppName = appName
	ProcessId = os.Getppid()
	HostName, _ = os.Hostname()
	StartupAt = dateTime.Now()
	rand.Seed(time.Now().UnixNano())
	snowflake.Init(parse.HashCode64(HostName), rand.Int63n(32))
	AppId = snowflake.GenerateId()
	AppIp = net.GetIp()

	flog.Println("AppName： ", flog.Colors[2](AppName))
	flog.Println("AppID：   ", flog.Colors[2](AppId))
	flog.Println("AppIP：   ", flog.Colors[2](AppIp))
	flog.Println("HostName：", flog.Colors[2](HostName))
	flog.Println("HostTime：", flog.Colors[2](StartupAt.ToString("yyyy-MM-dd hh:mm:ss")))
	flog.Println("PID：     ", flog.Colors[2](ProcessId))
	err := showComponentLog()
	if err != nil {
		panic(err)
	}
	flog.Println("---------------------------------------")

	var startupModule TModule
	flog.Println("Loading Module...")
	dependModules = modules.Distinct(modules.GetDependModule(startupModule))
	flog.Println("Loaded, " + strconv.Itoa(len(dependModules)) + " modules in total")
	flog.Println("---------------------------------------")

	modules.StartModules(dependModules)
	flog.Println("---------------------------------------")

	if len(callbackFnList) > 0 {
		for index, fn := range callbackFnList {
			sw.Restart()
			fn()
			flog.Println("Run " + strconv.Itoa(index+1) + "：" + reflect.TypeOf(fn).String() + "，Use：" + sw.GetMillisecondsText())
			flog.Println("---------------------------------------")
		}
	}
	flog.Println("Initialization completed, total time：" + sw.GetMillisecondsText())
	return nil
}

// 组件日志
func showComponentLog() error {
	err := configure.ReadInConfig()
	if err != nil { // 捕获读取中遇到的error
		flog.Errorf("An error occurred while reading: %s \n", err)
		return err
	}

	logConfig := configure.GetSubNodes("Log.Component")
	var logSets []string
	for k, v := range logConfig {
		if v == true {
			logSets = append(logSets, k)
		}
	}
	if len(logSets) > 0 {
		flog.Println("Log Switch：", flog.Colors[2](strings.Join(logSets, " ")))
	}
	return nil
}

// Exit 应用退出
func Exit() {
	modules.ShutdownModules(dependModules)
}

// AddInitCallback 添加框架启动完后执行的函数
func AddInitCallback(fn func()) {
	callbackFnList = append(callbackFnList, fn)
}
