package directory

import (
	"fmt"
	"testing"
)

// ToInfo 项目组信息
func TestGetFiles(t *testing.T) {
	path := "/Users/steden/Desktop/code/project/Farseer.Go"
	//dir, _ := ioutil.ReadDir(path)
	//for _, fileInfo := range dir {
	//	fmt.Println(fileInfo.Name())
	//}
	files := GetFiles(path, ".go", true)
	for _, filepath := range files {
		fmt.Println(filepath)
	}
}
