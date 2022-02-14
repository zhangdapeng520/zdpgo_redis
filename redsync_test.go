package zdpgo_redis

import (
	"testing"
)

func TestRedis_NewRedSync(t *testing.T) {
	r := New(RedisConfig{
		Host:  "192.168.33.101",
		Port:  6379,
		Debug: true,
	})

	mutex := r.NewRedSync("myglobalmutex")
	if err := mutex.Lock(); err != nil {
		panic(err)
	}

	// Do your work that requires the lock.

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}
}
