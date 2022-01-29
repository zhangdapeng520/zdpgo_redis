package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/rs/xid"
)

func main() {
	log.Println("Consumer started")

	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "192.168.33.101", "6379"),
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Unbale to connect to Redis", err)
	}

	log.Println("Connected to Redis server")

	subject := "tickets"
	consumersGroup := "tickets-consumer-group"

	// 使用XGROUPCREATE来创建消费者组
	err = redisClient.XGroupCreate(subject, consumersGroup, "0").Err()
	if err != nil {
		log.Println(err)
	}

	// 现在可以使用XREADGROUP来监听流中消息，并使用一个唯一id将消费者注册到消费者组里：
	// 为了生成唯一id，将使用xid库
	uniqueID := xid.New().String()

	for {
		entries, err := redisClient.XReadGroup(&redis.XReadGroupArgs{
			Group:    consumersGroup,
			Consumer: uniqueID,
			Streams:  []string{subject, ">"},
			Count:    2,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			log.Fatal(err)
		}

		// 然后在main.go中创建一个无限循环，我们调用XREADGROUP并在>位置，表示从该组的第一个待处理消息开始
		// 然后为每个ticket调用handNewTicket函数，并发送XACK命令到redis服务通知消息已经被消费。
		for i := 0; i < len(entries[0].Messages); i++ {
			messageID := entries[0].Messages[i].ID
			values := entries[0].Messages[i].Values
			eventDescription := fmt.Sprintf("%v", values["whatHappened"])
			ticketID := fmt.Sprintf("%v", values["ticketID"])
			ticketData := fmt.Sprintf("%v", values["ticketData"])

			if eventDescription == "ticket received" {
				err := handleNewTicket(ticketID, ticketData)
				if err != nil {
					log.Fatal(err)
				}
				redisClient.XAck(subject, consumersGroup, messageID)
			}
		}
	}
}

func handleNewTicket(ticketID string, ticketData string) error {
	log.Printf("Handling new ticket id : %s data %s\n", ticketID, ticketData)
	// time.Sleep(100 * time.Millisecond)
	return nil
}
