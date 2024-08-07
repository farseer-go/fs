package asyncLocal

import (
	"github.com/timandy/routine"
)

var list []routine.ThreadLocal

type AsyncLocal[T any] struct {
	threadLocal routine.ThreadLocal
}

// New 创建一个AsyncLocal
func New[T any]() AsyncLocal[T] {
	// 加入到list集合，用于手动GC
	threadLocal := routine.NewInheritableThreadLocal()
	list = append(list, threadLocal)

	return AsyncLocal[T]{
		threadLocal: threadLocal,
	}
}

// Get 获取值
func (receiver AsyncLocal[T]) Get() T {
	val := receiver.threadLocal.Get()
	if val == nil {
		var t T
		return t
	}
	return val.(T)
}

// Set 设置值
func (receiver AsyncLocal[T]) Set(t T) {
	receiver.threadLocal.Set(t)
}

// Remove 移除对象
func (receiver AsyncLocal[T]) Remove() {
	receiver.threadLocal.Remove()
}

// Release 释放
func Release() {
	for _, threadLocal := range list {
		threadLocal.Remove()
	}
	routineContext.Remove()
}
