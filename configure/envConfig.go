package configure

import (
	"os"
	"strings"
)

type envConfig struct {
}

func NewEnvConfig() *envConfig {
	return &envConfig{}
}

func (r *envConfig) LoadConfigure() error {
	return nil
}

func (r *envConfig) GetString(key string) string {
	key = strings.ReplaceAll(key, ".", "_")
	key = strings.ReplaceAll(key, "[", "_")
	key = strings.ReplaceAll(key, "]", "_")
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return ""
}

func (r *envConfig) Get(key string) (any, bool) {
	return nil, false
}
