package fs

import (
	"context"
	"fmt"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/fops"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/net"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/snowflake"
	"github.com/farseer-go/fs/sonyflake"
	"github.com/farseer-go/fs/stopwatch"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	Context        context.Context         // 最顶层的上下文
	dependModules  []modules.FarseerModule // 依赖的模块
	callbackFnList []callbackFn            // 回调函数列表
	isInit         bool                    // 是否初始化完
)

type callbackFn struct {
	f    func()
	name string
}

var onceInit sync.Once

// Initialize 初始化框架
func Initialize[TModule modules.FarseerModule](appName string) {
	sw := stopwatch.StartNew()
	Context = context.Background()
	onceInit.Do(func() {
		rand.Seed(time.Now().UnixNano())
		core.AppName = appName
		core.ProcessId = os.Getppid()
		core.HostName, _ = os.Hostname()
		core.StartupAt = dateTime.Now()
		snowflake.Init(parse.HashCode64(core.HostName), rand.Int63n(32))
		core.AppId = sonyflake.GenerateId()
		core.AppIp = net.GetIp()
	})

	flog.LogBuffer <- fmt.Sprint("AppName： ", flog.Colors[2](core.AppName))
	flog.LogBuffer <- fmt.Sprint("AppID：   ", flog.Colors[2](core.AppId))
	flog.LogBuffer <- fmt.Sprint("AppIP：   ", flog.Colors[2](core.AppIp))
	flog.LogBuffer <- fmt.Sprint("HostName：", flog.Colors[2](core.HostName))
	flog.LogBuffer <- fmt.Sprint("HostTime：", flog.Colors[2](core.StartupAt.ToString("yyyy-MM-dd hh:mm:ss")))
	flog.LogBuffer <- fmt.Sprint("PID：     ", flog.Colors[2](core.ProcessId))
	showComponentLog()
	flog.LogBuffer <- fmt.Sprint("---------------------------------------")

	// 加载模块依赖
	var startupModule TModule
	dependModules = modules.GetDependModules(startupModule)
	flog.LogBuffer <- fmt.Sprint("Loaded, " + flog.Red(len(dependModules)) + " modules in total")

	// 执行所有模块初始化
	modules.StartModules(dependModules)
	flog.LogBuffer <- fmt.Sprint("Initialization completed, total time：" + sw.GetMillisecondsText())
	flog.LogBuffer <- fmt.Sprint("---------------------------------------")

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
		if !isSuccess {
			//os.Exit(-1)
			panic("健康检查失败")
		}
	}

	// 定时向FOPS中心注册应用信息
	fops.RegisterApp()
	// 日志内容美化
	if len(healthChecks) > 0 || configure.GetString("Fops.Server") != "" {
		flog.LogBuffer <- fmt.Sprint("---------------------------------------")
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

	isInit = true
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
	// 未初始化完时，加入到列表中
	if !isInit {
		callbackFnList = append(callbackFnList, callbackFn{name: name, f: fn})
	} else { // 初始化完后，则立即执行
		fn()
	}
}
