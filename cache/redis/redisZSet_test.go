package redis

import (
	"fmt"
	"testing"
)

func Test_redisZSet(t *testing.T) {

	client := NewClient("default")
	defer client.Key.Del("key_set_z")
	//测试
	add, err := client.ZSet.Add("key_set_z", &redisZ{1, "小猫"}, &redisZ{2, "小狗"}, &redisZ{3, "小鸟"})
	fmt.Println("添加返回结果：", add, err)

	score, err := client.ZSet.Score("key_set_z", "小狗")
	fmt.Println("返回指定成员的score:", score, err)

	strings, err := client.ZSet.Range("key_set_z", 0, 1)
	fmt.Println("获取所有集合：", strings, err)

	revRange, err := client.ZSet.RevRange("key_set_z", 0, 3)
	fmt.Println("有序集合指定区间内的成员分数从高到低：", revRange, err)

	byScore, err := client.ZSet.RangeByScore("key_set_z", &redisZRangeBy{Max: "3", Min: "1", Offset: 1, Count: 3})
	fmt.Println("获取指定分数区间的成员列表：", byScore, err)

}
