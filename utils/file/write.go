package file

import (
	"os"
)

// WriteString 写入文件
func WriteString(filePath string, content string) {
	os.WriteFile(filePath, []byte(content), 0766)
}
