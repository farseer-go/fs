package linq

import (
	"log"
	"testing"
)

func TestRemove(t *testing.T) {
	lstYaml := []string{"1", "", "2"}
	item := FromC(lstYaml).Remove("")
	log.Println(item)
}
