package configure

import (
	"fmt"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/types"
	"reflect"
	"strings"
)

// ParseString 解析字符串，转成配置对象
func ParseString[TConfig any](configString string) TConfig {
	configRefVal := reflect.New(reflect.TypeOf(new(TConfig)).Elem()).Elem()
	return parseString(configRefVal, configString).Interface().(TConfig)
}

func parseString(configRefVal reflect.Value, configString string) reflect.Value {
	configRefType := configRefVal.Type()
	if configRefType.Kind() != reflect.Struct {
		panic("A generic type can only be a struct")
	}

	// 第一步：字符串转成map
	configMap := make(map[string]string)
	configs := strings.Split(configString, ",")
	for _, configSplit := range configs {
		kv := strings.Split(configSplit, "=")
		// 最少需要一个=符号
		if len(kv) == 1 {
			continue
		}
		configMap[strings.ToLower(kv[0])] = configSplit[strings.Index(configSplit, "=")+1:]
	}

	// 第二步：反射TConfig结构
	for i := 0; i < configRefVal.NumField(); i++ {
		if configRefVal.Field(i).CanSet() {
			fieldName := strings.ToLower(configRefType.Field(i).Name)
			s, exists := configMap[fieldName]
			if exists {
				value := parse.ConvertValue(s, configRefType.Field(i).Type)
				configRefVal.Field(i).Set(reflect.ValueOf(value))
			}
		}
	}
	return configRefVal
}

// ParseConfigs 将配置转换成数组对象
func ParseConfigs[TConfig any](key string) []TConfig {
	var configs []TConfig
	nodes := GetSliceNodes(key)
	// 遍历配置节点
	for nodeIndex := 0; nodeIndex < len(nodes); nodeIndex++ {
		config := ParseConfig[TConfig](fmt.Sprintf("%s[%d]", key, nodeIndex))
		configs = append(configs, config)
	}
	return configs
}

// ParseConfig 将配置转换成对象
func ParseConfig[TConfig any](key string) TConfig {
	config := new(TConfig)
	// 用于反射赋值
	configVal := reflect.ValueOf(config).Elem()
	node := GetSubNodes(key)

	// 遍历配置字段，进行逐个赋值
	for nodeKey, fieldConfigValue := range node {
		field, isExists := configVal.Type().FieldByNameFunc(func(fieldName string) bool {
			return strings.ToLower(fieldName) == strings.ToLower(nodeKey)
		})
		if isExists && field.IsExported() {
			fieldVal := parseConfig(field.Type, fieldConfigValue)
			configVal.FieldByName(field.Name).Set(fieldVal)
		}
	}
	return *config
}

func parseConfig(configType reflect.Type, node any) reflect.Value {
	// 字段是基础类型
	if types.IsGoBasicType(configType) {
		return reflect.ValueOf(parse.ConvertValue(node, configType))
	}

	configFieldValue := reflect.New(configType).Elem()
	// 确定这个节点是string类型还是map类型
	switch nodeVal := node.(type) {
	// map类型需要遍历配置字段，进行逐个解析地赋值。
	case map[string]any:
		// 遍历这个结构的字段
		for nodeKey, fieldConfigValue := range nodeVal {
			field, isExists := configType.FieldByNameFunc(func(fieldName string) bool {
				return strings.ToLower(fieldName) == strings.ToLower(nodeKey)
			})
			if isExists && field.IsExported() {
				configFieldValue.FieldByName(field.Name).Set(parseConfig(field.Type, fieldConfigValue))
			}
		}
	// 数组类型
	case []any:
		if configType.Kind() == reflect.Slice || configType.Kind() == reflect.Array {
			// 遍历数组
			for _, node := range nodeVal {
				configFieldValue = reflect.Append(configFieldValue, parseConfig(configType.Elem(), node))
			}
			return configFieldValue
		}
	case string:
		configFieldValue.Set(parseString(configFieldValue, nodeVal))
	}
	return configFieldValue
}
