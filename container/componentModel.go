package container

import (
	"github.com/farseer-go/fs/container/eumLifecycle"
	"reflect"
)

// 实现类模型
type componentModel struct {
	name            string            // 别名
	lifecycle       eumLifecycle.Enum // 生命周期
	interfaceType   reflect.Type      // 继承的接口
	getInstanceType any               // 函数的接口
	instance        any               // 实例
}

func NewComponentModel(name string, lifecycle eumLifecycle.Enum, interfaceType reflect.Type, getInstanceType any) componentModel {
	return componentModel{
		name:            name,
		lifecycle:       lifecycle,
		interfaceType:   interfaceType,
		getInstanceType: getInstanceType,
	}
}
