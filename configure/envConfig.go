package configure

import (
	"github.com/farseer-go/fs/parse"
	"os"
	"strings"
)

type envConfig struct {
}

func NewEnvConfig() *envConfig {
	return &envConfig{}
}

func (r *envConfig) Name() string {
	return "env"
}

func (r *envConfig) LoadConfigure() error {
	return nil
}

func (r *envConfig) Get(key string) (any, bool) {
	key = r.replace(key)
	val, exists := lookupEnvIgnoreCase(key)
	if exists {
		return val, true
	}
	return nil, false
}

func (r *envConfig) GetArray(key string) ([]any, bool) {
	prefixKey := strings.ToLower(r.replace(key) + "_")
	var result []any
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) != 2 {
			continue
		}

		// 有指定key的前缀，说明是数组
		envKey := strings.ToLower(r.replace(pair[0]))
		if strings.HasPrefix(envKey, prefixKey) {
			// 得到索引
			arrIndex := parse.ToInt(envKey[len(prefixKey):])
			for len(result) <= arrIndex {
				result = append(result, "")
			}
			result[arrIndex] = pair[1]
		}
	}
	return result, len(result) > 0
}

func (r *envConfig) GetSubNodes(key string) (map[string]any, bool) {
	m := make(map[string]any)
	prefixKey := strings.ToLower(r.replace(key) + "_")
	for _, e := range os.Environ() {
		if strings.HasPrefix(strings.ToLower(e), prefixKey) {
			index := strings.Index(e, "=")
			k := e[len(prefixKey):index]
			v := e[index+1:]
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

func lookupEnvIgnoreCase(key string) (string, bool) {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			envKey := pair[0]
			if strings.EqualFold(envKey, key) {
				return pair[1], true
			}
		}
	}
	return "", false
}
