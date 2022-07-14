package directory

import (
	"fmt"
	"testing"
)

// ToInfo 项目组信息
func TestGetFiles(t *testing.T) {
	files := GetFiles("/Users/steden/Desktop/code/project/Farseer.Go", ".go", true)
	fmt.Println(files)
}
