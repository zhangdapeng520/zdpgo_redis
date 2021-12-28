package zdpgo_redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/zhangdapeng520/zdpgo_log"
)

// 操作redis的核心对象
type Redis struct {
	db          *redis.Client  // redis连接对象
	log         *zdpgo_log.Log // 日志对象
	logFilePath string         // 日志路径
	debug       bool           // 是否为debug模式
}

// redis配置对象
type RedisConfig struct {
	Host        string // 主机地址
	Port        int    // 端口号
	Database    int    // 数据库
	Username    string // 用户名
	Password    string // 密码
	LogFilePath string // 日志路径
	Debug       bool   // 是否为debug模式
}

// 创建Redis操作对象
func New(config RedisConfig) *Redis {
	r := Redis{}

	// 初始化日志
	if config.LogFilePath != "" {
		r.logFilePath = config.LogFilePath
		r.log = zdpgo_log.New(config.LogFilePath)
	} else {
		r.logFilePath = "zdpgo_redis.log"
		r.log = zdpgo_log.New(r.logFilePath)
	}

	// 初始化debug模式
	r.log.SetDebug(config.Debug)
	r.debug = config.Debug

	// 初始化Redis连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port), // 连接地址
		Password: config.Password,                                // 密码
		DB:       config.Database,                                // 数据库
	})
	r.db = rdb

	return &r
}

// 设置debug模式
func (r *Redis) SetDebug(debug bool) {
	r.debug = debug
	r.log.SetDebug(debug)
}

// 是否为debug模式
func (r *Redis) IsDebug() bool {
	return r.debug
}
