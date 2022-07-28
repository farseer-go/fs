package parse

import "strconv"

// IsInt 判断source是否为int类型
func IsInt(source string) bool {
	_, err := strconv.Atoi(source)
	return err == nil
}
