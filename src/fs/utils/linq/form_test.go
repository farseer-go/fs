package linq

import (
	"fmt"
	"testing"
)

// ToInfo 项目组信息
func TestFind(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).Find(func(item int) bool {
		return item == 3
	})
	fmt.Println(item)
}
