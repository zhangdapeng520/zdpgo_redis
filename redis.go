package zdpgo_redis

import (
	"context"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_redis/core/common"
	"github.com/zhangdapeng520/zdpgo_redis/core/rhash"
	"github.com/zhangdapeng520/zdpgo_redis/core/rlist"
	"github.com/zhangdapeng520/zdpgo_redis/core/rset"
	"github.com/zhangdapeng520/zdpgo_redis/core/rstring"
	"github.com/zhangdapeng520/zdpgo_redis/libs/redis"
	"sync"
)

// Redis 操作redis的核心对象
type Redis struct {
	db     *redis.Client   // redis连接对象
	config *Config         // 配置对象
	lock   sync.Mutex      // 互斥锁对象
	Common *common.Common  // Redis通用操作对象
	String *rstring.String // 操作字符串的核心对象
	List   *rlist.List     // 操作列表的核心对象
	Hash   *rhash.Hash     // 操作hash的核心对象
	Set    *rset.Set       // 操作set的核心对象
}

// New 创建Redis操作对象
func New(config Config) *Redis {
	r := Redis{}

	// 获取默认配置
	config = getDefaultConfig(config)

	// 连接Redis
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,         // 连接地址
		Password: config.Password, // 密码
		DB:       config.Database, // 数据库
		PoolSize: config.PoolSize, // 连接池中的连接个数
	})
	r.db = rdb

	// 实例化操作对象
	r.Common = common.NewCommon(
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
		config.PoolSize,
	)
	r.String = rstring.NewString(
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
		config.PoolSize,
	)
	r.List = rlist.NewList(
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
		config.PoolSize,
	)
	r.Hash = rhash.NewHash(
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
		config.PoolSize,
	)
	r.Set = rset.NewSet(
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
		config.PoolSize,
	)

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
