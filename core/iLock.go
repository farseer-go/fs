package core

// ILock 锁
type ILock interface {
	// TryLock 尝试加锁
	TryLock() bool
	// GetLock 获取锁，直到获取成功
	GetLock()
	// ReleaseLock 释放锁
	ReleaseLock()
	// TryLockRun 尝试加锁，执行完后，自动释放锁（未获取到时直接退出）
	TryLockRun(fn func()) bool
	// GetLockRun 获取锁，直到获取成功，执行完后，自动释放锁
	GetLockRun(fn func())
}
