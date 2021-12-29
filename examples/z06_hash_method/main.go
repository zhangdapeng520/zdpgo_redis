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

	var (
		nums    int64
		results []string
		data    map[string]string
	)

	// 添加元素
	r.HSet("zhangdapeng", "name", "张大鹏")
	r.HSet("zhangdapeng", "age", 22)
	r.HSet("zhangdapeng", "gender", true)

	// 获取元素
	result, _ := r.HGet("zhangdapeng", "name")
	fmt.Println(result)

	// 批量获取元素
	result1, _ := r.HMGet("zhangdapeng", "name", "age", "gender")
	fmt.Println(result1)

	// 批量设置元素
	r.HMSet("lisi", "name", "李四", "age", 22, "gender", true)

	// 年龄自增
	r.HIncr("lisi", "age")
	r.HIncrBy("lisi", "age", 11)
	result1, _ = r.HMGet("lisi", "name", "age", "gender")
	fmt.Println(result1)

	// 判断字段是否存在
	result2, _ := r.HExists("lisi", "age")
	fmt.Println(result2)
	nums, _ = r.HLen("lisi")
	fmt.Println(nums)     // 字段数
	r.HDel("lisi", "age") // 删除age字段
	result2, _ = r.HExists("lisi", "age")
	fmt.Println(result2)
	nums, _ = r.HLen("lisi")
	fmt.Println(nums) // 字段数

	// 获取所有的键
	results, _ = r.HKeys("lisi")
	fmt.Println(results)

	// 获取所有值
	results, _ = r.HVals("lisi")
	fmt.Println(results)

	// 获取所有键值对
	data, _ = r.HGetAll("lisi")
	fmt.Println(data)
}
