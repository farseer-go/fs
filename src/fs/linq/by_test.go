package linq

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	lstYaml := []string{"1", "", "2"}
	item := By(lstYaml).Remove("")
	fmt.Println(item)
}
