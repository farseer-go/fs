package fs

import (
	"context"
	"fmt"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/net"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/snowflake"
	"github.com/farseer-go/fs/stopwatch"
	"math/rand"
	"os"
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

// Context 最顶层的上下文
var Context context.Context

// Log 日志
var Log core.ILog

// 依赖的模块
var dependModules []modules.FarseerModule

// 回调函数列表
var callbackFnList []callbackFn

type callbackFn struct {
	f    func()
	name string
}

// Initialize 初始化框架
func Initialize[TModule modules.FarseerModule](appName string) {
	sw := stopwatch.StartNew()
	Context = context.Background()
	AppName = appName
	ProcessId = os.Getppid()
	HostName, _ = os.Hostname()
	StartupAt = dateTime.Now()
	rand.Seed(time.Now().UnixNano())

	snowflake.Init(parse.HashCode64(HostName), rand.Int63n(32))
	AppId = snowflake.GenerateId()
	AppIp = net.GetIp()

	flog.LogBuffer <- fmt.Sprint("AppName： ", flog.Colors[2](AppName))
	flog.LogBuffer <- fmt.Sprint("AppID：   ", flog.Colors[2](AppId))
	flog.LogBuffer <- fmt.Sprint("AppIP：   ", flog.Colors[2](AppIp))
	flog.LogBuffer <- fmt.Sprint("HostName：", flog.Colors[2](HostName))
	flog.LogBuffer <- fmt.Sprint("HostTime：", flog.Colors[2](StartupAt.ToString("yyyy-MM-dd hh:mm:ss")))
	flog.LogBuffer <- fmt.Sprint("PID：     ", flog.Colors[2](ProcessId))
	showComponentLog()
	flog.LogBuffer <- fmt.Sprint("---------------------------------------")

	var startupModule TModule
	dependModules = modules.Distinct(modules.GetDependModule(startupModule))
	flog.LogBuffer <- fmt.Sprint("Loaded, " + flog.Red(len(dependModules)) + " modules in total")

	modules.StartModules(dependModules)
	flog.LogBuffer <- fmt.Sprint("Initialization completed, total time：" + sw.GetMillisecondsText())
	flog.LogBuffer <- fmt.Sprint("---------------------------------------")

	Log = container.Resolve[core.ILog]()
	flog.ClearLogBuffer(Log)
	go flog.LoadLogBuffer(Log)

	// 健康检查
	healthChecks := container.ResolveAll[core.IHealthCheck]()
	if len(healthChecks) > 0 {
		flog.LogBuffer <- fmt.Sprint("Health Check...")
		isSuccess := true
		for _, healthCheck := range healthChecks {
			item, err := healthCheck.Check()
			if err == nil {
				flog.LogBuffer <- fmt.Sprintf("%s%s", flog.Green("【✓】"), item)
			} else {
				flog.LogBuffer <- fmt.Sprintf("%s%s：%s", flog.Red("【✕】"), item, flog.Red(err.Error()))
				isSuccess = false
			}
		}
		flog.LogBuffer <- fmt.Sprint("---------------------------------------")

		if !isSuccess {
			//os.Exit(-1)
			panic("健康检查失败")
		}
	}

	// 加载callbackFnList，启动后才执行的模块
	if len(callbackFnList) > 0 {
		for index, fn := range callbackFnList {
			sw.Restart()
			fn.f()
			flog.LogBuffer <- fmt.Sprint("Run " + strconv.Itoa(index+1) + "：" + fn.name + "，Use：" + sw.GetMillisecondsText())
		}
		flog.LogBuffer <- fmt.Sprint("---------------------------------------")
	}
}

// 组件日志
func showComponentLog() {
	logConfig := configure.GetSubNodes("Log.Component")
	var logSets []string
	for k, v := range logConfig {
		if v == true {
			logSets = append(logSets, k)
		}
	}
	if len(logSets) > 0 {
		flog.LogBuffer <- fmt.Sprint("Log Switch：", flog.Colors[2](strings.Join(logSets, " ")))
	}
}

// Exit 应用退出
func Exit(code int) {
	modules.ShutdownModules(dependModules)
	os.Exit(code)
}

// AddInitCallback 添加框架启动完后执行的函数
func AddInitCallback(name string, fn func()) {
	callbackFnList = append(callbackFnList, callbackFn{name: name, f: fn})
}
