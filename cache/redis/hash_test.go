package redis

import (
	"fmt"
	"testing"
)

//String 测试
func TestClientHash(t *testing.T) {
	client := NewClient("default")

	defer client.Key.Del("key_has1")
	defer client.Key.Del("key_has2")

	err := client.Hash.Set("key_has1", "name", "小丽")
	err_v2 := client.Hash.Set("key_has1", "age", 40, "address", "上海")

	if err == nil {
		fmt.Printf("设置key_has1值成功.\n")
	} else {
		fmt.Printf("设置key_has1值错误:%v\n", err)
	}

	if err_v2 == nil {
		fmt.Printf("设置key_has1 v2 值成功.\n")
	} else {
		fmt.Printf("设置key_has1 v2 值错误:%v\n", err_v2)
	}

	get, _ := client.Hash.Get("key_has1", "name")
	fmt.Printf("获取key_has1  单个 name 值成功:%v\n", get)

	all, _ := client.Hash.GetAll("key_has1")
	fmt.Printf("获取key_has1  所有 值成功:%v\n", all)

	exists, _ := client.Hash.Exists("key_has1", "age")
	fmt.Printf("age值是否存在:%v\n", exists)

	get2, _ := client.Hash.Get("key_has1", "age")
	fmt.Printf("获取key_has2  单个 age 值成功:%v\n", get2)

	remove, _ := client.Hash.Del("key_has1", "age")
	fmt.Printf("移出age成员:%v\n", remove)

	err2 := client.Hash.Set("key_has2", "key1", "value1", "key2", 222)
	if err2 == nil {
		fmt.Printf("设置key_has2值成功.\n")
	} else {
		fmt.Printf("设置key_has2值错误:%v\n", err2)
	}
	all2, _ := client.Hash.GetAll("key_has2")
	fmt.Printf("获取key_has2  所有 值成功:%v\n", all2)

	//SetMap
	umap := map[string]string{"user": "harlen", "city": "河南", "age": "30"}
	err3 := client.Hash.SetMap("key_has3", umap)
	if err3 == nil {
		fmt.Printf("设置key_has3值成功.\n")
	}
	all3, _ := client.Hash.GetAll("key_has3")
	fmt.Printf("获取key_has3  所有 值成功:%v\n", all3)
}
