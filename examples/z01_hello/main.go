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

	// 设置值
	r.Set("name", "张大鹏")

	// 获取值
	value, err := r.Get("name")
	fmt.Println(value, err)
}
