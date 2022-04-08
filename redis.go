package zdpgo_redis

import (
	"context"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_redis/libs/redis"
	"sync"
)

// Redis 操作redis的核心对象
type Redis struct {
	db     *redis.Client // redis连接对象
	config *RedisConfig  // 配置对象
	lock   sync.Mutex    // 互斥锁对象
}

// New 创建Redis操作对象
func New(config RedisConfig) *Redis {
	r := Redis{}

	// 初始化配置
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_redis.log"
	}
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		config.Port = 6379
	}
	if config.StreamTag == "" {
		config.StreamTag = "zdpgo_redis_config_stream_tag"
	}
	r.config = &config

	// 初始化Redis连接
	if config.PoolSize == 0 {
		config.PoolSize = 33 // 默认是33个
	}
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,         // 连接地址
		Password: config.Password, // 密码
		DB:       config.Database, // 数据库
		PoolSize: config.PoolSize, // 连接池中的连接个数
	})
	r.db = rdb

	return &r
}

// Status 查看Redis服务状态
func (r *Redis) Status() bool {
	_, err := r.db.Ping(context.Background()).Result()
	if err != nil {
		return false
	}
	return true
}
