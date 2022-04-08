package rstring

import (
	"time"
)

type Get interface {
	// Get 根据键获取值
	Get(key string) (value string, err error)
	// MGet 同时获取多个值
	MGet(keys ...string) ([]interface{}, error)
}

type Set interface {
	// Set 设置键值对
	Set(key string, value interface{}) error

	// SetExpire 设置键值对并指定过期时间
	SetExpire(key string, value interface{}, expire time.Duration) error

	// MSet 同时获取多个键对应的值
	MSet(kvs ...interface{}) error
}

type Incr interface {
	// Incr 自增1
	Incr(key string)

	// IncrBy 自增n
	IncrBy(key string, n int64)
}

type Decr interface {
	// Decr 自减1
	Decr(key string)

	// DecrBy 自减n
	DecrBy(key string, n int64)
}

type Delete interface {
	// Del 删除键
	Del(keys ...string) error
}

type Expire interface {
	// Expire 设置过期时间
	Expire(key string, expire time.Duration) error
}
