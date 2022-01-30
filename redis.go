package zdpgo_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zhangdapeng520/zdpgo_random"
	"github.com/zhangdapeng520/zdpgo_zap"
	"sync"
)

// Redis 操作redis的核心对象
type Redis struct {
	db     *redis.Client        // redis连接对象
	log    *zdpgo_zap.Zap       // 日志对象
	config *RedisConfig         // 配置对象
	random *zdpgo_random.Random // 生成随机数据的核心对象
	lock   sync.Mutex           // 互斥锁对象
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port), // 连接地址
		Password: config.Password,                                // 密码
		DB:       config.Database,                                // 数据库
		PoolSize: config.PoolSize,                                // 连接池中的连接个数
	})
	r.db = rdb

	// 初始化日志
	go func(r *Redis) {
		r.lock.Lock()
		r.log = zdpgo_zap.New(zdpgo_zap.ZapConfig{
			Debug:        config.Debug,       // 是否为debug模式
			OpenGlobal:   true,               // 是否开启全局日志
			OpenFileName: true,               // 是否输出文件名和行号
			LogFilePath:  config.LogFilePath, // 日志路径
		})
		r.lock.Unlock()
	}(&r)

	// 初始化随机数
	go func(r *Redis) {
		r.lock.Lock()
		r.random = zdpgo_random.New(zdpgo_random.RandomConfig{
			Debug:       r.config.Debug,
			LogFilePath: r.config.LogFilePath,
		})
		r.lock.Unlock()
	}(&r)

	return &r
}

// SetDebug 设置debug模式
func (r *Redis) SetDebug(debug bool) {
	r.config.Debug = debug
	r.log = zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        r.config.Debug,       // 是否为debug模式
		OpenGlobal:   true,                 // 是否开启全局日志
		OpenFileName: true,                 // 是否输出文件名和行号
		LogFilePath:  r.config.LogFilePath, // 日志路径
	})
}

// IsDebug 是否为debug模式
func (r *Redis) IsDebug() bool {
	return r.config.Debug
}

// Status 查看Redis服务状态
func (r *Redis) Status() bool {
	pong, err := r.db.Ping(context.Background()).Result()
	if err != nil {
		r.log.Error("redis连接失败：", "ping", pong, "err", err)
		return false
	}
	return true
}
