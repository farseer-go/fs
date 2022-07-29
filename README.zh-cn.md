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
#### struct
* core
  * ApiResponseString （标准的API输出（默认string值））
  * ApiResponseInt （标准的API输出（默认int值））
  * ApiResponse （标准的API输出（泛型））
    * .SetData （设置Data字段的值）
  * PageList （用于分页数组，包含总记录数）
* data
  * dbConfig（数据库配置）
  * DbContext（数据库上下文）
  * TableSet （数据库表操作）
    * .SetTableName （设置表名）
    * .GetTableName （获取表名称）
    * .Select （筛选字段）
    * .Where （条件）
    * .Order （排序）
    * .Desc （倒序）
    * .Asc （正序）
    * .ToList （返回结果集）
    * .ToPageList （返回分页结果集）
    * .ToEntity （返回单个对象）
    * .Count （返回表中的数量）
    * .IsExists （是否存在记录）
    * .Insert （新增记录）
    * .Update （修改记录）
    * .UpdateValue （修改单个字段）
    * .Delete （删除记录）
    * .GetXXX （获取单个XXX类型字段值）
* eventBus
  * EventArgs （事件属性）
* linq
  * EventArgs （事件属性）
  * linqDictionary （针对字典的操作）
    * .ExistsKey （是否存在KEY）
  * linqForm （数据对集合数据筛选）
    * .Where （对数据进行筛选）
    * .Find （查找符合条件的元素）
    * .FindAll （查找符合条件的元素列表）
    * .First （查找符合条件的第一个元素）
    * .ToArray （查找符合条件的元素列表）
    * .RemoveAll （移除条件=true的元素）
    * .Count （获取数量）
    * .ToPageList （数组分页）
    * .Take （返回前多少条数据）
  * linqFormT （筛选子元素字段）
    * .Select （筛选子元素字段）
  * linqFromC （支持比较的集合）
    * .Where （对数据进行筛选）
    * .Contains （查找数组是否包含某元素）
    * .Remove （移除指定值的元素）
    * .Count （获取数量）
  * linqFromOrder （对集合进行排序）
    * .Where （对数据进行筛选）
    * .OrderBy （正序排序）
    * .OrderByDescending （倒序排序）
    * .Min （获取最小值）
    * .Max （获取最大值）

---

#### func
* configure
  * GetString （获取配置）
  * SetDefault （设置配置的默认值）
* core
  * Success （接口调用成功后返回的Json）
  * Error （接口调用失时返回的Json）
  * Error403 （接口调用失时返回的Json）
  * NewPageList （数据分页列表及总数）
* core/container
  * Register （注册接口）
  * Resolve （从容器中获取实例）
* data
  * NewDbContext （初始化上下文）
  * Init （数据库上下文初始化）
* eventBus
  * PublishEvent （阻塞发布事件）
  * PublishEventAsync （异步发布事件）
  * Subscribe （订阅事件）
* linq
  * Dictionary （针对字典的操作）
  * From （数据对集合数据筛选）
  * FromT （筛选子元素字段）
  * FromC （支持比较的集合）
  * FromOrder （对集合进行排序）
* mapper
  * Array （数组转换）
  * Single （单个转换）
* mq/queue
  * Push （添加数据到队列中）
  * Subscribe （订阅消息）
* utils/directory
  * GetFiles （读取指定目录下的文件）
  * CopyFolder （复制整个文件夹）
  * CopyFile （复制文件）
  * ClearFile （清空目录下的所有文件）
  * IsExists （判断路径是否存在）
* utils/encrypt
  * Md5 （对字符串做MD5加密）
* utils/exec
  * RunShell （执行shell命令）
* utils/net
    * LocalIPv4s （获取本机IP地址）
* utils/parse
    * Convert （通用的类型转换）
    * IsInt （是否为int类型）
* utils/str
  * CutRight （裁剪末尾标签）
  * MapToStringList （将map转成字符串数组）