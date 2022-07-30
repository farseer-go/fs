package file

import (
	"fmt"
	"github.com/farseernet/farseer.go/utils/str"
	"path/filepath"
	"testing"
)

// GetFiles
func TestGetFiles(t *testing.T) {
	path := "/Users/steden/Desktop/code/project/Farseer.Go"
	files := GetFiles(path, "*.md", true)
	for _, filepath := range files {
		fmt.Println(filepath)
	}
}

// ClearFile
func TestClearFile(t *testing.T) {
	path := "/Users/steden/Desktop/code/project/Farseer.Go2"
	ClearFile(path)
}

// IsExists
func TestIsExists(t *testing.T) {
	path := "/Users/steden/Desktop/code/project/Farseer.Go3"
	fmt.Println(IsExists(path))
}

func TestOther(t *testing.T) {
	git := "https://github.com/FarseerNet/farseer.go.git"
	git = filepath.Base(git)
	git = str.CutRight(git, ".git")
	fmt.Println(git)
}
