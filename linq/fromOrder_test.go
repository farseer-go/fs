package linq

import (
	"fmt"
	"testing"
)

// 正序排序
func TestOrderBy(t *testing.T) {
	lst := []int{3, 5, 6, 2, 1, 8, 7, 4}
	item := FromOrder[int, int](lst).OrderBy(func(item int) int {
		return item
	})
	fmt.Println(item)
}

// 倒序排序
func TestOrderByDescending(t *testing.T) {
	lst := []int{3, 5, 6, 2, 1, 8, 7, 4}
	item := FromOrder[int, int](lst).OrderByDescending(func(item int) int {
		return item
	})
	fmt.Println(item)
}
