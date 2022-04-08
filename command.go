package zdpgo_redis

import (
	"context"
)

// Execute 执行Redis命令
func (r *Redis) Execute(args ...interface{}) (string, error) {
	val, err := r.db.Do(context.Background(), args...).Text()
	return val, err
}
