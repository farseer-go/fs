package redis

type redisConfig struct {
	Name           string
	Server         string
	DB             int
	Password       string
	ConnectTimeout string
	SyncTimeout    string
}
