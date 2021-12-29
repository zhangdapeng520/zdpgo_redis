package zdpgo_redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// 根据键获取值
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

// 设置值，自定义过期时间
func (r *Redis) SetExpire(key string, value interface{}, expire time.Duration) {
	err := r.db.Set(context.Background(), key, value, expire).Err()
	if err != nil {
		r.log.Error("根据键设置值失败：", err)
	}
}

// 根据键设置值，过期时间默认为3小时
func (r *Redis) Set(key string, value interface{}) {
	r.SetExpire(key, value, 3*60*60*time.Second)
}

// 同时获取多个键对应的值
func (r *Redis) MGet(keys ...string) ([]interface{}, error) {
	result, err := r.db.MGet(context.Background(), keys...).Result()
	if err != nil {
		r.log.Error("根据多个键同时获取值失败：", err)
	}
	return result, err
}

// 自增1
func (r *Redis) Incr(key string) {
	r.db.Incr(context.Background(), key)
}

// 自增n
func (r *Redis) IncrBy(key string, n int64) {
	r.db.IncrBy(context.Background(), key, n)
}

// 自减1
func (r *Redis) Decr(key string) {
	r.db.Decr(context.Background(), key)
}

// 自减n
func (r *Redis) DecrBy(key string, n int64) {
	r.db.DecrBy(context.Background(), key, n)
}

// 追加字符串
func (r *Redis) Append(key string, value string) {
	r.db.Append(context.Background(), key, value)
}

// 截取字符串
func (r *Redis) Substr(key string, start, end int) (string, error) {
	result, err := r.db.Do(context.Background(), "substr", key, start, end-1).Result()
	if err != nil {
		r.log.Error("截取字符串失败：", err)
	}
	return result.(string), err
}
