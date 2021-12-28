package zdpgo_redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// 发布消息
func (r *Redis) Publish(key, value string) error {
	err := r.db.Publish(context.Background(), key, value).Err()
	return err
}

// 订阅消息
func (r *Redis) Subscribe(key string) *redis.PubSub {
	ps := r.db.Subscribe(context.Background(), key)
	return ps
}

// 获取消息
func (r *Redis) ReceiveMessage(ps *redis.PubSub) (*redis.Message, error) {
	msg, err := ps.ReceiveMessage(context.Background())
	if err != nil {
		r.log.Error("获取消息失败：", err)
	}
	// fmt.Println(msg.Channel, msg.Payload)
	return msg, err
}

// 关闭订阅
func (r *Redis) CloseSubscribe(ps *redis.PubSub) {
	err := ps.Close()
	if err != nil {
		r.log.Error("关闭订阅失败：", err)
		return
	}
}
