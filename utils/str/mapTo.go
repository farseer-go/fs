package str

import "github.com/farseernet/farseer.go/utils/parse"

// MapToStringList 将map转成字符串数组，map的kv转成：k=v格式
func MapToStringList[TKey comparable, TValue any](maps map[TKey]TValue) []string {
	var result []string
	for key, value := range maps {
		result = append(result, parse.Convert(key, "")+"="+parse.Convert(value, ""))
	}
	return result
}
