package main

import (
	"github.com/zhangdapeng520/zdpgo_redis"
	"math/rand"
)

func main() {
	r := zdpgo_redis.New(zdpgo_redis.RedisConfig{
		Host:  "192.168.33.101",
		Port:  6379,
		Debug: true,
	})

	for i := 0; i < 3000; i++ {
		values := map[string]interface{}{
			"whatHappened": string("ticket received"),
			"ticketID":     int(rand.Intn(100000000)),
			"ticketData":   string("some ticket data"),
		}
		r.PubStream(zdpgo_redis.PubStreamConfig{
			Subject: "test",
			Values:  values,
		})
	}

}
