package asyncLocal

import (
	"github.com/timandy/routine"
)

type AsyncLocal[T any] struct {
	threadLocal routine.ThreadLocal
}

// 将同一个协程所New的AsyncLocal放到同一个数组中，在使用完后释放
var list = make(map[int64][]routine.ThreadLocal)

// New 创建一个AsyncLocal
func New[T any]() AsyncLocal[T] {
	// 加入到list集合，用于手动GC
	al := AsyncLocal[T]{
		threadLocal: routine.NewInheritableThreadLocal(),
	}
	al.AddRelease()
	return al
}

// AddRelease 将同一个协程所New的AsyncLocal放到同一个数组中，在使用完后自动释放
func (receiver AsyncLocal[T]) AddRelease() {
	goId := routine.Goid()
	list[goId] = append(list[goId], receiver.threadLocal)
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

// Release 释放当前使用的AsyncLocal
func Release() {
	goId := routine.Goid()
	for _, threadLocal := range list[goId] {
		threadLocal.Remove()
	}
	delete(list, goId)
}
