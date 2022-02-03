package zdpgo_redis

import "github.com/go-redsync/redsync/v4"
import "github.com/go-redsync/redsync/v4/redis/goredis/v8"

// NewRedSync 创建redis分布式锁对象
func (r *Redis) NewRedSync(mutexName string) *redsync.Mutex {
	pool := goredis.NewPool(r.db)
	r.redSync = redsync.New(pool)
	mutex := r.redSync.NewMutex(mutexName)
	r.redSyncMutex = mutex
	return mutex
}
