package trace

import (
	"github.com/farseer-go/fs/asyncLocal"
)

// CurTraceContext 当前请求的Trace上下文
var CurTraceContext = asyncLocal.New[ITraceContext]()
