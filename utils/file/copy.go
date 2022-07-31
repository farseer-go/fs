package file

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

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

	// 先创建目标的目录
	stat, _ := os.Stat(srcPath)
	_ = os.MkdirAll(destPath, stat.Mode().Perm())

	err := filepath.WalkDir(srcPath, func(filePath string, dirInfo fs.DirEntry, err error) error {
		if srcPath == filePath {
			return nil
		}
		newPath := filepath.Join(destPath, filePath[len(srcPath):])
		if dirInfo.IsDir() { // 如果是目录，则创建目录即可
			// 创建目标目录
			perm, _ := dirInfo.Info()
			_ = os.MkdirAll(newPath, perm.Mode())
		} else {
			CopyFile(filePath, newPath)
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
	stat, _ := src.Stat()
	dst, err := os.OpenFile(destName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, stat.Mode())
	if err != nil {
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
}
