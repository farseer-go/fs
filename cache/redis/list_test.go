package redis

import (
	"fmt"
	"testing"
	"time"
)

func Test_redisList(t *testing.T) {

	client := NewClient("default")
	defer client.Key.Del("key_list")

	//测试push
	push, err := client.List.Push("key_list", "北京", "上海", "广州", "海南", "河南")
	fmt.Println("添加返回结果", push, err)

	set, err := client.List.Set("key_list", 0, "深圳")
	fmt.Println("设置指定值返回结果：", set, err)

	rem, err := client.List.Rem("key_list", 0, "上海")
	fmt.Println("移除指定值返回结果：", rem, err)

	i, err := client.List.Len("key_list")
	fmt.Println("获取指定长度返回结果：", i, err)

	strings, err := client.List.Range("key_list", 0, i-1)
	fmt.Println("遍历key下所有数据：", strings, err)

	pop, err := client.List.BLPop(3*time.Second, "key_list")
	fmt.Println("没有阻塞的情况下返回结果：", pop, err)
}
