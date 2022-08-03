package modules

type FarseerModule interface {
	// DependsModule 依赖的模块
	DependsModule() []FarseerModule
	// PreInitialize 预初始化
	PreInitialize()
	// Initialize 初始化
	Initialize()
	// PostInitialize 初始化之后
	PostInitialize()
	// Shutdown 应用关闭之前先关闭模块
	Shutdown()
}
