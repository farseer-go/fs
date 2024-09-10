package container

import (
	"fmt"
	"github.com/farseer-go/fs/container/eumLifecycle"
	"github.com/farseer-go/fs/flog"
	"reflect"
	"sync"
	"time"
)

// 容器
type container struct {
	name string
	//dependency map[reflect.Type][]*componentModel // 依赖
	dependency sync.Map // 依赖
	//component  []*componentModel                  // 实现类
	lock *sync.RWMutex
}

// NewContainer 实例化一个默认容器
func NewContainer() *container {
	return &container{
		name: "default",
		//dependency: make(map[reflect.Type][]*componentModel),
		dependency: sync.Map{},
		//component:  []*componentModel{},
		lock: &sync.RWMutex{},
	}
}

// 注册实例，添加到依赖列表
func (r *container) addComponent(model *componentModel) {
	r.lock.Lock()
	defer r.lock.Unlock()

	componentModels, exists := r.dependency.Load(model.interfaceType)
	if !exists {
		r.dependency.Store(model.interfaceType, []*componentModel{model})
	} else {
		comModels := componentModels.([]*componentModel)
		for index := 0; index < len(comModels); index++ {
			if comModels[index].name == model.name {
				panic(fmt.Sprintf("container：The same registration object already exists,interfaceType=%s,name=%s,instanceType=%s", model.interfaceType.String(), model.name, reflect.TypeOf(model.instanceType).String()))
			}
		}
		comModels = append(comModels, model)
		r.dependency.Store(model.interfaceType, comModels)
	}
	//r.component = append(r.component, model)
}

// 注册构造函数
func (r *container) registerConstructor(constructor any, name string, lifecycle eumLifecycle.Enum) {
	constructorType := reflect.TypeOf(constructor)

	if constructorType.Kind() != reflect.Func {
		panic("container：Constructor registration，Can only be func type")
	}

	// 检查出参，只能为1个出参
	if constructorType.NumOut() != 1 {
		panic("container：Constructor registration，There can only be 1 out participation")
	}
	interfaceType := constructorType.Out(0)
	if interfaceType.Kind() != reflect.Interface {
		panic("container：Constructor registration，The reference type can only be Interface")
	}

	// 入参中，不能包含当前接口类型
	for inIndex := 0; inIndex < constructorType.NumIn(); inIndex++ {
		if name == "" && constructorType.In(inIndex).String() == interfaceType.String() {
			panic("container：Constructor registration, when no alias is set, the type of the entry cannot be the same as the type of the returned interface")
		}

		if constructorType.In(inIndex).Kind() != reflect.Interface {
			panic("container：Constructor registration，The input type must be interface")
		}
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
		panic("container：实例注册，The interfaceType type can only be Interface")
	}
	model := NewComponentModelByInstance(name, lifecycle, interfaceTypeOf, ins)
	r.addComponent(model)
}

// 获取对象
func (r *container) resolve(interfaceType reflect.Type, name string) (any, error) {
	if interfaceType.Kind() == reflect.Pointer {
		interfaceType = interfaceType.Elem()
	}

	// 通过Interface查找注册过的container
	if interfaceType.Kind() == reflect.Interface {
		r.lock.RLock()
		componentModels, exists := r.dependency.Load(interfaceType)
		r.lock.RUnlock()
		if !exists {
			return nil, fmt.Errorf("container：%s Unregistered，name=%s", interfaceType.String(), name)
		}

		comModels := componentModels.([]*componentModel)
		for i := 0; i < len(comModels); i++ {
			// 找到了实现类
			if comModels[i].name == name {
				return r.getOrCreateIns(interfaceType, i), nil
			}
		}
		return nil, fmt.Errorf("container：%s Unregistered，name=%s", interfaceType.String(), name)

		// 结构对象，直接动态创建
	} else if interfaceType.Kind() == reflect.Struct {
		return r.injectByType(interfaceType), nil
	}
	return nil, fmt.Errorf("container：%s Types not supported，name=%s", interfaceType.String(), name)
}

// 获取所有对象
func (r *container) resolveAll(interfaceType reflect.Type) []any {
	if interfaceType.Kind() != reflect.Interface {
		_ = flog.Errorf("container：When resolve all objects，%s must is Interface type", interfaceType.String())
		return nil
	}

	// 通过Interface查找注册过的container
	r.lock.RLock()
	componentModels, exists := r.dependency.Load(interfaceType)
	r.lock.RUnlock()
	if !exists {
		return nil
	}

	var ins []any
	for i := 0; i < len(componentModels.([]*componentModel)); i++ {
		// 找到了实现类
		ins = append(ins, r.getOrCreateIns(interfaceType, i))
	}
	return ins
}

// 根据lifecycle获取实例
func (r *container) getOrCreateIns(interfaceType reflect.Type, index int) any {
	componentModels, _ := r.dependency.Load(interfaceType)
	comModels := componentModels.([]*componentModel)
	// 更新实例访问时间
	comModels[index].lastVisitAt = time.Now()
	// 单例
	if comModels[index].lifecycle == eumLifecycle.Single {
		if comModels[index].instance == nil {
			comModels[index].instance = r.createIns(comModels[index])
		}
		return comModels[index].instance
	} else {
		return r.createIns(comModels[index])
	}
}

// 根据类型，动态创建实例
func (r *container) createIns(model *componentModel) any {
	var arr []reflect.Value
	// 构造函数，需要分别取出入参值
	for inIndex := 0; inIndex < model.instanceType.NumIn(); inIndex++ {
		val := r.resolveDefaultOrFirstComponent(model.instanceType.In(inIndex))
		arr = append(arr, reflect.ValueOf(val))
	}
	if arr == nil {
		arr = []reflect.Value{}
	}
	return r.inject(model.instanceValue.Call(arr)[0].Interface())
}

// 获取对象，如果默认别名不存在，则使用第一个注册的实例
func (r *container) resolveDefaultOrFirstComponent(interfaceType reflect.Type) any {
	r.lock.Lock()
	componentModels, exists := r.dependency.Load(interfaceType)
	r.lock.Unlock()
	if !exists {
		_ = flog.Errorf("container：%s Unregistered", interfaceType.String())
		return nil
	}

	findIndex := 0
	// 优先找默认实例
	comModels := componentModels.([]*componentModel)
	for i := 0; i < len(comModels); i++ {
		// 找到了实现类
		if comModels[i].name == "" {
			findIndex = i
		}
	}
	return r.getOrCreateIns(interfaceType, findIndex)
}

// 解析注入
func (r *container) inject(ins any) any {
	if ins == nil {
		return ins
	}
	insVal := reflect.Indirect(reflect.ValueOf(ins))
	for i := 0; i < insVal.NumField(); i++ {
		field := insVal.Type().Field(i)
		if field.IsExported() && field.Type.Kind() == reflect.Interface && insVal.Field(i).IsNil() {
			fieldIns, err := r.resolve(field.Type, field.Tag.Get("inject"))
			if err != nil {
				_ = flog.Error(err)
				continue
			}
			insVal.Field(i).Set(reflect.ValueOf(fieldIns))
		}
	}
	return ins
}

// 解析注入
func (r *container) injectByType(instanceType reflect.Type) any {
	instanceVal := reflect.New(instanceType).Elem()
	for i := 0; i < instanceVal.NumField(); i++ {
		field := instanceVal.Type().Field(i)
		if field.IsExported() && field.Type.Kind() == reflect.Interface {
			fieldIns, err := r.resolve(field.Type, field.Tag.Get("inject"))
			if err != nil {
				_ = flog.Error(err)
				continue
			}
			instanceVal.Field(i).Set(reflect.ValueOf(fieldIns))
		}
	}
	return instanceVal.Interface()
}

// 是否注册过
func (r *container) isRegister(interfaceType reflect.Type, name string) bool {
	r.lock.RLock()
	componentModels, exists := r.dependency.Load(interfaceType)
	r.lock.RUnlock()
	if !exists {
		return false
	}
	comModels := componentModels.([]*componentModel)
	for i := 0; i < len(comModels); i++ {
		// 找到了实现类
		if comModels[i].name == name {
			return true
		}
	}
	return false
}

// 移除已注册的实例
func (r *container) removeComponent(interfaceType reflect.Type, name string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	componentModels, _ := r.dependency.Load(interfaceType)
	comModels := componentModels.([]*componentModel)
	// 遍历已注册的实例列表
	for index := 0; index < len(comModels); index++ {
		// 找到实例后，删除
		if comModels[index].name == name {
			comModels = append(comModels[:index], comModels[index+1:]...)
		}
	}
	r.dependency.Store(interfaceType, comModels)
}

// 移除长时间未使用的实例
func (r *container) removeUnused(interfaceType reflect.Type, ttl time.Duration) {
	r.lock.Lock()
	defer r.lock.Unlock()

	componentModels, _ := r.dependency.Load(interfaceType)
	comModels := componentModels.([]*componentModel)
	// 遍历已注册的实例列表
	for index := 0; index < len(comModels); index++ {
		// 删除超出ttl时间未访问的实例
		if time.Now().Sub(comModels[index].lastVisitAt) >= ttl {
			comModels = append(comModels[:index], comModels[index+1:]...)
		}
	}
	r.dependency.Store(interfaceType, comModels)
}
