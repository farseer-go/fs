package asyncLocal

import (
	"sync"

	"github.com/farseer-go/fs/fastReflect"
)

// 在一次请求中共享数据（适用于多层架构中不同层之间的数据共享，省去传值）
var routineContext AsyncLocal[*sync.Map] = New[*sync.Map]()

// InitContext 初始化同一协程上下文，避免在同一协程中多次初始化
func InitContext() *sync.Map {
	mVal := routineContext.Get()
	if mVal == nil {
		mVal = &sync.Map{}
		routineContext.Set(mVal)
	}
	return mVal
}

// GetContext 获取Context
func GetContext[T any](key string) T {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key

	mVal := routineContext.Get()
	if mVal == nil {
		return t
	}
	if val, exists := mVal.Load(key); exists {
		return val.(T)
	}
	return t
}

func getLock(mVal *sync.Map, key string) *sync.Mutex {
	lockKey := "ContextLock_" + key
	var lock *sync.Mutex
	if val, exists := mVal.Load(lockKey); exists {
		lock = val.(*sync.Mutex)
	} else {
		lock = &sync.Mutex{}
		mVal.Store(lockKey, lock)
	}
	return lock
}

// GetOrSetContext 获取Context
func GetOrSetContext[T any](key string, getValFunc func() T) T {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key

	mVal := InitContext()
	if val, exists := mVal.Load(key); exists {
		return val.(T)
	}

	// 使用锁来确保在单个协程环境下的线程安全
	lock := getLock(mVal, key)
	lock.Lock()
	defer lock.Unlock()

	// 再次检查，避免在锁内重复设置
	if val, exists := mVal.Load(key); exists {
		return val.(T)
	}

	t = getValFunc()
	mVal.Store(key, t)
	routineContext.Set(mVal)
	return t
}

// SetContext 写入上下文
func SetContext[T any](key string, getValFunc func() T) {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key

	mVal := routineContext.Get()
	if mVal == nil {
		mVal = &sync.Map{}
	}
	mVal.Store(key, getValFunc())
	routineContext.Set(mVal)
}

// SetContextIfNotExists 写入上下文（如果不存在）
func SetContextIfNotExists[T any](key string, getValFunc func() T) {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key

	mVal := routineContext.Get()
	if mVal == nil {
		mVal = &sync.Map{}
	}

	if _, exists := mVal.Load(key); !exists {
		mVal.Store(key, getValFunc())
		routineContext.Set(mVal)
	}
}

// Remove 移除缓存
func Remove[T any](key string) {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key
	if mVal := routineContext.Get(); mVal != nil {
		mVal.Delete(key)
		routineContext.Set(mVal)
	}
}
