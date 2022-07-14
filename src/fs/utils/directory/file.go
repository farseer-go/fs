package directory

import (
	"os"
	"path/filepath"
	"strings"
)

// GetFiles 读取指定目录下的文件
func GetFiles(path string, searchPattern string, searchSubDir bool) []string {
	var files []string
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if path == filePath {
			return nil
		}
		if info.IsDir() { // 如果是目录，并且要搜索子目录时，则用递归
			if searchSubDir {
				subFiles := GetFiles(filePath, searchPattern, true)
				files = append(files, subFiles...)
			}
		} else {
			if searchPattern != "" { // 带有文件名称过滤
				if strings.Contains(info.Name(), searchPattern) {
					files = append(files, filePath)
				}
			} else {
				files = append(files, filePath)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
