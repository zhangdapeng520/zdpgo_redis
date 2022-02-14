package zdpgo_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestRedis_ZRange(t *testing.T) {
	r := prepareRedis()
	client := r.db
	// 添加一个集合元素到集合中， 这个元素的分数是2.5，元素名是tizi
	err := client.ZAdd(context.Background(), "key", &redis.Z{2.5, "tizi"}).Err()
	if err != nil {
		panic(err)
	}

	size, err := client.ZCard(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(size)
}
