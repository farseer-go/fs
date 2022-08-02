package redis

// 定义redis配置结构
type redisConfig struct {
	Name           string
	Server         string
	DB             int
	Password       string
	ConnectTimeout string
	SyncTimeout    string
}
