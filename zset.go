package zdpgo_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// ZAdd 添加sorted set
func (r *Redis) ZAdd(key string, members ...*redis.Z) error {
	err := r.db.ZAdd(context.Background(), key, members...).Err()
	return err
}

func (r *Redis) ZCard(key string) (int64, error) {
	size, err := r.db.ZCard(context.Background(), key).Result()
	return size, err
}

// ZRange 取出sorted set
func (r *Redis) ZRange(key string, start int64, end int64) *redis.StringSliceCmd {
	return r.db.ZRange(context.Background(), key, start, end)
}
