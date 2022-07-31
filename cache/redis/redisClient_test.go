package redis

import (
	"fmt"
	"testing"
)

//测试
func TestClient(t *testing.T) {
	client := NewClient("default")
	err := client.string.Set("key1", "...12312")
	if err == nil {
		fmt.Printf("设置值:%v\n", "...12312")
	}
	get, err := client.string.Exists("key1")
	if err != nil {
		fmt.Printf("错误信息：%v\n", err)
	} else {
		fmt.Printf("是否存在：%v\n", get)
	}

	remove, _ := client.string.Remove("key1")
	fmt.Printf("是否删除：%v\n", remove)

	s, _ := client.string.Get("key1")
	fmt.Printf("获取值：%v\n", s)

	get2, err2 := client.string.Exists("key1")
	if err != nil {
		fmt.Printf("错误信息：%v\n", err2)
	} else {
		fmt.Printf("是否存在：%v\n", get2)
	}
}
