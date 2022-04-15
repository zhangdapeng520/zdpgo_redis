package common

import (
	"context"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_redis/libs/redis"
	"time"
)

type Common struct {
	db *redis.Client // redis连接对象
}

func NewCommon(host string, port int, username, password string, db, poolSize int) *Common {
	s := Common{}
	address := fmt.Sprintf("%s:%d", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,  // 连接地址
		Username: username, // 用户名
		Password: password, // 密码
		DB:       db,       // 数据库
		PoolSize: poolSize, // 连接池中的连接个数
	})
	s.db = rdb
	return &s
}

// Del 删除键
func (r *Common) Del(keys ...string) error {
	err := r.db.Del(context.Background(), keys...).Err()
	return err
}

// Expire 设置过期时间·
func (r *Common) Expire(key string, expire time.Duration) error {
	err := r.db.Expire(context.Background(), key, expire).Err()
	return err
}

// Keys 获取数据库中所有的keys
func (r *Common) Keys(pattern ...string) ([]string, error) {
	// 默认匹配所有
	p := "*"
	if len(pattern) > 0 {
		p = pattern[0]
	}

	// 执行查询
	keys := r.db.Keys(context.Background(), p)

	// 提取查询结果
	result, err := keys.Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
