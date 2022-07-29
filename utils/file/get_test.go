package file

import (
	"fmt"
	"strings"
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
	fmt.Println(strings.TrimRight("aaaabb", "ab"))
}
