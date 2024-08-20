package types

import (
	"reflect"
	"sync"
)

// string key=自定义标识
// value int=field or method对应的索引
var Cache = make(map[string][]int)
var lock sync.RWMutex

// ListNew 动态创建一个新的List
func ListNew(lstType reflect.Type) reflect.Value {
	key := lstType.String() + ".New"
	if _, isExists := getCache(key); !isExists {
		method, _ := reflect.New(lstType).Type().MethodByName("New")
		setCache(key, []int{method.Index})
	}

	lstValue := reflect.New(lstType)
	lstValue.Method(getCacheVal(key)[0]).Call(nil)
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
func GetAddMethod(lstValue reflect.Value) reflect.Value {
	// 初始化反射Method
	key := lstValue.String() + ".Add"
	if _, isExists := getCache(key); !isExists {
		method, _ := lstValue.Type().MethodByName("Add")
		setCache(key, []int{method.Index})
	}
	return lstValue.Method(getCacheVal(key)[0])
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
	key := lstValue.String() + ".ToArrayAny"
	if _, isExists := getCache(key); !isExists {
		method, _ := lstValue.Type().MethodByName("ToArrayAny")
		setCache(key, []int{method.Index})
	}
	arrValue := lstValue.Method(getCacheVal(key)[0]).Call(nil)[0]
	return arrValue.Interface().([]any)
}

// GetListToArrayValue 在集合中获取数据
func GetListToArrayValue(lstValue reflect.Value) reflect.Value {
	key := lstValue.String() + ".ToArray"
	if _, isExists := getCache(key); !isExists {
		method, _ := lstValue.Type().MethodByName("ToArray")
		setCache(key, []int{method.Index})
	}

	return lstValue.Method(getCacheVal(key)[0]).Call(nil)[0]
}

// ExecuteMapperInit 在集合中获取数据
func ExecuteMapperInit(targetVal reflect.Value) {
	key := targetVal.String() + ".MapperInit"
	if _, isExists := getCache(key); !isExists {
		method, _ := targetVal.Type().MethodByName("MapperInit")
		setCache(key, []int{method.Index})
	}

	targetVal.Method(getCacheVal(key)[0]).Call([]reflect.Value{})
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
