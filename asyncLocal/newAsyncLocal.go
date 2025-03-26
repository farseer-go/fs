package asyncLocal

import (
	"github.com/timandy/routine"
)

// 存储当前线程生成的共享变量，在线程结束后，释放
var lstRemoves routine.ThreadLocal[[]func()] = routine.NewInheritableThreadLocal[[]func()]()

// 当前线程共享的变量
type AsyncLocal[T any] struct {
	threadLocal routine.ThreadLocal[T]
}

// New 创建一个AsyncLocal
func New[T any]() AsyncLocal[T] {
	// 加入到list集合，用于手动GC
	threadLocal := routine.NewInheritableThreadLocal[T]()
	threadLocal.Remove()

	f := lstRemoves.Get()
	f = append(f, func() {
		threadLocal.Remove()
	})
	lstRemoves.Set(f)

	return AsyncLocal[T]{
		threadLocal: threadLocal,
	}
}

// Get 获取值
func (receiver AsyncLocal[T]) Get() T {
	val := receiver.threadLocal.Get()
	return val
	// if val == nil {
	// 	var t T
	// 	return t
	// }
	// return val.(T)
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
	for _, threadLocalRemove := range lstRemoves.Get() {
		threadLocalRemove()
	}
	lstRemoves.Remove()
	routineContext.Remove()
}
