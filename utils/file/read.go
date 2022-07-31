package file

import (
	"os"
	"strings"
)

// ReadString 读文件内容
func ReadString(filePath string) string {
	file, _ := os.ReadFile(filePath)
	return string(file)
}

// ReadAllLines 读文件内容，按行返回数组
func ReadAllLines(filePath string) []string {
	file, _ := os.ReadFile(filePath)
	return strings.Split(string(file), "\n")
}
