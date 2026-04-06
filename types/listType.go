package types

import (
	"reflect"
	"sync"
)

// string key=自定义标识
// value int=field or method对应的索引
var Cache = make(map[string][]int)
var lock sync.RWMutex

// methodIndexCache 以 reflect.Type + 后缀 作为 key，避免昂贵的 lstValue.String() 调用
// key: reflect.Type（类型指针，可直接比较），value: int（方法索引）
type methodKey struct {
	t   reflect.Type
	tag uint8 // 0=Add, 1=ToArray, 2=ToArrayAny, 3=New, 4=MapperInit
}

var methodCache sync.Map // map[methodKey]int

const (
	tagAdd         uint8 = 0
	tagToArray     uint8 = 1
	tagToArrayAny  uint8 = 2
	tagNew         uint8 = 3
	tagMapperInit  uint8 = 4
)

func getMethodIndex(t reflect.Type, tag uint8, methodName string) int {
	key := methodKey{t: t, tag: tag}
	if v, ok := methodCache.Load(key); ok {
		return v.(int)
	}
	method, ok := t.MethodByName(methodName)
	if !ok {
		// 方法不在值类型方法集中（可能是指针接收器方法），返回 -1 表示未找到
		methodCache.Store(key, -1)
		return -1
	}
	methodCache.Store(key, method.Index)
	return method.Index
}

// resolveMethod 获取 lstValue 上 methodName 对应的 reflect.Value（Method）。
// 若值类型方法集中找不到（指针接收器方法），则自动取指针后重试。
func resolveMethod(lstValue reflect.Value, tag uint8, methodName string) reflect.Value {
	t := lstValue.Type()
	idx := getMethodIndex(t, tag, methodName)
	if idx >= 0 {
		return lstValue.Method(idx)
	}
	// 值类型没有此方法，尝试指针接收器
	if lstValue.CanAddr() {
		ptrVal := lstValue.Addr()
		pt := ptrVal.Type()
		idx = getMethodIndex(pt, tag, methodName)
		return ptrVal.Method(idx)
	}
	// 不可寻址时创建临时指针副本
	ptrVal := reflect.New(t)
	ptrVal.Elem().Set(lstValue)
	pt := ptrVal.Type()
	idx = getMethodIndex(pt, tag, methodName)
	return ptrVal.Method(idx)
}

// ListNew 动态创建一个新的List
func ListNew(lstType reflect.Type, cap int) reflect.Value {
	key := lstType.String() + ".New"
	if _, isExists := getCache(key); !isExists {
		method, _ := reflect.New(lstType).Type().MethodByName("New")
		setCache(key, []int{method.Index})
	}

	lstValue := reflect.New(lstType)
	lstValue.Method(getCacheVal(key)[0]).Call([]reflect.Value{reflect.ValueOf(cap)})
	return lstValue
}

// ListAdd 动态添加元素
func ListAdd(lstValue reflect.Value, item any) {
	ListAddValue(lstValue, reflect.ValueOf(item))
}

// ListAddValue 动态添加元素
func ListAddValue(lstValue reflect.Value, itemValue reflect.Value) {
	method := GetAddMethod(lstValue)

	if itemValue.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
	}
	if itemValue.Kind() == reflect.Slice {
		method.CallSlice([]reflect.Value{itemValue})
	} else {
		method.Call([]reflect.Value{itemValue})
	}
}

// GetAddMethod 获取动态添加元素的Method
// 原来用 lstValue.String() 作 key 非常昂贵（format 操作），改为用类型指针
func GetAddMethod(lstValue reflect.Value) reflect.Value {
	return resolveMethod(lstValue, tagAdd, "Add")
}

// GetListItemArrayType 获取List的原始数组类型
func GetListItemArrayType(lstType reflect.Type) reflect.Type {
	key := lstType.String() + ".source"
	if _, isExists := getCache(key); !isExists {
		method, _ := lstType.FieldByName("source")
		setCache(key, method.Index)
	}
	if len(getCacheVal(key)) == 1 {
		return lstType.Field(getCacheVal(key)[0]).Type.Elem()
	}
	return lstType.FieldByIndex(getCacheVal(key)).Type.Elem()
}

// GetListItemType 获取List的元素Type
func GetListItemType(lstType reflect.Type) reflect.Type {
	key := lstType.String() + ".source"

	if _, isExists := getCache(key); !isExists {
		method, _ := lstType.FieldByName("source")
		setCache(key, method.Index)
	}

	var field reflect.Type
	if len(getCacheVal(key)) == 1 {
		field = lstType.Field(getCacheVal(key)[0]).Type
	} else {
		field = lstType.FieldByIndex(getCacheVal(key)).Type
	}

	return field.Elem().Elem()
}

// GetListToArray 在集合中获取数据
func GetListToArray(lstValue reflect.Value) []any {
	arrValue := resolveMethod(lstValue, tagToArrayAny, "ToArrayAny").Call(nil)[0]
	return arrValue.Interface().([]any)
}

// GetListToArrayValue 在集合中获取数据
// 原来用 lstValue.String() 作 key 非常昂贵，改为用 reflect.Type 指针作 key
func GetListToArrayValue(lstValue reflect.Value) reflect.Value {
	return resolveMethod(lstValue, tagToArray, "ToArray").Call(nil)[0]
}

// ExecuteMapperInit 在集合中获取数据
func ExecuteMapperInit(targetVal reflect.Value) {
	resolveMethod(targetVal, tagMapperInit, "MapperInit").Call([]reflect.Value{})
}

func setCache(key string, val []int) {
	lock.Lock()
	Cache[key] = val
	lock.Unlock()
}

func getCache(key string) ([]int, bool) {
	lock.RLock()
	val, isExists := Cache[key]
	lock.RUnlock()
	return val, isExists
}

func getCacheVal(key string) []int {
	lock.RLock()
	val, _ := Cache[key]
	lock.RUnlock()
	return val
}
