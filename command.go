package zdpgo_redis

import (
	"context"
)

// 执行Redis命令
func (r *Redis) Execute(args ...interface{}) (string, error) {
	val, err := r.db.Do(context.Background(), args...).Text()
	r.log.Info(val, err)
	return val, err
}
