package container

import (
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
				flog.Errorf("container：已存在同样的注册对象,interfaceType=%s,name=%s,getInstanceType=%s", model.interfaceType.String(), model.name, reflect.TypeOf(model.getInstanceType).String())
				os.Exit(-1)
			}
		}
		r.dependency[model.interfaceType] = append(componentModels, model)
	}
	r.component = append(r.component, model)
}

// 注册构造函数
func (r *container) registerConstructor(constructor any, name string, lifecycle eumLifecycle.Enum) {
	constructorType := reflect.TypeOf(constructor)
	if constructorType.NumIn() != 0 {
		flog.Error("container：构造函数注册，不能有入参")
		os.Exit(-1)
	}
	if constructorType.NumOut() != 1 {
		flog.Error("container：构造函数注册，只能有1个出参")
		os.Exit(-1)
	}
	interfaceType := constructorType.Out(0)
	if interfaceType.Kind() != reflect.Interface {
		flog.Error("container：构造函数注册，出参类型只能为Interface")
		os.Exit(-1)
	}
	model := NewComponentModel(name, lifecycle, interfaceType, constructor)
	r.addComponent(model)
}

// 注册实例
func (r *container) registerInstance(interfaceType any, ins struct{}, name string, lifecycle eumLifecycle.Enum) {

}

// 注册方法
func (r *container) registerMethod(interfaceType any, method any, name string, lifecycle eumLifecycle.Enum) {

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
			// 单例
			if componentModels[i].lifecycle == eumLifecycle.Single {
				if componentModels[i].instance == nil {
					componentModels[i].instance = r.createIns(componentModels[i])
				}
				return componentModels[i].instance
			} else {
				return r.createIns(componentModels[i])
			}
		}
	}
	flog.Errorf("container：%s未注册，name=%s", interfaceType.String(), name)
	return nil
}

// 根据类型，动态创建实例
func (r *container) createIns(model componentModel) any {
	getInstanceType := reflect.TypeOf(model.getInstanceType)

	// 构造函数注册
	if getInstanceType.NumIn() == 0 && getInstanceType.NumOut() == 1 {
		return reflect.ValueOf(model.getInstanceType).Call([]reflect.Value{})[0].Interface()
	}

	return nil
}
