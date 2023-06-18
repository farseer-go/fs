package modules

// FarseerModule 依赖的模块
type FarseerModule interface {
	// DependsModule 依赖的模块
	DependsModule() []FarseerModule
}

// FarseerPreInitializeModule 预初始化（常用于全局变量初始化）
type FarseerPreInitializeModule interface {
	FarseerModule
	// PreInitialize 预初始化
	PreInitialize()
}

// FarseerInitializeModule 初始化（常用于根据配置设置初始化对象）
type FarseerInitializeModule interface {
	FarseerModule
	// Initialize 初始化
	Initialize()
}

// FarseerPostInitializeModule 初始化之后（常用于启动协程服务）
type FarseerPostInitializeModule interface {
	FarseerModule
	// PostInitialize 初始化之后
	PostInitialize()
}

// FarseerShutdownModule 应用关闭之前先关闭模块
type FarseerShutdownModule interface {
	FarseerModule
	// Shutdown 应用关闭之前先关闭模块
	Shutdown()
}
