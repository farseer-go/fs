package file

import (
	"os"
)

// WriteString 写入文件
func WriteString(filePath string, content string) {
	os.WriteFile(filePath, []byte(content), 0766)
}

// AppendString 追加文件
func AppendString(filePath string, content string) {
	oldContent := ReadString(filePath)
	os.WriteFile(filePath, []byte(oldContent+content), 0766)
}

// AppendLine 换行追加文件
func AppendLine(filePath string, content string) {
	oldContent := ReadString(filePath)
	os.WriteFile(filePath, []byte(oldContent+"\n"+content), 0766)
}
