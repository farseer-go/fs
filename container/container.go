package container

import (
	"fmt"
	"github.com/farseer-go/fs/container/eumLifecycle"
	"github.com/farseer-go/fs/flog"
	"os"
	"reflect"
)

// 容器
type container struct {
	name       string
	dependency map[reflect.Type][]componentModel // 依赖
	component  []componentModel                  // 实现类
}

// NewContainer 实例化一个默认容器
func NewContainer() *container {
	return &container{
		name:       "default",
		dependency: make(map[reflect.Type][]componentModel),
		component:  []componentModel{},
	}
}

// 注册实例
func (r *container) addComponent(model componentModel) {
	componentModels, exists := r.dependency[model.interfaceType]
	if !exists {
		r.dependency[model.interfaceType] = []componentModel{model}
	} else {
		for index := 0; index < len(componentModels); index++ {
			if componentModels[index].name == model.name && componentModels[index].getInstanceType == model.getInstanceType {
				panic(fmt.Sprintf("container：已存在同样的注册对象,interfaceType=%s,name=%s,getInstanceType=%s", model.interfaceType.String(), model.name, reflect.TypeOf(model.getInstanceType).String()))
			}
		}
		r.dependency[model.interfaceType] = append(componentModels, model)
	}
	r.component = append(r.component, model)
}

// 注册构造函数
func (r *container) registerConstructor(constructor any, name string, lifecycle eumLifecycle.Enum) {
	constructorType := reflect.TypeOf(constructor)
	for inIndex := 0; inIndex < constructorType.NumIn(); inIndex++ {
		if name == "" && constructorType.In(inIndex).String() == constructorType.String() {
			panic("container：构造函数注册，当未设置别名时，入参的类型不能与返回的接口类型一样")
		}

		if constructorType.In(inIndex).Kind() != reflect.Interface {
			panic("container：构造函数注册，入参类型必须为interface")
		}
	}
	if constructorType.NumOut() != 1 {
		panic("container：构造函数注册，只能有1个出参")
	}
	interfaceType := constructorType.Out(0)
	if interfaceType.Kind() != reflect.Interface {
		panic("container：构造函数注册，出参类型只能为Interface")
	}
	model := NewComponentModel(name, lifecycle, interfaceType, constructor)
	r.addComponent(model)
}

// 注册实例
func (r *container) registerInstance(interfaceType any, ins any, name string, lifecycle eumLifecycle.Enum) {
	interfaceTypeOf := reflect.TypeOf(interfaceType)
	if interfaceTypeOf.Kind() == reflect.Pointer {
		interfaceTypeOf = interfaceTypeOf.Elem()
	}
	if interfaceTypeOf.Kind() != reflect.Interface {
		flog.Error("container：实例注册，interfaceType类型只能为Interface")
		os.Exit(-1)
	}
	model := NewComponentModel(name, lifecycle, interfaceTypeOf, ins)
	model.instance = ins
	r.addComponent(model)
}

// 获取对象
func (r *container) resolve(interfaceType reflect.Type, name string) any {
	componentModels, exists := r.dependency[interfaceType]
	if !exists {
		flog.Errorf("container：%s未注册", interfaceType.String())
		return nil
	}

	for i := 0; i < len(componentModels); i++ {
		// 找到了实现类
		if componentModels[i].name == name {
			return r.getOrCreateIns(interfaceType, i)
		}
	}
	flog.Errorf("container：%s未注册，name=%s", interfaceType.String(), name)
	return nil
}

// 根据lifecycle获取实例
func (r *container) getOrCreateIns(interfaceType reflect.Type, index int) any {
	// 单例
	if r.dependency[interfaceType][index].lifecycle == eumLifecycle.Single {
		if r.dependency[interfaceType][index].instance == nil {
			r.dependency[interfaceType][index].instance = r.createIns(r.dependency[interfaceType][index])
		}
		return r.dependency[interfaceType][index].instance
	} else {
		return r.createIns(r.dependency[interfaceType][index])
	}
}

// 根据类型，动态创建实例
func (r *container) createIns(model componentModel) any {
	getInstanceType := reflect.TypeOf(model.getInstanceType)

	if getInstanceType.Kind() == reflect.Func {
		var arr []reflect.Value
		// 构造函数，需要分别取出入参值
		for inIndex := 0; inIndex < getInstanceType.NumIn(); inIndex++ {
			val := reflect.ValueOf(r.resolveDefaultOrFirstComponent(getInstanceType.In(inIndex)))
			arr = append(arr, val)
		}
		if arr == nil {
			arr = []reflect.Value{}
		}
		return reflect.ValueOf(model.getInstanceType).Call(arr)[0].Interface()
	}
	if getInstanceType.Kind() == reflect.Struct {
		return model.getInstanceType
	}
	return nil
}

// 获取对象，如果默认别名不存在，则使用第一个注册的实例
func (r *container) resolveDefaultOrFirstComponent(interfaceType reflect.Type) any {
	componentModels, exists := r.dependency[interfaceType]
	if !exists {
		flog.Errorf("container：%s未注册", interfaceType.String())
		return nil
	}

	findIndex := 0
	// 优先找默认实例
	for i := 0; i < len(componentModels); i++ {
		// 找到了实现类
		if componentModels[i].name == "" {
			findIndex = i
		}
	}
	return r.getOrCreateIns(interfaceType, findIndex)
}
