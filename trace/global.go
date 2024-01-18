package trace

import (
	"github.com/farseer-go/fs/asyncLocal"
)

// CurTraceContext 当前请求的Trace上下文
var CurTraceContext = asyncLocal.New[ITraceContext]()
var DetailComment = asyncLocal.New[string]()

// SetComment 添加操作的注释
func SetComment(cmt string) {
	DetailComment.Set(cmt)
}

// ClearComment 移除注释
func ClearComment() {
	DetailComment.Remove()
}
