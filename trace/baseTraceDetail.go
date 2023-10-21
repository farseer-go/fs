package trace

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/asyncLocal"
	"github.com/farseer-go/linkTrace/eumCallType"
	"time"
)

// ScopeLevel 层级列表
var ScopeLevel = asyncLocal.New[collections.List[BaseTraceDetail]]()

// BaseTraceDetail 埋点明细（基类）
type BaseTraceDetail struct {
	DetailId         int64            // 明细ID
	ParentDetailId   int64            // 父级明细ID
	Level            int              // 当前层级（入口为0层）
	CallMethod       string           // 调用方法
	CallType         eumCallType.Enum // 调用类型
	Timeline         time.Duration    // 从入口开始统计
	UnTraceTs        time.Duration    // 上一次结束到现在开始之间未Trace的时间
	StartTs          int64            // 调用开始时间戳
	EndTs            int64            // 调用停止时间戳
	UseTs            time.Duration    // 总共使用时间毫秒
	IsException      bool             // 是否执行异常
	ExceptionMessage string           // 异常信息
	ignore           bool             // 忽略这次的链路追踪
}

func (receiver *BaseTraceDetail) SetSql(DbName string, tableName string, sql string) {}

// End 链路明细执行完后，统计用时
func (receiver *BaseTraceDetail) End(err error) {
	receiver.EndTs = time.Now().UnixMicro()
	receiver.UseTs = time.Duration(receiver.EndTs-receiver.StartTs) * time.Microsecond

	if err != nil {
		receiver.IsException = true
		receiver.ExceptionMessage = err.Error()
	}

	// 移除层级
	lstScope := ScopeLevel.Get()
	if !lstScope.IsNil() {
		lstScope.RemoveAt(lstScope.Count() - 1)
		ScopeLevel.Set(lstScope)
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
