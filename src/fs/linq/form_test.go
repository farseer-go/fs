package linq

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).Find(func(item int) bool {
		return item == 3
	})
	fmt.Println(item)
}
func TestRemove(t *testing.T) {
	lstYaml := []string{"1", "", "2"}
	item := From(lstYaml).Remove("")
	fmt.Println(item)
}
