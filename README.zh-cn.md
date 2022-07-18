## Farseer.Go是什么?

---
[English](https://github.com/FarseerGo/Farseer.Go) | [中文](https://github.com/FarseerGo/Farseer.Go/blob/main/README.zh-cn.md)

针对 **golang** 平台下的一套标准制定的框架。

我们为您选型出目前最为流行的组件，并按模块化来提供使用这些组件。

框架完美支持 **DDD领域驱动** 的技术实现，如`仓储资源库`、`应用层事务`、`领域事件`、`应用层动态WebAPI`。

我们希望，在日常开发中，只需要直接依赖这一套框架便可应付常用的技术组件

它有一个[.net core](https://github.com/FarseerNet/Farseer.Net/tree/dev/Doc) 成熟版本，已经使用了10多年，非常棒

> *目前golang版本是`Alpha`,意味着功能还不完善，我们在努力向框架中逐渐添加新功能。您可以先关注我们*

### 有什么特点？

---
**优雅**

我们使用`IOC`技术，遍布整个框架及您的业务系统。

**简单**

我们使用`AOP`技术，让您无需额外编写非业务功能代码，如事务、缓存、异常捕获、日志、链路Track

**轻量**

框架内大量使用`集合池化`技术，使您的应用占用内存更小。

**链路追踪**

如果您使用我们提供的Orm、Redis、Http、Grpc、Elasticsearch、MQ(Rabbit、RedisStream、Rocker、本地Queue)、EventBus、Task、FSS等等，您什么都不需要做，系统将隐式为您实现链路追踪，并提供API请求日志、慢查询（前面提到的都会记录）。

结合[FOPS](https://github.com/FarseerNet/FOPS) 项目（自动构建、链路追踪控制台、K8S集群日志收集）支持代码无侵入的全链路实时监控。

### 有哪些功能？

---
| 组件名称               | 描述                                                                                           |
|--------------------|----------------------------------------------------------------------------------------------|
| fs                 | 启动器                                                                                          |
| fs/core/container  | Ioc容器的注册与获取                                                                                  |
| fs/mapper          | 对象间的转换，如DO转DTO                                                                               |
| fs/eventBus        | 事件的发布订阅                                                                                      |
| fs/mq/queue        | 本地队列，多次写入、批量消费，多个订阅者                                                                         |
| fs/linq            | 支持linq的where、first、toArray、remove、removeAll、contains、orderBy、orderByDescending、min、max、count |
| fs/utils/directory | 获取目录下的文件                                                                                     |
| fs/utils/encrypt   | MD5加密                                                                                        |
| fs/utils/net       | 获取本机IP                                                                                       |