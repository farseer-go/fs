package asyncLocal

import (
	"github.com/timandy/routine"
)

// 存储当前线程生成的共享变量，在线程结束后，释放
var lstRemoves routine.ThreadLocal[[]func()] = routine.NewInheritableThreadLocal[[]func()]()

// New 创建一个AsyncLocal
func New[T any]() routine.ThreadLocal[T] {
	// 加入到list集合，用于手动GC
	threadLocal := routine.NewInheritableThreadLocal[T]()
	//threadLocal.Remove()

	// 用于后续可以遍历执行清除动作
	f := lstRemoves.Get()
	lstRemoves.Set(append(f, func() {
		threadLocal.Remove()
	}))

	return threadLocal
}

// Release 释放
func Release() {
	// 将所有创建的对象遍历执行Remove()操作
	for _, removeFunc := range lstRemoves.Get() {
		removeFunc()
	}
	lstRemoves.Remove()

	// 清除上下存储的数据缓存(如在不同的分层中,从数据库查询到的数据)
	routineContext.Remove()
}
