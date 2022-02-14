package zdpgo_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// HandStreamFunc 处理流中数据的方法
type HandStreamFunc func(map[string]interface{}) error

// PubStream 发布流
// @param config 流配置
func (r *Redis) PubStream(config PubStreamConfig) error {
	r.log.Info("PubStream 正在发布消息", "config", config)

	// 添加唯一标识符
	config.Values["zdpgo_redis_stream_tag"] = r.config.StreamTag

	// 简单地使用XADD命令向Redis流发送一些消息
	err := r.db.XAdd(context.Background(), &redis.XAddArgs{
		Stream: config.Subject,
		MaxLen: config.MaxLen,
		ID:     config.ID,
		Values: config.Values,
	}).Err()

	return err
}

// SubStream 订阅流
func (r *Redis) SubStream(config SubStreamConfig) (err error) {
	// 边界条件
	if config.HandStreamFunc == nil {
		r.log.Panic("消费Redis Stream失败，处理流的函数不能为空")
	}

	// 使用XGROUPCREATE来创建消费者组
	err = r.db.XGroupCreate(context.Background(), config.Subject, config.ConsumerGroupName, "0").Err()
	if err != nil {
		r.log.Error("r.db.XGroupCreate 创建消费者组失败", "error", err)
	}

	// 现在可以使用XREADGROUP来监听流中消息，并使用一个唯一id将消费者注册到消费者组里：
	// 为了生成唯一id，将使用xid库
	uniqueID := r.random.Xid()
	r.log.Info("生成唯一ID", "id", uniqueID)

	var entries []redis.XStream
	for {
		r.log.Info("r.db.XReadGroup 读取数据")
		entries, err = r.db.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Group:    config.ConsumerGroupName,
			Consumer: uniqueID,
			Streams:  []string{config.Subject, ">"},
			Count:    2,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			r.log.Fatal("r.db.XReadGroup 读取组数据失败", "error", err)
		}

		// 然后在main.go中创建一个无限循环，我们调用XREADGROUP并在>位置，表示从该组的第一个待处理消息开始
		// 然后为每个ticket调用handNewTicket函数，并发送XACK命令到redis服务通知消息已经被消费。
		// 循环消费组数据
		for i := 0; i < len(entries[0].Messages); i++ {
			// 获取消息id和values数据
			messageID := entries[0].Messages[i].ID
			values := entries[0].Messages[i].Values

			// 获取唯一标识符
			tag := fmt.Sprintf("%v", values["zdpgo_redis_stream_tag"])
			r.log.Info("获取唯一标识符：", "tag", tag, "configStreamTag", r.config.StreamTag)

			// 消费数据
			if tag == r.config.StreamTag {
				r.log.Info("config.HandStreamFunc 消费数据", "values", values)
				err = config.HandStreamFunc(values)
				if err != nil {
					r.log.Fatal("config.HandStreamFunc 消费数据失败", "error", err)
				}
				r.db.XAck(context.Background(), config.Subject, config.ConsumerGroupName, messageID)
				r.log.Info("r.db.XAck 通知Redis指定数据已被消费")
			}
		}
	}
}
