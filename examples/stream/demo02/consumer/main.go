package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_redis"
)

func main() {
	r := zdpgo_redis.New(zdpgo_redis.RedisConfig{
		Host:  "192.168.33.101",
		Port:  6379,
		Debug: true,
	})
	r.SubStream(zdpgo_redis.SubStreamConfig{
		Subject:           "test",
		ConsumerGroupName: "test_group",
		HandStreamFunc:    handleNewTicket,
	})
}

func handleNewTicket(values map[string]interface{}) error {
	fmt.Println("正在消费数据：", "values", values)
	return nil
}
