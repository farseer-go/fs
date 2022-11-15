## What is Farseer.Go?

[English](https://github.com/farseer-go/fs) | [中文](https://github.com/farseer-go/fs/blob/main/README.zh-cn.md)

A standard set of frameworks developed for the **golang** platform.

We have selected the most popular components for you and provide the use of these components in a modular way.

The framework perfectly supports **DDD domain-driven** technical implementations such as `repository`, `application-layer transactions`, `domain events`, `application-layer`, `dynamic WebAPI`.

We hope that in the daily development, only need to directly rely on this set of frameworks can cope with the common technical components

It has a [.net core](https://github.com/FarseerNet/Farseer.Net/tree/dev/Doc) Mature version, has been used for more than 10 years, very good

> *The current version of golang is `Alpha`, which means the features are not perfect yet, we are working on adding new features to the framework gradually. You can follow us first*.

## What are the features?

**Unified Configuration**

Global unified configuration management

**Elegant**

We use `IOC` technology throughout the framework and your business systems.

**Simple**

We use `AOP` technology so that you don't have to write additional non-business functional code such as transactions, caching, exception catching, logging, linking Track

**Lightweight**

The framework makes extensive use of `collection pooling` technology to make your application take up less memory.

**Tracking**

If you use Orm, Redis, Http, Grpc, Elasticsearch, MQ (Rabbit, RedisStream, Rocker, local Queue), EventBus, Task, FSS, etc. that we provide, you don't need to do anything, the system will implicitly implement link tracking for you and provide API request logs, slow queries (all of the previously mentioned will be logged).

[FOPS](https://github.com/FarseerNet/FOPS) Project (automatic build, link trace console, K8S cluster log collection) supports code non-intrusive full link real-time monitoring.

## What packages are available?

| Component                                                    | Description                                                          |
|--------------------------------------------------------------|----------------------------------------------------------------------|
| [async](https://github.com/farseer-go/async)                 | Parallel asynchronous execution, uniform await to get the result     |
| [cache](https://github.com/farseer-go/cache)                 | Multi-level cache                                                    |
| [cacheMemory](https://github.com/farseer-go/cacheMemory)     | Memory Cache                                                         |
| [collections](https://github.com/farseer-go/collections)     | Support for List collections and linq syntax                         |
| [data](https://github.com/farseer-go/data)                   | Database ORM                                                         |
| [elasticSearch](https://github.com/farseer-go/elasticSearch) | elasticSearch client                                                 |
| [eventBus](https://github.com/farseer-go/eventBus)           | Publish subscription for events                                      |
| [fs](https://github.com/farseer-go/fs)                       | Farseer Basic                                                        |
| [fss](https://github.com/farseer-go/fss)                     | fss client                                                           |
| [linkTrack](https://github.com/farseer-go/linkTrack)         | Link Tracking                                                        |
| [mapper](https://github.com/farseer-go/mapper)               | Conversions between objects, such as DO to DTO                       |
| [mvc](https://github.com/farseer-go/mvc)                     | web mvc                                                              |
| [queue](https://github.com/farseer-go/queue)                 | Local queue, multiple writes, bulk consumption, multiple subscribers |
| [rabbit](https://github.com/farseer-go/rabbit)               | rabbit client                                                        |
| [redis](https://github.com/farseer-go/redis)                 | redis client                                                         |
| [redisStream](https://github.com/farseer-go/redisStream)     | redisStream client                                                   |
| [tasks](https://github.com/farseer-go/tasks)                 | Local job                                                            |
| [utils](https://github.com/farseer-go/utils)                 | General Tools                                                        |
| [webapi](https://github.com/farseer-go/webapi)               | webapi mvc                                                           |

## What are the functions?
* [fs（框架初始化）](#how-to-start)
    * Initialize （初始化框架）
* [configure（配置读写）](configure/)
    * GetString （获取配置）
    * SetDefault （设置配置的默认值）
* [container（容器IOC）](container/)
    * func
        * Use （自定义注册）
          * Transient（临时模式（默认为单例模式））
          * Name（Ioc别名）
          * Register （注册到容器）
        * Register （单例且没有别名注册到容器）
        * Resolve （从容器中获取实例）
        * ResolveName （指定ioc别名从容器中获取实例）
* [core（通用类型）](core/)
    * struct
        * ApiResponseString （标准的API输出（默认string值））
        * ApiResponseInt （标准的API输出（默认int值））
        * ApiResponseLong （标准的API输出（默认int64值））
        * ApiResponse （标准的API输出（泛型））
            * .SetData （设置Data字段的值）
    * func
        * Success （接口调用成功后返回的Json）
        * Error （接口调用失时返回的Json）
        * Error403 （接口调用失时返回的Json）
    * core/eumLogLevel
        * Enum （日志等级）
* [dateTime（时间日期）](dateTime/)
  * func
    * ToString（转字符串）
    * Now（当前时间）
    * New（初始化）
    * Year（获取年）
    * Month（获取月）
    * Day（获取日）
    * Hour（获取小时）
    * Minute（获取分钟）
    * Second（获取秒）
    * Date（获取Date部份）
    * AddYears（添加年）
    * AddMonths（添加月份）
    * AddDays（添加天数）
    * AddHours（添加小时）
    * AddMinutes（添加分钟）
    * AddSeconds（添加秒）
    * AddDate（添加Date）
    * AddTime（添加Time）
    * ToTime（获取time.Time类型）
* [exception（异常处理）](exception/)
    * struct
        * RefuseException
    * func
        * ThrowRefuseException （抛出RefuseException异常）
        * ThrowRefuseExceptionf （抛出RefuseException异常）
        * Catch（捕获异常）
            * .RefuseException（捕获RefuseException异常）
                * .ContinueRecover（是否继续让下一个捕获继续处理）
            * .String（捕获String异常）
            * .Any（捕获Any异常）
* [flog（日志打印）](flog/)
    * func
        * Trace（打印Trace日志）
        * Debug（打印Debug日志）
        * Info（打印Info日志）
        * Warning（打印Warning日志）
        * Error（打印Error日志）
        * Critical（打印Critical日志）
        * Log（打印日志）
        * Print（打印日志）
* [modules](modules/)
    * StartModules （启动模块）
* [net](net/)
    * LocalIPv4s （获取本机IP地址）
* [parse（类型转换）](parse/)
    * Convert （通用的类型转换）
    * IsInt （是否为int类型）
    * IsEqual（两个any值是否相等）
* [snowflake（雪花算法）](snowflake/)
    * Init（全局初始化一次)
    * GenerateId（生成唯一ID）
* [stopwatch](stopwatch/)
    * func
        * StartNew（创建计时器，并开始计时）
    * struct
        * Stopwatch
            * Restart（重置计时器）
            * Start（继续计时）
            * Stop（停止计时）
            * ElapsedMilliseconds（返回当前已计时的时间（毫秒））
* [types](types/)
  * func
    * GetRealType（获取真实类型）
    * IsSlice（是否为切片类型）
    * IsMap（是否为Map类型）
    * IsStruct（是否为Struct）
    * IsList（判断类型是否为List）
    * IsDictionary（是否为Dictionary）
    * IsPageList（是否为PageList）

## How to start？
StartupModule is the startup module you define, as detailed in：[modules](modules/)
```go
// The framework starts and performs module initialization
fs.Initialize[StartupModule]("FOPS")
```

## Log Print
`Log.Component` Set the switch for component printing logs

if true, the Component will print Detailed Log
```yaml
Log:
  LogLevel: "Information"
  Component:
    task: true
    cacheManage: true
    webapi: true
    event: true
    httpRequest: true
    queue: true
    fss: true
```