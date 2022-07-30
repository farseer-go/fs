package file

import "os"

// Delete 删除文件
func Delete(filePath string) {
	os.Remove(filePath)
}
