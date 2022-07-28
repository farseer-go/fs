package directory

import (
	"fmt"
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

// CopyFolder
func TestCopyFolder(t *testing.T) {
	path1 := "/Users/steden/Desktop/code/project/Farseer.Go"
	path2 := "/Users/steden/Desktop/code/project/Farseer.Go2"

	CopyFolder(path1, path2)
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
