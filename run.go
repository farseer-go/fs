package Farseer_Go

import (
	"fmt"
	"github.com/farseernet/farseer.go/utils/net"
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

func Run(appName string) {
	rand.Seed(time.Now().UnixNano())
	StartupAt = time.Now()
	appidString := strings.Join([]string{strconv.FormatInt(StartupAt.UnixMicro(), 10), strconv.Itoa(rand.Intn(999-100) + 100)}, "")
	AppId, _ = strconv.ParseInt(appidString, 10, 64)
	AppIp = net.Ip
	AppName = appName

	fmt.Println("系统时间：", StartupAt)
	fmt.Println("进程ID：", os.Getppid())
	fmt.Println("应用ID：", AppId)
	fmt.Println("应用IP：", AppIp)
	fmt.Println("---------------------------------------")
}
