package core

import "time"

// IConnectionChecker 统一的连接检查接口
// 用于检查配置字符串是否能成功连接到相应的服务（数据库、缓存、消息队列等）
type IConnectionChecker interface {
	// Check 检查连接字符串是否能成功连接到服务
	// configString: 连接配置字符串
	// 返回值：(连接成功, 错误信息)
	Check(configString string) (bool, error)

	// CheckWithTimeout 带超时时间的连接检查
	// configString: 连接配置字符串
	// timeout: 超时时间，如果为0则使用默认超时
	// 返回值：(连接成功, 错误信息)
	CheckWithTimeout(configString string, timeout time.Duration) (bool, error)
}
