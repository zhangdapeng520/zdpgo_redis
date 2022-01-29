package zdpgo_redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Get 根据键获取值
func (r *Redis) Get(key string) (value string, err error) {
	value, err = r.db.Get(context.Background(), key).Result()
	info := ""
	switch {
	case err == redis.Nil:
		info = fmt.Sprintf("键【%s】不存在", key)
		err = errors.New(info)
		r.log.Error(info)
	case err != nil:
		info = fmt.Sprintf("根据键获取值失败：%s", err.Error())
		err = errors.New(info)
		r.log.Error(info)
	case value == "":
		info = "值不存在"
		r.log.Warning(info)
	}
	return
}

// SetExpire 设置值，自定义过期时间
func (r *Redis) SetExpire(key string, value interface{}, expire time.Duration) {
	err := r.db.Set(context.Background(), key, value, expire).Err()
	if err != nil {
		r.log.Error("根据键设置值失败：", err)
	}
}

// Set 根据键设置值，过期时间默认为30天
func (r *Redis) Set(key string, value interface{}) {
	r.SetExpire(key, value, 30*24*60*60*time.Second)
}

// MGet 同时获取多个键对应的值
func (r *Redis) MGet(keys ...string) ([]interface{}, error) {
	result, err := r.db.MGet(context.Background(), keys...).Result()
	if err != nil {
		r.log.Error("根据多个键同时获取值失败：", err)
	}
	return result, err
}

// Add1 自增1
func (r *Redis) Add1(key string) {
	r.db.Incr(context.Background(), key)
}

// AddN 自增n
func (r *Redis) AddN(key string, n int64) {
	r.db.IncrBy(context.Background(), key, n)
}

// Sub1 自减1
func (r *Redis) Sub1(key string) {
	r.db.Decr(context.Background(), key)
}

// SubN 自减n
func (r *Redis) SubN(key string, n int64) {
	r.db.DecrBy(context.Background(), key, n)
}

// Append 追加字符串
func (r *Redis) Append(key string, value string) {
	r.db.Append(context.Background(), key, value)
}

// Substr 截取字符串
func (r *Redis) Substr(key string, start, end int) (string, error) {
	result, err := r.db.Do(context.Background(), "substr", key, start, end-1).Result()
	if err != nil {
		r.log.Error("截取字符串失败：", err)
	}
	return result.(string), err
}
