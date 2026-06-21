package trace

import (
	"github.com/farseer-go/fs/asyncLocal"
)

// ScopeLevel 层级列表
var ScopeLevel = asyncLocal.New[[]*TraceDetail]()

// cachedManager 缓存的链路追踪管理器（单例，注册时由 SetManager 写入一次）
var cachedManager IManager

// SetManager 在模块初始化注册 IManager 时调用一次，缓存单例实例。
// 避免热路径每个请求多次走容器查找（原本每请求解析3次）的加锁与开销。
// 放在 trace 包以避免 trace -> container -> flog -> trace 的循环依赖。
func SetManager(manager IManager) {
	cachedManager = manager
}

// Manager 获取缓存的链路追踪管理器（单例）。
// 若尚未通过 SetManager 设置（理论上不会发生），返回空实现兜底。
func Manager() IManager {
	if cachedManager == nil {
		return &EmptyManager{}
	}
	return cachedManager
}

// CurTraceContext 当前请求的Trace上下文
var CurTraceContext = asyncLocal.New[*TraceContext]()
var detailComment = asyncLocal.New[string]()

// SetComment 添加操作的注释
func SetComment(cmt ...string) {
	if len(cmt) > 0 {
		detailComment.Set(cmt[0])
	}
}

// ClearComment 移除注释
func ClearComment() {
	detailComment.Remove()
}

// GetComment 获取操作的注释，并删除
func GetComment() string {
	cmt := detailComment.Get()
	detailComment.Remove()
	return cmt
}

func GetTraceId() string {
	if traceContext := CurTraceContext.Get(); traceContext != nil {
		return traceContext.TraceId
	}
	return ""
}
