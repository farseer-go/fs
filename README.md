# 概述
- [文档](https://farseer-go.github.io/doc/)
- [GitHub](https://github.com/farseer-go)

## 什么是farseer-go
针对 `golang` 平台下的一套技术框架。

我们为您选型出目前最为流行的组件，并按模块化来提供使用这些组件。

框架完美支持 `DDD领域驱动` 的战术设计，如`仓储资源库`、`应用层事务`、`领域事件`、`应用层动态WebAPI`。

只需要这一套框架便可应付常用的项目应用

它有一个[.net core](https://github.com/FarseerNet/Farseer.Net/) 成熟版本，已经使用了10多年，非常棒

?> 不用担心框架会让你依赖过多的包，我们的组件都是独立的包，意味着如果你使用`webapi组件`则不会依赖`redis包`

!> 每个组件都是单独的包，因此版本号也是单独发布的

## 有什么特点？

### 1、统一配置

常用的框架中都需要配置各种连接字符串，而`farseer-go`将这些配置都整合到`farseer.yaml`

### 2、优雅

我们会让写代码跟阅读英文书一样流畅，没有晦涩难懂、无意义的参数

### 3、极简

所有的组件功能在提供使用时，最简单的调用，是我们最主要的考虑目标之一

### 4、轻量

我们将所有组件打散在单独地包中，或许你只需要用到像webapi组件时，不会让你依赖其它不需要的包

### 5、链路追踪

如果您使用我们提供的Orm、Redis、Http、Grpc、Elasticsearch、MQ(Rabbit、RedisStream、Rocker、本地Queue)、EventBus、Task、FSS等等，您什么都不需要做，系统将隐式为您实现链路追踪，并提供API请求日志、慢查询（前面提到的都会记录）。

> 结合[FOPS](https://github.com/FarseerNet/FOPS) 项目（自动构建、链路追踪控制台、K8S集群日志收集）支持代码无侵入的全链路实时监控。

## 集成的组件
- `collections`：数据集合
- `webapi`：web api服务
- `async`：异步编程
- `mapper`：对象转换
- `cacheMemory`：本地缓存
- `redis`：redis client
- `data`：数据库ORM
- `elasticSearch`：es client
- `eventBus`：事件总线
- `queue`：本地队列
- `tasks`：本地任务
- `fss`：分布试调度中心client
- `utils`：工具集
- linkTrack：链路追踪（下一版本推出）
- rabbit：rabbit client（下一版本推出）
- redisStream：redis mq（下一版本推出）

## 如何开始

_main.go_
```go
package main
import "github.com/farseer-go/fs"

func main() {
	fs.Initialize[StartupModule]("your project Name")
}
```

?> 只需要在main函数第一行，执行`fs.Initialize`，即可初始化框架

运行后，会在控制台打印加载信息：

```
2022-12-01 17:07:24 应用名称： your project Name
2022-12-01 17:07:24 主机名称： MacBook-Pro.local
2022-12-01 17:07:24 系统时间： 2022-12-01 17:07:24
2022-12-01 17:07:24   进程ID： 6123
2022-12-01 17:07:24   应用ID： 193337022963818496
2022-12-01 17:07:24   应用IP： 192.168.1.4
2022-12-01 17:07:24 日志开关： 
2022-12-01 17:07:24 ---------------------------------------
2022-12-01 17:07:24 加载模块...
2022-12-01 17:07:24 加载模块：webapi.Module
2022-12-01 17:07:24 加载模块：domain.Module
2022-12-01 17:07:24 加载模块：application.Module
2022-12-01 17:07:24 加载模块：interfaces.Module
2022-12-01 17:07:24 加载模块：data.Module
2022-12-01 17:07:24 加载模块：eventBus.Module
2022-12-01 17:07:24 加载模块：queue.Module
2022-12-01 17:07:24 加载模块：infrastructure.Module
2022-12-01 17:07:24 加载模块：main.StartupModule
2022-12-01 17:07:24 加载完毕，共加载 10 个模块
2022-12-01 17:07:24 ---------------------------------------
2022-12-01 17:07:24 Modules模块初始化...
2022-12-01 17:07:24 耗时：0 ms modules.FarseerKernelModule.PreInitialize()
2022-12-01 17:07:24 耗时：0 ms webapi.Module.PreInitialize()
2022-12-01 17:07:24 耗时：0 ms domain.Module.PreInitialize()
2022-12-01 17:07:24 耗时：0 ms application.Module.PreInitialize()
2022-12-01 17:07:24 耗时：0 ms interfaces.Module.PreInitialize()
2022-12-01 17:07:24 耗时：0 ms data.Module.PreInitialize()
2022-12-01 17:07:24 耗时：0 ms eventBus.Module.PreInitialize()
2022-12-01 17:07:24 耗时：0 ms queue.Module.PreInitialize()
2022-12-01 17:07:24 耗时：0 ms infrastructure.Module.PreInitialize()
2022-12-01 17:07:24 耗时：0 ms main.StartupModule.PreInitialize()
2022-12-01 17:07:24 ---------------------------------------
2022-12-01 17:07:24 耗时：0 ms modules.FarseerKernelModule.Initialize()
2022-12-01 17:07:24 耗时：0 ms webapi.Module.Initialize()
2022-12-01 17:07:24 耗时：0 ms domain.Module.Initialize()
2022-12-01 17:07:24 耗时：0 ms application.Module.Initialize()
2022-12-01 17:07:24 耗时：0 ms interfaces.Module.Initialize()
2022-12-01 17:07:24 耗时：0 ms data.Module.Initialize()
2022-12-01 17:07:24 耗时：0 ms eventBus.Module.Initialize()
2022-12-01 17:07:24 耗时：0 ms queue.Module.Initialize()
2022-12-01 17:07:24 耗时：0 ms infrastructure.Module.Initialize()
2022-12-01 17:07:24 耗时：0 ms main.StartupModule.Initialize()
2022-12-01 17:07:24 ---------------------------------------
2022-12-01 17:07:24 耗时：0 ms modules.FarseerKernelModule.PostInitialize()
2022-12-01 17:07:24 耗时：0 ms webapi.Module.PostInitialize()
2022-12-01 17:07:24 耗时：0 ms domain.Module.PostInitialize()
2022-12-01 17:07:24 耗时：0 ms application.Module.PostInitialize()
2022-12-01 17:07:24 耗时：0 ms interfaces.Module.PostInitialize()
2022-12-01 17:07:24 耗时：0 ms data.Module.PostInitialize()
2022-12-01 17:07:24 耗时：0 ms eventBus.Module.PostInitialize()
2022-12-01 17:07:24 耗时：0 ms queue.Module.PostInitialize()
2022-12-01 17:07:24 耗时：0 ms infrastructure.Module.PostInitialize()
2022-12-01 17:07:24 耗时：0 ms main.StartupModule.PostInitialize()
2022-12-01 17:07:24 基础组件初始化完成
2022-12-01 17:07:24 初始化完毕，共耗时：1 ms 
2022-12-01 17:07:24 ---------------------------------------
2022-12-01 17:07:24 [Info] Web服务已启动：http://localhost:8888/

```