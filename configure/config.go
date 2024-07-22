package configure

import (
	"github.com/farseer-go/fs/parse"
	"strings"
)

var configurationBuilder = newConfigurationBuilder()

// yml提供者
var ymlProvider *yamlConfig

type config struct {
	def            map[string]any    // 默认配置
	envKeyReplacer *strings.Replacer // 环境变量替换
	configProvider []IConfigProvider // 配置提供者
}

func newConfigurationBuilder() config {
	return config{
		def:            make(map[string]any),
		configProvider: []IConfigProvider{},
	}
}

// AddYamlFile 设置yaml文件配置
func (c *config) AddYamlFile(configFile string) {
	ymlProvider = NewYamlConfig(configFile)
	c.configProvider = append(c.configProvider, ymlProvider)
}

// AddEnvironmentVariables 加载环境变量
func (c *config) AddEnvironmentVariables() {
	c.configProvider = append([]IConfigProvider{NewEnvConfig()}, c.configProvider...)
}

// SetEnvKeyReplacer 环境变量替换
func (c *config) SetEnvKeyReplacer(r *strings.Replacer) {
	c.envKeyReplacer = r
}

// Build 找到并读取配置文件
func (c *config) Build() error {
	for _, provider := range c.configProvider {
		err := provider.LoadConfigure()
		if err != nil {
			return err
		}
	}
	return nil
}

// Get 获取配置
func (c *config) Get(key string) any {
	// 遍历配置提供者
	for _, provider := range c.configProvider {
		v, exists := provider.Get(key)
		if exists {
			return v
		}
	}

	// 是否有默认配置
	val, exists := c.def[key]
	if exists {
		return val
	}

	return nil
}

// GetSubNodes 获取所有子节点
func (c *config) GetSubNodes(key string) map[string]any {
	m := make(map[string]any)
	// 先加载默认值
	prefixKey := key + "."
	for k, v := range c.def {
		if strings.HasPrefix(k, prefixKey) {
			m[k[len(prefixKey):]] = v
		}
	}

	// 这里需要倒序获取列表，利用后面覆盖前面的方式来获取
	// 再添加yaml、环境变量
	for i := len(c.configProvider) - 1; i >= 0; i-- {
		if subMap, exists := c.configProvider[i].GetSubNodes(key); exists {
			for k, v := range subMap {
				// 尝试从之前的map中找到key（忽略大小写）
				// 目的是以yaml的key为准
				if c.configProvider[i].Name() == "env" {
					k = lookupMapKeyIgnoreCase(m, k)
				}
				m[k] = v
			}
		}
	}

	return m
}

// GetSlice 获取数组
func (c *config) GetSlice(key string) []string {
	var result []string
	// 先加载默认值
	if defVal, exists := c.def[key]; exists {
		switch defArr := defVal.(type) {
		case []string:
			result = defArr
		}
	}

	// 这里需要倒序获取列表，利用后面覆盖前面的方式来获取
	// 再添加yaml、环境变量
	for i := len(c.configProvider) - 1; i >= 0; i-- {
		if arrVal, exists := c.configProvider[i].GetArray(key); exists {
			for arrIndex, val := range arrVal {
				for len(result) <= arrIndex {
					result = append(result, "")
				}
				result[arrIndex] = parse.ToString(val)
			}
		}
	}
	return result
}

// GetSliceNodes 获取数组节点
func (c *config) GetSliceNodes(key string) []map[string]any {
	// 遍历配置提供者
	for _, provider := range c.configProvider {
		v, exists := provider.Get(key)
		if exists {
			arr, isOk := v.([]any)
			if isOk {
				var sliceNodes []map[string]any
				for _, node := range arr {
					sliceNodes = append(sliceNodes, node.(map[string]any))
				}
				return sliceNodes
			}
		}
	}
	return []map[string]any{}
}

// 尝试从之前的map中找到key（忽略大小写）
// 目的是以yaml的key为准
func lookupMapKeyIgnoreCase(m map[string]any, key string) string {
	for k, _ := range m {
		if strings.EqualFold(k, key) {
			return k
		}
	}
	return key
}
