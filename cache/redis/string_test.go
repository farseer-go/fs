package redis

import (
	"fmt"
	"testing"
	"time"
)

//String 测试
func TestClientString(t *testing.T) {
	client := NewClient("default")
	err := client.String.Set("key1", "...3456")
	if err == nil {
		fmt.Printf("设置值:%v\n", "...3456")
	}
	get, _ := client.String.Get("key1")
	fmt.Printf("获取值：%v\n", get)

	//如果key值存在，设置这个会返回false
	nx, _ := client.String.SetNX("key2", "1231", 100*time.Second)
	fmt.Printf("设置过期时间：%v\n", nx)

	get2, _ := client.String.Get("key2")
	fmt.Printf("获取值：%v\n", get2)

	ttl, _ := client.Key.TTL("key2")
	fmt.Printf("获取过期时间：%v\n", ttl)
}
