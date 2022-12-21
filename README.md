# Overview
- `Document`
    - [English](https://farseer-go.gitee.io/en-us/)
    - [中文](https://farseer-go.gitee.io/)
    - [English](https://farseer-go.github.io/doc/en-us/)
- Source
    - [github](https://github.com/farseer-go/fs)

![](https://img.shields.io/github/stars/farseer-go?style=social)
![](https://img.shields.io/github/license/farseer-go/fs)
![](https://img.shields.io/github/go-mod/go-version/farseer-go/fs)
![](https://img.shields.io/github/v/release/farseer-go/fs)
![](https://img.shields.io/github/languages/code-size/farseer-go/fs)
![](https://img.shields.io/github/directory-file-count/farseer-go/fs)
![](https://goreportcard.com/badge/github.com/farseer-go/fs)

## Introduction

A **modular** and **complete infrastructure** framework based on Golang，create modern web applications and APIs

Use the most popular components available today and provide them with modular technology.

The framework perfectly supports `DDD domain-driven` tactical design, such as `warehousing repository`
, `application-layer transactions`, `domain events`, `application-layer dynamic WebAPI`.

It has a mature version of [.net core](https://github.com/FarseerNet/Farseer.Net/) that has been in use for over 10
years and is great

![](https://farseer-go.gitee.io/images/farseer-go.png)

> Don't worry about the framework making you depend on too many packages, farseer-go's components are all separate
> packages and unused packages are not downloaded into your application

> Each component is a separate package, so the version number is also released separately

## Features

- `Unified configuration`: all configurations are consolidated into `. /farseer.yaml`

- `Elegant`: all modules follow the principle of developer experience first.

- `Modularity`: provides a complete modular system that allows you to develop reusable application modules.

- `Domain-driven`: helps you implement a hierarchical architecture based on DDD and build a maintainable code base.

- `link tracking` (coming in the next version): if you use the framework Orm, Redis, Http, Grpc, ES, MQ, EventBus, Task, FSS, will implicitly implement link tracking for you and provide API request logs, slow queries.

> Combined with [FOPS](https://github.com/FarseerNet/FOPS) project (automatic build, link tracing console, K8S cluster log collection) support code non-intrusive full link real-time monitoring.

## Components

| Package Name  | Description        |                                                                                                                                                                                                                                                                                                             |
|---------------|--------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| fs            | Basic Core Package | ![](https://img.shields.io/github/v/release/farseer-go/fs)![](https://img.shields.io/github/languages/code-size/farseer-go/fs)![](https://img.shields.io/github/directory-file-count/farseer-go/fs)![](https://goreportcard.com/badge/github.com/farseer-go/fs)                                             |
| collections   | Data Collection    | ![](https://img.shields.io/github/v/release/farseer-go/collections)![](https://img.shields.io/github/languages/code-size/farseer-go/collections)![](https://img.shields.io/github/directory-file-count/farseer-go/collections)![](https://goreportcard.com/badge/github.com/farseer-go/collections)         |
| webapi        | web api            | ![](https://img.shields.io/github/v/release/farseer-go/webapi)![](https://img.shields.io/github/languages/code-size/farseer-go/webapi)![](https://img.shields.io/github/directory-file-count/farseer-go/webapi)![](https://goreportcard.com/badge/github.com/farseer-go/webapi)                             |
| async         | Async Programming  | ![](https://img.shields.io/github/v/release/farseer-go/async)![](https://img.shields.io/github/languages/code-size/farseer-go/async)![](https://img.shields.io/github/directory-file-count/farseer-go/async)![](https://goreportcard.com/badge/github.com/farseer-go/async)                                 |
| mapper        | Object Conversion  | ![](https://img.shields.io/github/v/release/farseer-go/mapper)![](https://img.shields.io/github/languages/code-size/farseer-go/mapper)![](https://img.shields.io/github/directory-file-count/farseer-go/mapper)![](https://goreportcard.com/badge/github.com/farseer-go/mapper)                             |
| cacheMemory   | Local Cache        | ![](https://img.shields.io/github/v/release/farseer-go/cacheMemory)![](https://img.shields.io/github/languages/code-size/farseer-go/cacheMemory)![](https://img.shields.io/github/directory-file-count/farseer-go/cacheMemory)![](https://goreportcard.com/badge/github.com/farseer-go/cacheMemory)         |
| redis         | client             | ![](https://img.shields.io/github/v/release/farseer-go/redis)![](https://img.shields.io/github/languages/code-size/farseer-go/redis)![](https://img.shields.io/github/directory-file-count/farseer-go/redis)![](https://goreportcard.com/badge/github.com/farseer-go/redis)                                 |
| data          | DataBase ORM       | ![](https://img.shields.io/github/v/release/farseer-go/data)![](https://img.shields.io/github/languages/code-size/farseer-go/data)![](https://img.shields.io/github/directory-file-count/farseer-go/data)![](https://goreportcard.com/badge/github.com/farseer-go/data)                                     |
| elasticSearch | client             | ![](https://img.shields.io/github/v/release/farseer-go/elasticSearch)![](https://img.shields.io/github/languages/code-size/farseer-go/elasticSearch)![](https://img.shields.io/github/directory-file-count/farseer-go/elasticSearch)![](https://goreportcard.com/badge/github.com/farseer-go/elasticSearch) |
| eventBus      | eventBus           | ![](https://img.shields.io/github/v/release/farseer-go/eventBus)![](https://img.shields.io/github/languages/code-size/farseer-go/eventBus)![](https://img.shields.io/github/directory-file-count/farseer-go/eventBus)![](https://goreportcard.com/badge/github.com/farseer-go/eventBus)                     |
| queue         | Local Queue        | ![](https://img.shields.io/github/v/release/farseer-go/queue)![](https://img.shields.io/github/languages/code-size/farseer-go/queue)![](https://img.shields.io/github/directory-file-count/farseer-go/queue)![](https://goreportcard.com/badge/github.com/farseer-go/queue)                                 |
| tasks         | Local tasks        | ![](https://img.shields.io/github/v/release/farseer-go/tasks)![](https://img.shields.io/github/languages/code-size/farseer-go/tasks)![](https://img.shields.io/github/directory-file-count/farseer-go/tasks)![](https://goreportcard.com/badge/github.com/farseer-go/tasks)                                 |
| fss           | client             | ![](https://img.shields.io/github/v/release/farseer-go/fss)![](https://img.shields.io/github/languages/code-size/farseer-go/fss)![](https://img.shields.io/github/directory-file-count/farseer-go/fss)![](https://goreportcard.com/badge/github.com/farseer-go/fss)                                         |
| utils         | utils              | ![](https://img.shields.io/github/v/release/farseer-go/utils)![](https://img.shields.io/github/languages/code-size/farseer-go/utils)![](https://img.shields.io/github/directory-file-count/farseer-go/utils)![](https://goreportcard.com/badge/github.com/farseer-go/utils)                                 |
| linkTrack     | linkTrack          | （Coming soon）                                                                                                                                                                                                                                                                                               |
| rabbit        | client             | （Coming soon）                                                                                                                                                                                                                                                                                               |
| redisStream   | redis mq           | （Coming soon）                                                                                                                                                                                                                                                                                               |

## Quick Start

_main.go_
```go
package main
import "github.com/farseer-go/fs"

func main() {
	fs.Initialize[StartupModule]("your project Name")
}
```

> In the first line of the main function, execute `fs.Initialize` to start initializing the framework

After running the console prints the loading message.

```
2022-12-01 17:07:24 Application Name： your project Name
2022-12-01 17:07:24 Host Name： MacBook-Pro.local
2022-12-01 17:07:24 System time： 2022-12-01 17:07:24
2022-12-01 17:07:24   ProcessID： 6123
2022-12-01 17:07:24   ApplicationID： 193337022963818496
2022-12-01 17:07:24   ApplicationIP： 192.168.1.4
2022-12-01 17:07:24 ---------------------------------------
2022-12-01 17:07:24 Loading Module...
2022-12-01 17:07:24 Loading Module：webapi.Module
2022-12-01 17:07:24 Loading Module：domain.Module
2022-12-01 17:07:24 Loading Module：application.Module
2022-12-01 17:07:24 Loading Module：interfaces.Module
2022-12-01 17:07:24 Loading Module：data.Module
2022-12-01 17:07:24 Loading Module：eventBus.Module
2022-12-01 17:07:24 Loading Module：queue.Module
2022-12-01 17:07:24 Loading Module：infrastructure.Module
2022-12-01 17:07:24 Loading Module：main.StartupModule
2022-12-01 17:07:24 Loaded, 10 modules in total
2022-12-01 17:07:24 ---------------------------------------
2022-12-01 17:07:24 Initialization completed, total time: 1 ms 
2022-12-01 17:07:24 ---------------------------------------
2022-12-01 17:07:24 [Info] Web service is started：http://localhost:8888/
```
## Stargazers

[![Stargazers repo roster for @farseer-go/fs](https://reporoster.com/stars/farseer-go/fs)](https://github.com/farseer-go/fs/stargazers)

## Forks

[![Forks repo roster for @farseer-go/fs](https://reporoster.com/forks/farseer-go/fs)](https://github.com/farseer-go/fs/network/members)