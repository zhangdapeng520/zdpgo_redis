package zdpgo_redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// 根据键获取值
func (r *Redis) Get(key string) (value string, err error) {
	value, err = r.db.Get(context.Background(), "key").Result()
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
