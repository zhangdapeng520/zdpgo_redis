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
func (r *Redis) SetExpire(key, value string, expire time.Duration) {
	err := r.db.Set(context.Background(), key, value, expire).Err()
	if err != nil {
		r.log.Error("根据键设置值失败：", err)
	}
}

// 根据键设置值，过期时间默认为3小时
func (r *Redis) Set(key, value string) {
	r.SetExpire(key, value, 3*60*60*time.Second)
}
