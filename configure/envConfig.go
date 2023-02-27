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
	m := make(map[string]any)
	prefixKey := key + "_"
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, prefixKey) {
			index := strings.Index(env, "=")
			k := env[len(prefixKey):index]
			v := env[index+1:]
			m[k] = v
		}
	}
	return m, len(m) > 0
}

func (r *envConfig) replace(key string) string {
	key = strings.ReplaceAll(key, ".", "_")
	key = strings.ReplaceAll(key, "[", "_")
	key = strings.ReplaceAll(key, "]", "_")
	return key
}
