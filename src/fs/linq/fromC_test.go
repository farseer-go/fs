package linq

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	lstYaml := []string{"1", "", "2"}
	item := FromC(lstYaml).Remove("")
	fmt.Println(item)
}
