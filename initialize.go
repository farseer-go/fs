package fs

import (
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/net"
	"github.com/farseer-go/fs/stopwatch"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// StartupAt 应用启动时间
var StartupAt time.Time

// AppName 应用名称
var AppName string

// AppId 应用ID
var AppId int64

// AppIp 应用IP
var AppIp string

// 依赖的模块
var dependModules []modules.FarseerModule

// Initialize 初始化框架
func Initialize[TModule modules.FarseerModule](appName string) {
	sw := stopwatch.StartNew()
	rand.Seed(time.Now().UnixNano())
	StartupAt = time.Now()
	appidString := strings.Join([]string{strconv.FormatInt(StartupAt.UnixMilli(), 10), strconv.Itoa(rand.Intn(999-100) + 100)}, "")
	AppId, _ = strconv.ParseInt(appidString, 10, 64)
	AppIp = net.Ip
	AppName = appName

	flog.Println("应用名称：", AppName)
	flog.Println("系统时间：", StartupAt)
	flog.Println("进程ID：", os.Getppid())
	flog.Println("应用ID：", AppId)
	flog.Println("应用IP：", AppIp)
	flog.Println("---------------------------------------")

	var startupModule TModule
	flog.Println("加载模块...")
	dependModules = modules.GetDependModule(startupModule)
	flog.Println("加载完毕，共加载 " + strconv.Itoa(len(dependModules)) + " 个模块")
	flog.Println("---------------------------------------")

	modules.StartModules(dependModules)
	flog.Println("初始化完毕，共耗时" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms")
	flog.Println("---------------------------------------")
}

// Exit 应用退出
func Exit() {
	modules.ShutdownModules(dependModules)
}
