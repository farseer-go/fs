package linq

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).Where(func(item int) bool {
		return item == 3
	}).ToArray()
	fmt.Println(item)
}

func TestToPageList(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).ToPageList(3, 2)
	fmt.Println(item)
}

func TestTake(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).Take(3)
	fmt.Println(item)
}
