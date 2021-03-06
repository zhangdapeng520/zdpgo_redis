package zdpgo_redis

import (
	"testing"
)

func prepareRedis() *Redis {
	r := New(Config{
		Host: "10.1.3.12",
	})
	if r.Status() {
		return r
	}
	return nil
}

// 测试新建redis
func TestRedis_New(t *testing.T) {
	r := prepareRedis()
	t.Log(r.Status())
}
