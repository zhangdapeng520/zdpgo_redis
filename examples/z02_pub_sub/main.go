package main

import (
	"fmt"
	"time"

	"github.com/zhangdapeng520/zdpgo_redis"
)

func main() {
	config := zdpgo_redis.RedisConfig{
		Host: "192.168.18.101",
		Port: 6379,
	}
	r := zdpgo_redis.New(config)
	r.SetDebug(true)

	// 发布消息
	go func() {
		v := 0
		for {
			r.Publish("test", fmt.Sprintf("hello: %d", v))
			v += 1
			time.Sleep(time.Second)
		}
	}()

	// 订阅消息
	go func() {
		ps := r.Subscribe("test")
		for {
			msg, _ := r.ReceiveMessage(ps)
			fmt.Println(msg.Channel, msg.Payload)
		}
	}()

	time.Sleep(60 * time.Second)
}
