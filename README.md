## What is Farseer.Go?

---
[English](https://github.com/FarseerGo/Farseer.Go) | [中文](https://github.com/FarseerGo/Farseer.Go/blob/main/README.zh-cn.md)

A standard set of frameworks developed for the **golang** platform.

We have selected the most popular components for you and provide the use of these components in a modular way.

The framework perfectly supports **DDD domain-driven** technical implementations such as `repository`, `application-layer transactions`, `domain events`, `application-layer dynamic WebAPI`.

We hope that in the daily development, only need to directly rely on this set of frameworks can cope with the common technical components

It has a [.net core](https://github.com/FarseerNet/Farseer.Net/tree/dev/Doc) Mature version, has been used for more than 10 years, very good

> *The current version of golang is `Alpha`, which means the features are not perfect yet, we are working on adding new features to the framework gradually. You can follow us first*.

## What are the features?

---
**Elegant**

We use `IOC` technology throughout the framework and your business systems.

**Simple**

We use `AOP` technology so that you don't have to write additional non-business functional code such as transactions, caching, exception catching, logging, linking Track

**Lightweight**

The framework makes extensive use of `collection pooling` technology to make your application take up less memory.

**Tracking**

If you use Orm, Redis, Http, Grpc, Elasticsearch, MQ (Rabbit, RedisStream, Rocker, local Queue), EventBus, Task, FSS, etc. that we provide, you don't need to do anything, the system will implicitly implement link tracking for you and provide API request logs, slow queries (all of the previously mentioned will be logged).

[FOPS](https://github.com/FarseerNet/FOPS) Project (automatic build, link trace console, K8S cluster log collection) supports code non-intrusive full link real-time monitoring.

## What are the package?

---
| Component       | Description                                                          |
|-----------------|----------------------------------------------------------------------|
| cache/redis     | redis client                                                         |
| configure       | config read write                                                    |
| core/container  | Registration and acquisition of Ioc containers                       |
| data            | Database ORM                                                         |
| elasticSearch   | elasticSearch client                                                 |
| eventBus        | Publish subscription for events                                      |
| fsApp           | Module Starter                                                       |
| linq            | Support linq methods                                                 |
| mapper          | Conversions between objects, such as DO to DTO                       |
| modules         | Module                                                               |
| mq/queue        | Local queue, multiple writes, bulk consumption, multiple subscribers |
| utils/encrypt   | encryption eg:md5                                                    |
| utils/exec      | shell cmd                                                            |
| utils/file      | File IO operations                                                   |
| utils/http      | Http GET POST                                                        |
| utils/net       | Get local IP                                                         |
| utils/parse     | Universal type conversion                                            |
| utils/stopwatch | Timer                                                                |
| utils/str       | String Operations                                                    |
| utils/times     | time.Time Operations                                                 |

## What are the functions?
### 框架初始化
* fs
    * `Initialize （初始化框架）`

* modules
    * `StartModules （启动模块）`

### 配置读写
* configure
    * GetString （获取配置）
    * SetDefault （设置配置的默认值）

---
### 容器IOC
* core/container
    * func
        * Register （注册接口）
        * Resolve （从容器中获取实例）

---
### 异常处理
* `exception`
    * `struct`
        * `RefuseException`
    * `func`
        * `ThrowRefuseException （抛出RefuseException异常）`
        * `ThrowRefuseExceptionf （抛出RefuseException异常）`
        * `Catch（捕获异常）`
            * `.RefuseException（捕获RefuseException异常）`
                * `.ContinueRecover（是否继续让下一个捕获继续处理）`
            * `.String（捕获String异常）`
            * `.Any（捕获Any异常）`
