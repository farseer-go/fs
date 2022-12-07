# 概述
- [文档（国内）](https://farseer-go.gitee.io/)
- [文档（国外）](https://farseer-go.github.io/doc/)
- [开源（Github）](https://github.com/farseer-go/fs)

![](https://img.shields.io/github/stars/farseer-go?style=social)
![](https://img.shields.io/github/license/farseer-go/fs)
![](https://img.shields.io/github/go-mod/go-version/farseer-go/fs)
![](https://img.shields.io/github/v/release/farseer-go/fs)
![](https://img.shields.io/github/languages/code-size/farseer-go/fs)
![](https://img.shields.io/github/directory-file-count/farseer-go/fs)
![](https://img.shields.io/github/last-commit/farseer-go/fs)

## 什么是farseer-go
针对 `golang` 平台下的一套技术框架。

我们为您选型出目前最为流行的组件，并按模块化来提供使用这些组件。

框架完美支持 `DDD领域驱动` 的战术设计，如`仓储资源库`、`应用层事务`、`领域事件`、`应用层动态WebAPI`。

只需要这一套框架便可应付常用的项目应用

它有一个[.net core](https://github.com/FarseerNet/Farseer.Net/) 成熟版本，已经使用了10多年，非常棒

?> 不用担心框架会让你依赖过多的包，我们的组件都是独立的包，意味着如果你使用`webapi组件`则不会依赖`redis包`

!> 每个组件都是单独的包，因此版本号也是单独发布的

## 有什么特点？

- `统一配置`：`farseer-go`将常用包的各种连接字符串都整合到`./farseer.yaml`

- `优雅`：所有的模块都遵循开发者体验第一为原则。

- `极简`：尽可能让您少依赖模块中的包、参数。

- `模块化`：未使用的模块不会下载到你的环境中。

- `链路追踪`（下个版本推出）：如果您使用框架中的Orm、Redis、Http、Grpc、ES、MQ、EventBus、Task、FSS，将隐式为您实现链路追踪，并提供API请求日志、慢查询。

> 结合[FOPS](https://github.com/FarseerNet/FOPS) 项目（自动构建、链路追踪控制台、K8S集群日志收集）支持代码无侵入的全链路实时监控。

## 集成的组件

| 包名            | 描述          |                                                                                                                                                                                                                                                                                                             |
|---------------|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| collections   | 数据集合        | ![](https://img.shields.io/github/v/release/farseer-go/collections)![](https://img.shields.io/github/languages/code-size/farseer-go/collections)![](https://img.shields.io/github/directory-file-count/farseer-go/collections)![](https://img.shields.io/github/last-commit/farseer-go/collections)         |
| webapi        | web api服务   | ![](https://img.shields.io/github/v/release/farseer-go/webapi)![](https://img.shields.io/github/languages/code-size/farseer-go/webapi)![](https://img.shields.io/github/directory-file-count/farseer-go/webapi)![](https://img.shields.io/github/last-commit/farseer-go/webapi)                             |
| async         | 异步编程        | ![](https://img.shields.io/github/v/release/farseer-go/async)![](https://img.shields.io/github/languages/code-size/farseer-go/async)![](https://img.shields.io/github/directory-file-count/farseer-go/async)![](https://img.shields.io/github/last-commit/farseer-go/async)                                 |
| mapper        | 对象转换        | ![](https://img.shields.io/github/v/release/farseer-go/mapper)![](https://img.shields.io/github/languages/code-size/farseer-go/mapper)![](https://img.shields.io/github/directory-file-count/farseer-go/mapper)![](https://img.shields.io/github/last-commit/farseer-go/mapper)                             |
| cacheMemory   | 本地缓存        | ![](https://img.shields.io/github/v/release/farseer-go/cacheMemory)![](https://img.shields.io/github/languages/code-size/farseer-go/cacheMemory)![](https://img.shields.io/github/directory-file-count/farseer-go/cacheMemory)![](https://img.shields.io/github/last-commit/farseer-go/cacheMemory)         |
| redis         | client      | ![](https://img.shields.io/github/v/release/farseer-go/redis)![](https://img.shields.io/github/languages/code-size/farseer-go/redis)![](https://img.shields.io/github/directory-file-count/farseer-go/redis)![](https://img.shields.io/github/last-commit/farseer-go/redis)                                 |
| data          | 数据库ORM      | ![](https://img.shields.io/github/v/release/farseer-go/data)![](https://img.shields.io/github/languages/code-size/farseer-go/data)![](https://img.shields.io/github/directory-file-count/farseer-go/data)![](https://img.shields.io/github/last-commit/farseer-go/data)                                     |
| elasticSearch | client      | ![](https://img.shields.io/github/v/release/farseer-go/elasticSearch)![](https://img.shields.io/github/languages/code-size/farseer-go/elasticSearch)![](https://img.shields.io/github/directory-file-count/farseer-go/elasticSearch)![](https://img.shields.io/github/last-commit/farseer-go/elasticSearch) |
| eventBus      | 事件总线        | ![](https://img.shields.io/github/v/release/farseer-go/eventBus)![](https://img.shields.io/github/languages/code-size/farseer-go/eventBus)![](https://img.shields.io/github/directory-file-count/farseer-go/eventBus)![](https://img.shields.io/github/last-commit/farseer-go/eventBus)                     |
| queue         | 本地队列        | ![](https://img.shields.io/github/v/release/farseer-go/queue)![](https://img.shields.io/github/languages/code-size/farseer-go/queue)![](https://img.shields.io/github/directory-file-count/farseer-go/queue)![](https://img.shields.io/github/last-commit/farseer-go/queue)                                 |
| tasks         | 本地任务        | ![](https://img.shields.io/github/v/release/farseer-go/tasks)![](https://img.shields.io/github/languages/code-size/farseer-go/tasks)![](https://img.shields.io/github/directory-file-count/farseer-go/tasks)![](https://img.shields.io/github/last-commit/farseer-go/tasks)                                 |
| fss           | 分布试调度client | ![](https://img.shields.io/github/v/release/farseer-go/fss)![](https://img.shields.io/github/languages/code-size/farseer-go/fss)![](https://img.shields.io/github/directory-file-count/farseer-go/fss)![](https://img.shields.io/github/last-commit/farseer-go/fss)                                         |
| utils         | 工具集         | ![](https://img.shields.io/github/v/release/farseer-go/utils)![](https://img.shields.io/github/languages/code-size/farseer-go/utils)![](https://img.shields.io/github/directory-file-count/farseer-go/utils)![](https://img.shields.io/github/last-commit/farseer-go/utils)                                 |
| linkTrack     | 链路追踪        | （即将推出）                                                                                                                                                                                                                                                                                                      |
| rabbit        | client      | （即将推出）                                                                                                                                                                                                                                                                                                      |
| redisStream   | redis mq    | （即将推出）                                                                                                                                                                                                                                                                                                      |

## 如何开始

_main.go_
```go
package main
import "github.com/farseer-go/fs"

func main() {
	fs.Initialize[StartupModule]("your project Name")
}
```

?> 在main函数第一行，执行`fs.Initialize`，开始初始化框架

运行后控制台打印加载信息：

```
2022-12-01 17:07:24 应用名称： your project Name
2022-12-01 17:07:24 主机名称： MacBook-Pro.local
2022-12-01 17:07:24 系统时间： 2022-12-01 17:07:24
2022-12-01 17:07:24   进程ID： 6123
2022-12-01 17:07:24   应用ID： 193337022963818496
2022-12-01 17:07:24   应用IP： 192.168.1.4
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
2022-12-01 17:07:24 初始化完毕，共耗时：1 ms 
2022-12-01 17:07:24 ---------------------------------------
2022-12-01 17:07:24 [Info] Web服务已启动：http://localhost:8888/
```
## Stargazers

[![Stargazers repo roster for @farseer-go/fs](https://reporoster.com/stars/farseer-go/fs)](https://github.com/farseer-go/fs/stargazers)

## Forkers

[![Forkers repo roster for @farseer-go/fs](https://reporoster.com/forks/farseer-go/fs)](https://github.com/farseer-go/fs/network/members)