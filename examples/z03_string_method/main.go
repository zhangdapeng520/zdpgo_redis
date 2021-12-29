package main

import (
	"fmt"
	"reflect"
	"strconv"

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
	r.Set("age", 22)
	r.Set("gender", true)

	// 自增1
	r.Incr("age")

	// 自增n
	r.IncrBy("age", 10)

	// 自减1
	r.Decr("age")

	// 自减n
	r.DecrBy("age", 2)

	// 追加
	r.Append("name", "1")
	r.Append("name", "2")
	r.Append("name", "3")

	// 截取字符串
	substr, _ := r.Substr("name", 0, 9) // 一个中文字符占3个长度
	fmt.Println(substr)

	// 获取值
	result, err := r.MGet("name", "age", "gender")
	fmt.Println(result, err)
	name := result[0].(string)
	age, _ := strconv.Atoi(result[1].(string))
	gender, _ := strconv.ParseBool(result[2].(string))
	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(age, reflect.TypeOf(age))
	fmt.Println(gender, reflect.TypeOf(gender))
}
