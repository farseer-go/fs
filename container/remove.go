package container

import (
	"reflect"
	"time"
)

// Remove 移除已注册的实例
func Remove[TInterface any](iocName ...string) {
	name := getIocName(iocName...)
	interfaceType := reflect.TypeOf((*TInterface)(nil)).Elem()
	defContainer.removeComponent(interfaceType, name)
}

// RemoveUnused 移除长时间未使用的实例
func RemoveUnused[TInterface any](ttl time.Duration) {
	interfaceType := reflect.TypeOf((*TInterface)(nil)).Elem()
	defContainer.removeUnused(interfaceType, ttl)
}
