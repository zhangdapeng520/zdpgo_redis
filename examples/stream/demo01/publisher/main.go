package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/go-redis/redis"
)

func main() {
	log.Println("Publisher started")

	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "192.168.33.101", "6379"),
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Unbale to connect to Redis", err)
	}

	log.Println("Connected to Redis server")

	for i := 0; i < 3000; i++ {
		err = publishTicketReceivedEvent(redisClient)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func publishTicketReceivedEvent(client *redis.Client) error {
	log.Println("Publishing event to Redis")

	// 简单地使用XADD命令向Redis流发送一些消息
	err := client.XAdd(&redis.XAddArgs{
		Stream:       "tickets",
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values: map[string]interface{}{
			"whatHappened": string("ticket received"),
			"ticketID":     int(rand.Intn(100000000)),
			"ticketData":   string("some ticket data"),
		},
	}).Err()

	return err
}
