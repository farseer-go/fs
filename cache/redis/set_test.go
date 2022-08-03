package redis

import (
	"fmt"
	"testing"
)

func Test_redisSet(t *testing.T) {

	client := NewClient("default")
	defer client.Key.Del("key_set")
	defer client.Key.Del("key_set2")
	defer client.Key.Del("key_set_diff")
	defer client.Key.Del("key_set_inter")
	defer client.Key.Del("key_set_union")
	add2, err2 := client.Set.Add("key_set2", "小丽", "小王", "小白", "小赵", "小钱", "小孙")
	fmt.Println("添加2返回结果：", add2, err2)
	//添加
	add, err := client.Set.Add("key_set", "小丽", "小王", "小白", "小小")
	fmt.Println("添加返回结果：", add, err)

	card, err := client.Set.Card("key_set")
	fmt.Println("获取数量：", card, err)

	rem, err := client.Set.Rem("key_set", "小王")
	fmt.Println("移除指定成员返回结果：", rem, err)

	members, err := client.Set.Members("key_set")
	fmt.Println("获取所有成员：", members, err)

	member, err := client.Set.IsMember("key_set", "小白")
	fmt.Println("判断指定成员是否存在：", member, err)

	diff, err2 := client.Set.Diff("key_set", "key_set2")
	fmt.Println("获取差集：", diff, err2)

	store, err2 := client.Set.DiffStore("key_set_diff", "key_set", "key_set2")
	fmt.Println("存储差集到指定集合：", store, err2)

	inter, err2 := client.Set.Inter("key_set", "key_set2")
	fmt.Println("获取交集：", inter, err2)

	interStore, err2 := client.Set.InterStore("key_set_inter", "key_set", "key_set2")
	fmt.Println("存储交集到指定集合：", interStore, err2)

	union, err2 := client.Set.Union("key_set", "key_set2")
	fmt.Println("获取并集：", union, err2)

	unionStore, err2 := client.Set.UnionStore("key_set_union", "key_set", "key_set2")
	fmt.Println("存储并集到指定集合：", unionStore, err2)

}
