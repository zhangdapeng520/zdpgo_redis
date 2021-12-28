package main

import (
	"github.com/zhangdapeng520/zdpgo_redis"
)

func main() {
	config := zdpgo_redis.RedisConfig{
		Host: "192.168.18.101",
		Port: 6379,
	}
	r := zdpgo_redis.New(config)
	r.SetDebug(true)
}
