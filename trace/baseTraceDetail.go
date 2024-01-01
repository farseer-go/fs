package trace

import (
	"fmt"
	"github.com/farseer-go/fs/asyncLocal"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/path"
	"github.com/farseer-go/fs/trace/eumCallType"
	"runtime"
	"strings"
	"time"
)

// ScopeLevel 层级列表
var ScopeLevel = asyncLocal.New[[]BaseTraceDetail]()

// BaseTraceDetail 埋点明细（基类）
type BaseTraceDetail struct {
	TraceIdN       string            // 上下文ID
	TraceId        int64             // 上下文ID
	AppId          int64             // 应用ID
	AppName        string            // 应用名称
	AppIp          string            // 应用IP
	ParentAppName  string            // 上游应用
	DetailId       int64             // 明细ID
	ParentDetailId int64             // 父级明细ID
	Level          int               // 当前层级（入口为0层）
	MethodName     string            // 调用方法
	CallType       eumCallType.Enum  // 调用类型
	Timeline       time.Duration     // 从入口开始统计（微秒）
	UnTraceTs      time.Duration     // 上一次结束到现在开始之间未Trace的时间（微秒）
	StartTs        int64             // 调用开始时间戳（微秒）
	EndTs          int64             // 调用停止时间戳（微秒）
	UseTs          time.Duration     // 总共使用时间微秒
	UseDesc        string            // 总共使用时间（描述）
	ignore         bool              // 忽略这次的链路追踪
	Exception      *ExceptionStack   // 异常信息
	CreateAt       dateTime.DateTime // 请求时间
}

type ExceptionStack struct {
	ExceptionCallFile     string // 调用者文件路径
	ExceptionCallLine     int    // 调用者行号
	ExceptionCallFuncName string // 调用者函数名称
	ExceptionIsException  bool   // 是否执行异常
	ExceptionMessage      string // 异常信息
}

func (receiver ExceptionStack) IsNil() bool {
	return receiver.ExceptionCallFile == "" && receiver.ExceptionCallLine == 0 && receiver.ExceptionCallFuncName == "" && receiver.ExceptionIsException == false && receiver.ExceptionMessage == ""
}

func (receiver *BaseTraceDetail) SetSql(connectionString string, dbName string, tableName string, sql string, rowsAffected int64) {
}
func (receiver *BaseTraceDetail) SetHttpRequest(url string, head map[string]any, requestBody string, responseBody string, statusCode int) {
}

// End 链路明细执行完后，统计用时
func (receiver *BaseTraceDetail) End(err error) {
	receiver.EndTs = time.Now().UnixMicro()
	receiver.UseTs = time.Duration(receiver.EndTs-receiver.StartTs) * time.Microsecond

	if err != nil {
		receiver.Exception = &ExceptionStack{
			ExceptionIsException: true,
			ExceptionMessage:     err.Error(),
		}
		// 调用者
		receiver.Exception.ExceptionCallFile, receiver.Exception.ExceptionCallFuncName, receiver.Exception.ExceptionCallLine = GetCallerInfo()
	}

	// 移除层级
	lstScope := ScopeLevel.Get()
	if len(lstScope) > 0 {
		ScopeLevel.Set(lstScope[:len(lstScope)-1])
	}
}

func (receiver *BaseTraceDetail) Ignore() {
	receiver.ignore = true
}
func (receiver *BaseTraceDetail) IsIgnore() bool {
	return receiver.ignore
}
func (receiver *BaseTraceDetail) GetLevel() int {
	return receiver.Level
}
func (receiver *BaseTraceDetail) SetTraceIdN() {
	receiver.TraceIdN = parse.ToString(receiver.TraceId)
}

var ComNames = []string{"/farseer-go/async/", "/farseer-go/cache/", "/farseer-go/cacheMemory/", "/farseer-go/collections/", "/farseer-go/data/", "/farseer-go/elasticSearch/", "/farseer-go/etcd/", "/farseer-go/eventBus/", "/farseer-go/fs/", "/farseer-go/linkTrace/", "/farseer-go/mapper/", "/farseer-go/queue/", "/farseer-go/rabbit/", "/farseer-go/redis/", "/farseer-go/redisStream/", "/farseer-go/tasks/", "/farseer-go/utils/", "/farseer-go/webapi/", "/src/reflect/", "/usr/local/go/src/", "gorm.io/"}

func IsSysCom(file string) bool {
	for _, comName := range ComNames {
		if strings.Contains(file, comName) {
			return true
		}
	}
	return false
}

func GetCallerInfo() (string, string, int) {
	// 获取调用栈信息
	pc := make([]uintptr, 15) // 假设最多获取 10 层调用栈
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])

	// 遍历调用栈帧
	for {
		frame, more := frames.Next()
		if !strings.HasSuffix(frame.File, "_test.go") && (!IsSysCom(frame.File) || strings.HasSuffix(frame.File, "healthCheck.go")) { // !strings.HasPrefix(file, gormSourceDir) ||
			frame.Function = strings.TrimPrefix(frame.Function, "github.com/")
			// 移除绝对路径
			prefixFunc := frame.Function[0 : strings.LastIndex(frame.Function, path.PathSymbol)+len(path.PathSymbol)]
			packageIndex := strings.Index(frame.File, prefixFunc)
			if packageIndex == -1 {
				fmt.Print(packageIndex)
			}
			file := frame.File[packageIndex:]

			// 只要最后的方法名
			funcName := frame.Function[strings.LastIndex(frame.Function, path.PathSymbol)+len(path.PathSymbol):] + "()"
			return file, funcName, frame.Line
		}
		if !more {
			break
		}
	}
	return "", "", 0
}
