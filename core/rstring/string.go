package rstring

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zhangdapeng520/zdpgo_redis/libs/redis"
)

type String struct {
	db *redis.Client // redis连接对象
}

func NewString(host string, port int, username, password string, db, poolSize int) *String {
	s := String{}
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

// Get 根据键获取值
func (r *String) Get(key string) (value string, err error) {
	value, err = r.db.Get(context.Background(), key).Result()
	info := ""
	switch {
	case err == redis.Nil:
		info = fmt.Sprintf("键【%s】不存在", key)
		err = errors.New(info)
	case err != nil:
		info = fmt.Sprintf("根据键获取值失败：%s", err.Error())
		err = errors.New(info)
	case value == "":
		info = "值不存在"
	}
	return
}

// SetExpire 设置值，自定义过期时间
func (r *String) SetExpire(key string, value interface{}, expire time.Duration) error {
	err := r.db.Set(context.Background(), key, value, expire).Err()
	return err
}

// Set 根据键设置值，过期时间默认为30天
func (r *String) Set(key string, value interface{}) {
	r.SetExpire(key, value, 30*24*60*60*time.Second)
}

// MGet 同时获取多个键对应的值
func (r *String) MGet(keys ...string) ([]interface{}, error) {
	result, err := r.db.MGet(context.Background(), keys...).Result()
	return result, err
}

// MSet 同时获取多个键对应的值
func (r *String) MSet(kvs ...interface{}) error {
	_, err := r.db.MSet(context.Background(), kvs...).Result()
	return err
}

// Incr 自增1
func (r *String) Incr(key string) {
	r.db.Incr(context.Background(), key)
}

// IncrBy 自增n
func (r *String) IncrBy(key string, n int64) {
	r.db.IncrBy(context.Background(), key, n)
}

// Decr 自减1
func (r *String) Decr(key string) {
	r.db.Decr(context.Background(), key)
}

// DecrBy 自减n
func (r *String) DecrBy(key string, n int64) {
	r.db.DecrBy(context.Background(), key, n)
}

// Append 追加字符串
func (r *String) Append(key string, value string) {
	r.db.Append(context.Background(), key, value)
}

// Del 删除键
func (r *String) Del(keys ...string) error {
	err := r.db.Del(context.Background(), keys...).Err()
	return err
}

// Expire 设置过期时间·
func (r *String) Expire(key string, expire time.Duration) error {
	err := r.db.Expire(context.Background(), key, expire).Err()
	return err
}

// Substr 截取字符串
func (r *String) Substr(key string, start, end int) (string, error) {
	result, err := r.db.Do(context.Background(), "substr", key, start, end-1).Result()
	return result.(string), err
}
