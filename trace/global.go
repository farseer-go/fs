package trace

import (
	"github.com/farseer-go/fs/asyncLocal"
)

// ScopeLevel 层级列表
var ScopeLevel = asyncLocal.New[[]BaseTraceDetail]()

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
