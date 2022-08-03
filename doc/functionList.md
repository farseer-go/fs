### 框架初始化
* fsApp
  * `Initialize （Initialize）`

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
### ORM
* data
    * struct
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
    * func
        * NewDbContext （初始化上下文）
        * Init （数据库上下文初始化）

---    
### linq
* linq
    * func
        * Dictionary （针对字典的操作）
            * .ExistsKey （是否存在KEY）
        * Form （数据对集合数据筛选）
            * .Where （对数据进行筛选）
            * .First （查找符合条件的第一个元素）
            * .ToArray （查找符合条件的元素列表）
            * .RemoveAll （移除条件=true的元素）
            * .RemoveItem （移除指定值的元素）
            * .Count （获取数量）
            * .ToPageList （数组分页）
            * .Take （返回前多少条数据）
            * `.ContainsItem （查找数组是否包含某元素）`
            * .Select （筛选子元素字段）
            * .OrderBy （正序排序）
            * `.OrderByItem （正序排序）`
            * .OrderByDescending （倒序排序）
            * `.OrderByDescendingItem （倒序排序）`
            * .Min （获取最小值）
            * .MinItem （获取最小值）
            * .Max （获取最大值）
            * .MaxItem （获取最大值）
            * `.Sum （求总和）`
            * `.SumItem （求总和）`
            * `.Avg （求平均数）`
            * `.AvgItem （求平均数）`
            * `GroupBy （将数组进行分组后返回map）`
            * `Any （是否存在）`
            * `All （是否所有数据都满足fn条件）`

---
### 事件总线
* eventBus
    * struct
        * EventArgs （事件属性）
    * func
        * PublishEvent （阻塞发布事件）
        * PublishEventAsync （异步发布事件）
        * Subscribe （订阅事件）

---
### 对象转换
* mapper
    * func
        * Array （数组转换）
        * Single （单个转换）

---
### 本地队列
* mq/queue
    * func
        * Push （添加数据到队列中）
        * Subscribe （订阅消息）

---
### 通用类型
* core
    * struct
        * ApiResponseString （标准的API输出（默认string值））
        * ApiResponseInt （标准的API输出（默认int值））
        * ApiResponseLong （标准的API输出（默认int64值））
        * ApiResponse （标准的API输出（泛型））
            * .SetData （设置Data字段的值）
        * PageList （用于分页数组，包含总记录数）
    * func
        * Success （接口调用成功后返回的Json）
        * Error （接口调用失时返回的Json）
        * Error403 （接口调用失时返回的Json）
        * NewPageList （数据分页列表及总数）
* `core/eumLogLevel`
  * `Enum （日志等级）`

---
### 文件操作
* utils/file
    * func
        * GetFiles （读取指定目录下的文件）
        * CopyFolder （复制整个文件夹）
        * CopyFile （复制文件）
        * ClearFile （清空目录下的所有文件）
        * IsExists （判断路径是否存在）
        * `Delete （删除文件）`
        * `WriteString （写入文件）`
        * `AppendString （追加文件）`
        * `AppendLine （换行追加文件）`
        * `AppendAllLine （换行追加文件）`
        * `CreateDir766 （创建所有目录，权限为766）`
        * `CreateDir （创建所有目录）`
        * `ReadString （读文件内容）`
        * `ReadAllLines （读文件内容，按行返回数组）`

---
### 常用工具
* utils/encrypt
  * Md5 （对字符串做MD5加密）
* utils/exec
  * RunShell （执行shell命令）
  * RunShellContext （执行shell命令）
* utils/net
  * LocalIPv4s （获取本机IP地址）
* utils/parse
  * Convert （通用的类型转换）
  * IsInt （是否为int类型）
* utils/str
  * CutRight （裁剪末尾标签）
  * MapToStringList （将map转成字符串数组）
  * `ToDateTime （将时间转换为yyyy-MM-dd HH:mm:ss）`
* utils/http
  * `Post （http post，支持超时设置）`
  * `PostForm （http post，默认x-www-form-urlencoded）`
  * `PostFormWithoutBody （http post，默认x-www-form-urlencoded）`
  * `PostJson （Post方式将结果反序列化成TReturn）`
  * `Get （http get，支持超时设置）`
  * `GetForm （http get，默认x-www-form-urlencoded）`
  * `GetFormWithoutBody （http get，默认x-www-form-urlencoded）`
  * `GetJson （Get方式将结果反序列化成TReturn）`
* utils/times
  * `GetTime（根据time.Duration转换成天、小时、分钟、秒）`
  * `GetDesc（返回时间中文的描述）`
  * `GetSubDesc（返回时间中文的描述）`
---
### redis
* cache/redis
  * func
    * NewClient（初始化客户端）
      * Remove（删除key）
      * Exists（key是否存在）
  * struct
    * Key
      * TTL（获取过期时间）
      * Del（删除）
      * Exists（key值是否存在）
    * String
      * Set（设置缓存）
      * Get（获取缓存）
      * SetNX（设置过期时间）
    * Hash
      * Set（设置缓存）
      * SetEntity（添加并序列化成json）
      * SetMap（设置map缓存）
      * Get（获取单个field值）
      * GetAll（获取key下所有数据）
      * Exists（判断单个field是否存在）
      * Remove（移出指定field成员）
      * `ToList（将hash.value反序列化成切片对象）`
      * `ToEntity（获取单个对象）`
      * `Count（获取hash的数量）`
    * List
      * Push（添加）
      * Set（设置指定索引的值）
      * Rem（移除指定索引的值）
      * Len（获取长度）
      * Range（获取集合数据）
      * BLPop（命令移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。）
    * Set
      * Add（添加）
      * Card（获取数量）
      * Rem（移除指定成员）
      * Members（获取所有成员）
      * IsMember（判断成员是否存在）
      * Diff（获取差集）
      * DiffStore（将差集，保存在指定集合中）
      * Inter（获取交集）
      * InterStore（将交集，保存在指定集合中）
      * Union（获取并集）
      * UnionStore（将并集，保存在指定集合中）
    * ZSet
      * Add（添加）
      * Score（获取指定成员的score)
      * Range（获取有序集合指定区间内的成员）
      * RevRange（获取有序集合指定区间内的成员分数从高到低）
      * RangeByScore（获取指定分数区间的成员列表）