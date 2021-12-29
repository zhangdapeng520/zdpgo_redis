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

	// 从右边追加值
	// r.RPush("list", 1, 2, 3)

	// 从左边追加值
	// r.LPush("list", 3, 2, 1)

	// 获取列表长度
	result1, err := r.LLen("list")
	fmt.Println(result1, err)

	// 获取列表元素
	result3, err := r.LRange("list", 0, 6)
	fmt.Println(result3)

	// 截取列表内容
	r.LTrim("list", 0, 6)
	r.LTrim("list", 0, 7)

	result4, err := r.LRange("list", 0, 7)
	fmt.Println(result4, err)

	// 设置指定位置的元素
	r.LSet("list", 3, 333)

	// 查看指定索引位置的元素
	result2, _ := r.LIndex("list", 3)
	fmt.Println(result2)
}
