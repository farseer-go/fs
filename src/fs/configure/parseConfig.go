package configure

import (
	"fs/utils/parse"
	"reflect"
	"strings"
)

// ParseConfig 解析字符串，转成配置
func ParseConfig[TConfig any](configString string) TConfig {

	var config = new(TConfig)
	configRefVal := reflect.ValueOf(config).Elem()

	if configRefVal.Type().Kind() != reflect.Struct {
		panic("泛型只能是struct结构")
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
	for i := 0; i < configRefVal.Type().NumField(); i++ {
		if configRefVal.Field(i).CanSet() {
			fieldName := strings.ToLower(configRefVal.Type().Field(i).Name)
			s, exists := configMap[fieldName]
			if exists {
				switch configRefVal.Type().Field(i).Type.Kind() {
				case reflect.Int:
					configRefVal.Field(i).SetInt(parse.Convert(s, int64(0)))
					break
				case reflect.Bool:
					configRefVal.Field(i).SetBool(parse.Convert(s, false))
					break
				default:
					configRefVal.Field(i).Set(reflect.ValueOf(s))
					break
				}

			}
		}
	}
	return *config
}
