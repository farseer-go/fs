package configure

import (
	"strings"
)

var configurationBuilder config

type config struct {
	def            map[string]any    // 默认配置
	envKeyReplacer *strings.Replacer // 环境变量替换
	configProvider []IConfigProvider // 配置提供者
}

func NewConfigurationBuilder() config {
	return config{
		def:            make(map[string]any),
		configProvider: []IConfigProvider{},
	}
}

// AddYamlFile 设置yaml文件配置
func (c *config) AddYamlFile(configFile string) {
	var yConfig IConfigProvider = NewYamlConfig(configFile)
	c.configProvider = append(c.configProvider, yConfig)
}

// AddEnvironmentVariables 加载环境变量
func (c *config) AddEnvironmentVariables() {
	var envConfig IConfigProvider = NewEnvConfig()
	c.configProvider = append([]IConfigProvider{envConfig}, c.configProvider...)
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

// GetString 读取配置
func (c *config) GetString(key string) string {
	// 遍历配置提供者
	for _, provider := range c.configProvider {
		v := provider.GetString(key)
		if v != "" {
			return v
		}
	}

	// 是否有默认配置
	val, exists := c.def[key]
	if exists {
		return val.(string)
	}

	return ""
}

// GetSubNodes 获取所有子节点
func (c *config) GetSubNodes(key string) map[string]any {
	// 遍历配置提供者
	for _, provider := range c.configProvider {
		v, exists := provider.Get(key)
		if exists {
			m, isOk := v.(map[string]any)
			if isOk {
				return m
			}
		}
	}
	return make(map[string]any)
}

// GetSlice 获取数组
func (c *config) GetSlice(key string) []string {
	// 遍历配置提供者
	for _, provider := range c.configProvider {
		v, exists := provider.Get(key)
		if exists {
			m, isOk := v.([]any)
			if isOk {
				var arr []string
				for _, s := range m {
					arr = append(arr, s.(string))
				}
				return arr
			}
		}
	}
	return []string{}
}
