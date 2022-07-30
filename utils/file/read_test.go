package file

import (
	"fmt"
	"testing"
)

func TestReadString(t *testing.T) {
	file := "/Users/steden/Desktop/code/project/Farseer.Go/go.mod"
	fmt.Println(ReadString(file))
}
