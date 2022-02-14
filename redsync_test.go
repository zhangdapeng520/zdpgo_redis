package zdpgo_redis

import (
	"testing"
)

func TestRedis_NewRedSync(t *testing.T) {
	r := prepareRedis()

	mutex := r.NewRedSync("myglobalmutex")
	if err := mutex.Lock(); err != nil {
		panic(err)
	}

	// Do your work that requires the lock.
	t.Log("使用分布式锁锁住的逻辑")

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}
}
