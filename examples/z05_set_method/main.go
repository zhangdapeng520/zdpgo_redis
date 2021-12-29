package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_redis"
)

func main() {
	// 创建Redis对象
	config := zdpgo_redis.RedisConfig{
		Host: "192.168.18.101",
		Port: 6379,
	}
	r := zdpgo_redis.New(config)
	r.SetDebug(true)

	// 查看连接状态
	fmt.Println(r.Status())

	// 添加元素
	r.SAdd("set", "a", "b", "c")

	// 删除元素
	r.SRem("set", "a")

	// 获取所有元素
	result, _ := r.SMembers("set")
	fmt.Println(result)

	// 求交集
	r.SAdd("set1", "c", "d", "e")
	result, _ = r.SInter("set", "set1")
	fmt.Println(result)

	// 求并集
	result, _ = r.SUnion("set", "set1")
	fmt.Println(result)

	// 求差集
	result, _ = r.SDiff("set", "set1")
	fmt.Println(result)
}
