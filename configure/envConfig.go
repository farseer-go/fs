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

func (r *envConfig) Get(key string) (any, bool) {
	key = r.replace(key)
	val, exists := os.LookupEnv(key)
	if exists {
		return val, true
	}
	return nil, false
}

func (r *envConfig) GetSubNodes(key string) (map[string]any, bool) {
	return nil, false
}

func (r *envConfig) replace(key string) string {
	key = strings.ReplaceAll(key, ".", "_")
	key = strings.ReplaceAll(key, "[", "_")
	key = strings.ReplaceAll(key, "]", "_")
	return key
}
