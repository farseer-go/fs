package file

import "testing"

// CopyFolder
func TestCopyFolder(t *testing.T) {
	path1 := "/Users/steden/Desktop/code/project/Farseer.Go"
	path2 := "/Users/steden/Desktop/code/project/Farseer.Go2"

	CopyFolder(path1, path2)
}
