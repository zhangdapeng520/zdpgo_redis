package zdpgo_redis

import (
	"fmt"
	"testing"
)

func prepareRedis() *Redis {
	r := New(RedisConfig{
		Host: "192.168.33.101",
	})
	return r
}

// 测试新建redis
func TestRedis_New(t *testing.T) {
	r := prepareRedis()
	fmt.Println(r.Status())
}
