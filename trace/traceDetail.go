package trace

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/farseer-go/fs/color"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/path"
	"github.com/farseer-go/fs/sonyflake"
	"github.com/farseer-go/fs/trace/eumCallType"
)

// TraceDetail 埋点明细（基类）
type TraceDetail struct {
	// TraceId    string // 上下文ID
	// AppId          string            // 应用ID
	// AppName        string            // 应用名称
	// AppIp          string            // 应用IP
	// ParentAppName  string            // 上游应用
	TraceLevel     int               // 逐层递增（显示上下游顺序）
	DetailId       string            // 明细ID
	ParentDetailId string            // 父级明细ID
	Level          int               // 当前层级（入口为0层）
	Comment        string            // 调用注释
	MethodName     string            // 调用方法
	CallType       eumCallType.Enum  // 调用类型
	Timeline       time.Duration     // 从入口开始统计（微秒）
	UnTraceTs      time.Duration     // 上一次结束到现在开始之间未Trace的时间（微秒）
	StartTs        int64             // 调用开始时间戳（微秒）
	EndTs          int64             // 调用停止时间戳（微秒）
	ignore         bool              // 忽略这次的链路追踪
	Exception      *ExceptionStack   // 异常信息
	CreateAt       dateTime.DateTime // 请求时间
	TraceDetailHand
	TraceDetailDatabase
	TraceDetailEs
	TraceDetailEtcd
	TraceDetailEvent
	TraceDetailGrpc
	TraceDetailHttp
	TraceDetailMq
	TraceDetailRedis
	// UseTs          time.Duration     // 总共使用时间微秒
	// UseDesc        string            // 总共使用时间（描述）
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

func (receiver *TraceDetail) SetSql(connectionString string, dbName string, tableName string, sql string, rowsAffected int64) {
}
func (receiver *TraceDetail) SetHttpRequest(url string, reqHead map[string]any, rspHead map[string]string, requestBody string, responseBody string, statusCode int) {
}
func (receiver *TraceDetail) SetRows(rows int) {
}

// End 链路明细执行完后，统计用时
func (receiver *TraceDetail) End(err error) {
	receiver.EndTs = time.Now().UnixMicro()
	// useTs = time.Duration(receiver.EndTs-receiver.StartTs) * time.Microsecond
	// receiver.UseDesc = useTs.String()

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

func (receiver *TraceDetail) Ignore() {
	receiver.ignore = true
}
func (receiver *TraceDetail) IsIgnore() bool {
	return receiver.ignore
}

func (receiver *TraceDetail) Run(fn func()) {
	defer func() {
		exp := recover()
		if exp != nil {
			receiver.End(fmt.Errorf("%s", exp))
			panic(exp)
		}
	}()
	fn()
	receiver.End(nil)
}

func (receiver *TraceDetail) ToString() string {
	useTs := time.Duration(receiver.EndTs-receiver.StartTs) * time.Microsecond
	switch receiver.CallType {
	case eumCallType.Database:
		if sql := receiver.TraceDetailDatabase.DbSql; sql != "" {
			if len(sql) > 1000 {
				sql = sql[:1000] + "......"
			}
			sql = color.ReplaceBlues(sql, "SELECT ", "UPDATE ", "DELETE ", "INSERT INTO ", " FROM ", " WHERE ", " LIMIT ", " SET ", " ORDER BY ", " VALUES ", " and ", " or ", "`")
			sql = strings.ReplaceAll(sql, receiver.TraceDetailDatabase.DbTableName, color.Green(receiver.TraceDetailDatabase.DbTableName))
			return fmt.Sprintf("[%s]耗时：%s，[影响%d行]%s", color.Yellow(receiver.CallType.ToString()), color.Red(useTs.String()), receiver.TraceDetailDatabase.DbRowsAffected, sql)
		} else if receiver.TraceDetailDatabase.DbConnectionString != "" {
			return fmt.Sprintf("[%s]耗时：%s， 连接数据库：%s", color.Yellow(receiver.CallType.ToString()), color.Red(useTs.String()), receiver.TraceDetailDatabase.DbConnectionString)
		}
		return ""
	case eumCallType.Elasticsearch:
		return fmt.Sprintf("[%s]耗时：%s，%s IndexName=%s，AliasesName=%s", color.Yellow(receiver.CallType.ToString()), color.Red(useTs.String()), receiver.MethodName, receiver.TraceDetailEs.EsIndexName, receiver.TraceDetailEs.EsAliasesName)
	case eumCallType.Etcd:
		return fmt.Sprintf("[%s]耗时：%s，%s Key=%s LeaseID=%v", color.Yellow(receiver.CallType.ToString()), color.Red(useTs.String()), receiver.MethodName, receiver.TraceDetailEtcd.EtcdKey, receiver.EtcdLeaseID)
	case eumCallType.EventPublish:
		return fmt.Sprintf("[%s]耗时：%s， %s", color.Yellow(receiver.CallType.ToString()), color.Red(useTs.String()), receiver.TraceDetailEvent.EventName)
	case eumCallType.Grpc:
		return fmt.Sprintf("[%s]耗时：%s，%s [%s]%s", color.Yellow(receiver.CallType.ToString()), color.Red(useTs.String()), receiver.MethodName, receiver.TraceDetailGrpc.GrpcMethod, receiver.TraceDetailGrpc.GrpcUrl)
	case eumCallType.Http:
		return fmt.Sprintf("[%s]耗时：%s，%s [%s]%s", color.Yellow(receiver.CallType.ToString()), color.Red(useTs.String()), receiver.MethodName, receiver.TraceDetailHttp.HttpMethod, receiver.TraceDetailHttp.HttpUrl)
	case eumCallType.Mq:
		return fmt.Sprintf("[%s]耗时：%s，%s Server=%s，Exchange=%s，RoutingKey=%s", color.Yellow(receiver.CallType.ToString()), color.Red(useTs.String()), receiver.MethodName, receiver.TraceDetailMq.MqServer, receiver.TraceDetailMq.MqExchange, receiver.TraceDetailMq.MqRoutingKey)
	case eumCallType.Redis:
		return fmt.Sprintf("[%s]耗时：%s，%s Key=%s，Field=%s", color.Yellow(receiver.CallType.ToString()), color.Red(useTs.String()), receiver.MethodName, receiver.TraceDetailRedis.RedisKey, receiver.TraceDetailRedis.RedisField)
	}
	return fmt.Sprintf("[%s]耗时：%s， %s", receiver.CallType.ToString(), useTs.String(), receiver.MethodName)
}

var ComNames = []string{"/farseer-go/async/", "/farseer-go/cache/", "/farseer-go/cacheMemory/", "/farseer-go/collections/", "/farseer-go/data/", "/farseer-go/elasticSearch/", "/farseer-go/etcd/", "/farseer-go/eventBus/", "/farseer-go/fs/", "/farseer-go/fSchedule/", "/farseer-go/linkTrace/", "/farseer-go/mapper/", "/farseer-go/queue/", "/farseer-go/rabbit/", "/farseer-go/redis/", "/farseer-go/redisStream/", "/farseer-go/tasks/", "/farseer-go/utils/", "/farseer-go/webapi/", "/src/reflect/", "/usr/local/go/src/", "gorm.io/"}

func isSysCom(file string) bool {
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
		if !strings.HasSuffix(frame.File, "_test.go") && (!isSysCom(frame.File) || strings.HasSuffix(frame.File, "healthCheck.go")) { // !strings.HasPrefix(file, gormSourceDir) ||
			frame.Function = strings.TrimPrefix(frame.Function, "github.com/")
			// 移除绝对路径
			prefixFunc := frame.Function[0 : strings.LastIndex(frame.Function, path.PathSymbol)+len(path.PathSymbol)]
			packageIndex := strings.Index(strings.ToLower(frame.File), strings.ToLower(prefixFunc))
			if packageIndex == -1 {
				packageIndex = 0 // 未找到，则使用0作为起始位置
			}

			var file string
			if len(frame.File) > packageIndex {
				file = frame.File[packageIndex:]
			}

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

func NewTraceDetail(callType eumCallType.Enum, methodName string) TraceDetail {
	// 获取当前层级列表
	lstScope := ScopeLevel.Get()
	var parentDetailId string
	if len(lstScope) > 0 {
		parentDetailId = lstScope[len(lstScope)-1].DetailId
	}
	baseTraceDetail := TraceDetail{
		DetailId:       strconv.FormatInt(sonyflake.GenerateId(), 10),
		Level:          len(lstScope) + 1,
		ParentDetailId: parentDetailId,
		MethodName:     methodName,
		CallType:       callType,
		StartTs:        time.Now().UnixMicro(),
		EndTs:          time.Now().UnixMicro(),
		Comment:        GetComment(),
		CreateAt:       dateTime.Now(),
	}
	// 加入到当前层级列表
	ScopeLevel.Set(append(lstScope, baseTraceDetail))
	return baseTraceDetail
}
