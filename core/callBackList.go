package core

type CallbackFn struct {
	F    func()
	Name string
}

var CallbackInitList []CallbackFn // 回调函数初始列表
// AddInitCallback 添加框架启动完后执行的函数
func AddInitCallback(name string, fn func()) {
	CallbackInitList = append(CallbackInitList, CallbackFn{Name: name, F: fn})
}

var CallbackExitList []CallbackFn // 回调函数退出列表
// AddExitCallback 添加应用退出时执行的函数
func AddExitCallback(name string, fn func()) {
	CallbackExitList = append(CallbackExitList, CallbackFn{Name: name, F: fn})
}
