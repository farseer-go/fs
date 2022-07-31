package file

import "os"

// CreateDir766 创建所有目录，权限为766
func CreateDir766(path string) {
	os.MkdirAll(path, 0766)
}

// CreateDir 创建所有目录
func CreateDir(path string, perm os.FileMode) {
	os.MkdirAll(path, perm)
}
