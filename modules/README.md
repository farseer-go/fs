# Getting Started with modules

## What's the use？
Use instead of `func init() { }` to accomplish what you expect by an explicit initialization sequence

## What modules are available for farseer-go？
* cache.Module
* cacheMemory.Module
* data.Module
* elasticSearch.Module
* eventBus.Module
* fSchedule.Module
* queue.Module
* rabbit.Module
* redis.Module
* tasks.Module

## Add your StartupModule Files
```go
type StartupModule struct {
}

// Dependent modules
func (module StartupModule) DependsModule() []modules.FarseerModule {
    return []modules.FarseerModule{interfaces.Module{}, infrastructure.Module{}}
}

// Pre-initialization
func (module StartupModule) PreInitialize() {
}

// Initialize
func (module StartupModule) Initialize() {
}

// PostInitialize
func (module StartupModule) PostInitialize() {
}

// Shutdown
func (module StartupModule) Shutdown() {
}
```

## DependsModule
DependsModule is used to load modules that you need to depend on, such as data.Module, redis.Module, cache.Module, or your business module

## Initialization
```go
// First all dependent modules will be loaded according to DependsModule's dependencies.
// and perform the initialization of each module in the order of dependency
fs.Initialize[StartupModule]("FOPS")
```