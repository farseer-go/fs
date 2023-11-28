package asyncLocal

import (
	"github.com/timandy/routine"
)

type AsyncLocal[T any] struct {
	threadLocal routine.ThreadLocal
}

// New 创建一个AsyncLocal
func New[T any]() AsyncLocal[T] {
	return AsyncLocal[T]{
		threadLocal: routine.NewInheritableThreadLocal(),
	}
}

// Get 获取值
func (receiver AsyncLocal[T]) Get() T {
	val := receiver.threadLocal.Get()
	if val == nil {
		var t T
		return t
	}
	return val
}

// Set 设置值
func (receiver AsyncLocal[T]) Set(t T) {
	receiver.threadLocal.Set(t)
}

// Remove 移除对象
func (receiver AsyncLocal[T]) Remove() {
	receiver.threadLocal.Remove()
}
