package file

import (
	"os"
)

// ReadString 读文件内容
func ReadString(filePath string) string {
	file, _ := os.ReadFile(filePath)
	return string(file)
}
