package flog

type Config struct {
	Component componentConfig
	Default   levelFormat
	Console   levelFormat // 输出到控制台
	File      fileConfig  // 写到文件
	Fops      levelFormat // 上传到FOPS
}

// 组件日志
type componentConfig struct {
	Task        bool
	CacheManage bool
}

type levelFormat struct {
	LogLevel   string // 只记录的等级
	Format     string // 记录格式
	TimeFormat string // 时间格式
	Disable    bool   // 停用
}

type fileConfig struct {
	levelFormat
	Path            string // 日志存放的目录位置
	FileName        string // 日志文件名称
	RollingInterval string // 日志滚动间隔
	FileSizeLimitMb int64  // 日志文件大小限制（MB）
	FileCountLimit  int    // 日志文件数量限制
	RefreshInterval int    // 写入到文件的时间间隔，最少为1（秒）
}
