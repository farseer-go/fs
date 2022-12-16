package configure

type envConfig struct {
}

func NewEnvConfig() *envConfig {
	return &envConfig{}
}

func (r *envConfig) LoadConfigure() error {
	return nil
}

func (r *envConfig) GetString(key string) string {
	return ""
}

func (r *envConfig) Get(key string) (any, bool) {
	return nil, false
}
