package configure

type IConfigProvider interface {
	// LoadConfigure 加载配置
	LoadConfigure() error
	// Get 读取配置
	Get(key string) (any, bool)
	// GetString 读取配置
	GetString(key string) string
}
