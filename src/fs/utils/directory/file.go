package directory

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// GetFiles 读取指定目录下的文件
// path：目录路径
// searchPattern：匹配文件名要包含的名称，搜索全部，传入""即可
// searchSubDir：是否要搜索子目录
func GetFiles(path string, searchPattern string, searchSubDir bool) []string {
	var files []string
	err := filepath.WalkDir(path, func(filePath string, dirInfo fs.DirEntry, err error) error {
		if path == filePath {
			return nil
		}
		// 目录不需要判断，filepath.Walk执行就包含递归了
		if !dirInfo.IsDir() && (searchPattern == "" || strings.Contains(dirInfo.Name(), searchPattern)) {
			files = append(files, filePath)
		} else if dirInfo.IsDir() && !searchSubDir {
			return fs.SkipDir
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

// CopyFolder 复制整个文件夹
// srcPath 要复制的原目录路径
// destPath 复制到位置的目录路径
func CopyFolder(srcPath string, destPath string) {
	// 如果位置路径最后不带/，则自动加上
	if srcPath[len(srcPath)-1] != '/' {
		srcPath += "/"
	}
	if destPath[len(destPath)-1] != '/' {
		destPath += "/"
	}

	err := filepath.WalkDir(srcPath, func(filePath string, dirInfo fs.DirEntry, err error) error {
		if srcPath == filePath {
			return nil
		}
		if dirInfo.IsDir() { // 如果是目录，则创建目录即可
			// 创建目标目录
			_ = os.MkdirAll(destPath+dirInfo.Name(), 0766)
		} else {
			destName := destPath + dirInfo.Name()
			CopyFile(filePath, destName)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// CopyFile 复制文件
// srcName 复制的原文件
// destName 复制到目标位置（需带上文件名）
func CopyFile(srcName string, destName string) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(destName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
}
