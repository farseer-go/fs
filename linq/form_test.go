package linq

import (
	"log"
	"testing"
)

func TestFind(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).Where(func(item int) bool {
		return item == 3
	}).ToArray()
	log.Println(item)
}

func TestToPageList(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).ToPageList(3, 2)
	log.Println(item)
}

func TestTake(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	item := From(lst).Take(3)
	log.Println(item)
}
