package fs

import (
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/net"
	"github.com/farseer-go/fs/stopwatch"
	"log"
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

// Initialize 初始化框架
func Initialize[TModule modules.FarseerModule](appName string) {
	sw := stopwatch.StartNew()
	rand.Seed(time.Now().UnixNano())
	StartupAt = time.Now()
	appidString := strings.Join([]string{strconv.FormatInt(StartupAt.UnixMicro(), 10), strconv.Itoa(rand.Intn(999-100) + 100)}, "")
	AppId, _ = strconv.ParseInt(appidString, 10, 64)
	AppIp = net.Ip
	AppName = appName

	log.Println("系统时间：", StartupAt)
	log.Println("进程ID：", os.Getppid())
	log.Println("应用ID：", AppId)
	log.Println("应用IP：", AppIp)
	log.Println("---------------------------------------")

	var startupModule TModule
	modules.StartModules(startupModule)
	log.Println("初始化完毕，共耗时" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms")
	log.Println("---------------------------------------")
}
