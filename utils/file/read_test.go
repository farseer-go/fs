package file

import (
	"log"
	"testing"
)

func TestReadString(t *testing.T) {
	file := "/Users/steden/Desktop/code/project/Farseer.Go/go.mod"
	log.Println(ReadString(file))
}
