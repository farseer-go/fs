package configure

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type yamlConfig struct {
	data       map[string]any // 从yaml读取的数据
	configFile string         // 配置文件路径
}

func NewYamlConfig(configFile string) *yamlConfig {
	return &yamlConfig{
		data:       make(map[string]any),
		configFile: configFile,
	}
}
func (r *yamlConfig) Name() string {
	return "yaml"
}

func (r *yamlConfig) LoadConfigure() error {
	data, err := os.ReadFile(r.configFile)
	if err != nil {
		return err
	}
	return r.LoadContent(data)
}

// LoadContent 解析yml内容
func (r *yamlConfig) LoadContent(data []byte) error {
	var m map[string]any
	err := yaml.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	// 结构化转成扁平化
	r.flattening("", m)
	return nil
}

func (r *yamlConfig) Get(key string) (any, bool) {
	v, exists := r.data[key]
	if exists {
		switch v.(type) {
		case map[string]any:
			data, err := yaml.Marshal(&v)
			if err == nil {
				return string(data), exists
			}
		}
	}
	return v, exists
}

func (r *yamlConfig) GetArray(key string) ([]any, bool) {
	v, exists := r.data[key]
	if exists {
		switch arr := v.(type) {
		case []any:
			return arr, exists
		}
	}
	return nil, exists
}

func (r *yamlConfig) GetSubNodes(key string) (map[string]any, bool) {
	v, exists := r.data[key]
	if exists {
		switch m := v.(type) {
		case map[string]any:
			return m, exists
		}
	}
	return nil, false
}

// 扁平化map
func (r *yamlConfig) flattening(keyPrefix string, m map[string]any) {
	// 遍历节点
	for k, v := range m {
		// 与之前的key，组合成:a.b形式
		var key = k
		if keyPrefix != "" {
			key = keyPrefix + "." + k
		}
		r.flatteningAny(key, v)
	}
}

// 扁平化any
func (r *yamlConfig) flatteningAny(key string, v any) {
	switch subNode := v.(type) {
	// 需要继续往里面遍历子节点map
	case map[string]any:
		r.data[key] = subNode
		r.flattening(key, subNode)
	// 需要继续往里面遍历子节点数组
	case []any:
		r.data[key] = subNode
		for subIndex := 0; subIndex < len(subNode); subIndex++ {
			r.flatteningAny(key+fmt.Sprintf("[%d]", subIndex), subNode[subIndex])
		}
	default:
		r.data[key] = v
	}
}
