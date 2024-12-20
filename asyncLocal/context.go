package asyncLocal

import (
	"github.com/farseer-go/fs/fastReflect"
)

// 在一次请求中共享数据（适用于多层架构中不同层之间的数据共享，省去传值）
var routineContext = New[map[string]any]()

// GetContext 获取Context
func GetContext[T any](key string) T {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key

	mVal := routineContext.Get()
	if mVal == nil {
		return t
	}
	if val, exists := mVal[key]; exists {
		return val.(T)
	}
	return t
}

// GetOrSetContext 获取Context
func GetOrSetContext[T any](key string, getValFunc func() T) T {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key

	mVal := routineContext.Get()
	if mVal == nil {
		mVal = make(map[string]any)
	}
	if val, exists := mVal[key]; exists {
		return val.(T)
	}

	t = getValFunc()
	mVal[key] = t
	routineContext.Set(mVal)
	return t
}

// SetContext 写入上下文
func SetContext[T any](key string, getValFunc func() T) {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key

	mVal := routineContext.Get()
	if mVal == nil {
		mVal = make(map[string]any)
	}
	mVal[key] = getValFunc()
	routineContext.Set(mVal)
}

// SetContextIfNotExists 写入上下文（如果不存在）
func SetContextIfNotExists[T any](key string, getValFunc func() T) {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key

	mVal := routineContext.Get()
	if mVal == nil {
		mVal = make(map[string]any)
	}

	if _, exists := mVal[key]; !exists {
		mVal[key] = getValFunc()
		routineContext.Set(mVal)
	}
}

// Remove 移除缓存
func Remove[T any](key string) {
	var t T
	key = fastReflect.PointerOf(t).ReflectTypeString + "_" + key
	if mVal := routineContext.Get(); mVal != nil {
		delete(mVal, key)
		routineContext.Set(mVal)
	}
}
