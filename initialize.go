package fs

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/farseer-go/fs/color"
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/net"
	"github.com/farseer-go/fs/sonyflake"
	"github.com/farseer-go/fs/stopwatch"
)

var (
	Context       context.Context         // 最顶层的上下文
	contextCancel context.CancelFunc      // 最顶层的上下文
	dependModules []modules.FarseerModule // 依赖的模块
	isInit        bool                    // 是否初始化完
)

var onceInit sync.Once

// Initialize 初始化框架
func Initialize[TModule modules.FarseerModule](appName string) {
	sw := stopwatch.StartNew()
	Context, contextCancel = context.WithCancel(context.Background())
	onceInit.Do(func() {
		core.AppName = appName
		core.ProcessId = os.Getppid()
		core.HostName, _ = os.Hostname()
		core.StartupAt = dateTime.Now()
		core.AppId = sonyflake.GenerateId()
		core.AppIp = net.GetIp()
	})
	flog.LogBuffer <- fmt.Sprint("FarsVer: ", color.Colors[2](core.Version))
	flog.LogBuffer <- fmt.Sprint("AppName: ", color.Colors[2](core.AppName))
	flog.LogBuffer <- fmt.Sprint("AppID:   ", color.Colors[2](core.AppId))
	flog.LogBuffer <- fmt.Sprint("AppIP:   ", color.Colors[2](core.AppIp))
	flog.LogBuffer <- fmt.Sprint("HostName:", color.Colors[2](core.HostName))
	flog.LogBuffer <- fmt.Sprint("HostTime:", color.Colors[2](core.StartupAt.ToString("yyyy-MM-dd hh:mm:ss")))
	flog.LogBuffer <- fmt.Sprint("PID:     ", color.Colors[2](core.ProcessId))
	showComponentLog()
	flog.LogBuffer <- "---------------------------------------"

	// 加载模块依赖
	var startupModule TModule
	dependModules = modules.GetDependModules(startupModule)

	var moduleNames []string
	for _, farseerModule := range dependModules {
		moduleName := reflect.TypeOf(farseerModule).String()
		moduleNames = append(moduleNames, color.Colors[5](moduleName))
	}
	flog.LogBuffer <- fmt.Sprint("Loading Module: " + strings.Join(moduleNames, "->") + " (" + color.Red(len(dependModules)) + " modules) ")
	//flog.LogBuffer <- fmt.Sprint("Loaded, " + color.Red(len(dependModules)) + " modules in total")

	// 执行所有模块初始化
	modules.StartModules(dependModules)

	flog.CloseBuffer()
	flog.Println("Initialization completed, total time: " + sw.GetMillisecondsText())
	flog.Println("---------------------------------------")

	if proxy := configure.GetString("Proxy"); proxy != "" {
		flog.Println("http使用代理: ", color.Blue(proxy))
	}

	// 健康检查
	healthChecks := container.ResolveAll[core.IHealthCheck]()
	if len(healthChecks) > 0 {
		flog.Println("Health Check...")
		isSuccess := true
		for _, healthCheck := range healthChecks {
			item, err := healthCheck.Check()
			if err == nil {
				flog.Printf("%s%s\n", color.Green("【✓】"), item)
			} else {
				flog.Printf("%s%s：%s\n", color.Red("【✕】"), item, color.Red(err.Error()))
				isSuccess = false
			}
		}
		if !isSuccess {
			//os.Exit(-1)
			panic("健康检查失败")
		}
	}

	// 日志内容美化
	if len(healthChecks) > 0 {
		flog.Println("---------------------------------------")
	}

	isInit = true
	// 加载callbackFnList，启动后才执行的模块
	if len(core.CallbackInitList) > 0 {
		for index, fn := range core.CallbackInitList {
			sw.Restart()
			fn.F()
			flog.Println("Run " + strconv.Itoa(index+1) + ": " + fn.Name + ", Use: " + sw.GetText())
		}
		flog.Println("---------------------------------------")
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
		flog.LogBuffer <- fmt.Sprint("Log Switch: ", color.Colors[2](strings.Join(logSets, " ")))
	}
}

// Exit 应用退出
func Exit() {
	contextCancel()
	// 关闭模块
	modules.ShutdownModules(dependModules)
}

// AddInitCallback 添加框架启动完后执行的函数
func AddInitCallback(name string, fn func()) {
	// 未初始化完时，加入到列表中
	if !isInit {
		core.AddInitCallback(name, fn)
	} else { // 初始化完后，则立即执行
		sw := stopwatch.StartNew()
		fn()
		flog.Println("Run : " + name + ", Use: " + sw.GetMillisecondsText())
	}
}

// Run 运行应用，等待退出信号
func Run() {
	flog.Println("应用已成功启动!")
	// 注册退出信号
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-exitSignal:
		// 收到信号
	case <-Context.Done():
		// 被 Stop 调用
	}

	// 关闭模块
	modules.ShutdownModules(dependModules)
}
