package rhash

import (
	"testing"
)

func getHash() *Hash {
	return NewHash("10.1.3.52", 6379, "", "", 0, 20)
}
func TestRedis_HGet(t *testing.T) {
	r := getHash()

	r.HSet("user_2", "username", "张大鹏")
	get, err := r.HGet("user_2", "username")
	if err != nil {
		t.Error(err)
	}
	t.Log(get, err)
}

func TestRedis_HMSet(t *testing.T) {
	r := getHash()

	r.HMSet("user_2", "username", "张大鹏", "age", 22)
	get, err := r.HMGet("user_2", "username", "age")
	if err != nil {
		t.Error(err)
	}
	t.Log(get, err)
}

func TestRedis_HIncr(t *testing.T) {
	r := getHash()

	r.HMSet("user_2", "username", "张大鹏", "age", 22)
	err := r.HIncr("user_2", "age")
	if err != nil {
		t.Error(err)
	}

	mGet, err := r.HGet("user_2", "age")
	if err != nil {
		t.Error(err)
	}
	t.Log(mGet)
}

func TestRedis_HDel(t *testing.T) {
	r := getHash()

	key := "user_2"
	r.HMSet(key, "username", "张大鹏", "age", 22)
	err := r.HIncr(key, "age")
	if err != nil {
		t.Error(err)
	}

	err = r.HDel(key, "username")
	if err != nil {
		t.Error(err)
	}

	mGet, err := r.HGet(key, "username")
	if err == nil {
		t.Error(err)
	}
	t.Log(mGet, err)
}

func TestRedis_HLen(t *testing.T) {
	r := getHash()

	key := "user_2"
	r.HMSet(key, "username", "张大鹏", "age", 22)
	hLen, err := r.HLen(key)
	if err != nil {
		t.Error(err)
	}
	t.Log(hLen)
}

func TestRedis_HKeys(t *testing.T) {
	r := getHash()

	key := "user_2"
	r.HMSet(key, "username", "张大鹏", "age", 22)
	hLen, err := r.HKeys(key)
	if err != nil {
		t.Error(err)
	}
	t.Log(hLen)
}

func TestRedis_HVals(t *testing.T) {
	r := getHash()

	key := "user_2"
	r.HMSet(key, "username", "张大鹏", "age", 22)
	hLen, err := r.HVals(key)
	if err != nil {
		t.Error(err)
	}
	t.Log(hLen)
}

func TestRedis_HGetAll(t *testing.T) {
	r := getHash()

	key := "user_2"
	r.HMSet(key, "username", "张大鹏", "age", 22)
	hLen, err := r.HGetAll(key)
	if err != nil {
		t.Error(err)
	}
	t.Log(hLen)
}

func TestRedis_HSetNX(t *testing.T) {
	r := getHash()

	key := "user_2"
	r.HSetNX(key, "username", "张大鹏")
	r.HSetNX(key, "username", "张大鹏111")

	get, err := r.HGet(key, "username")
	if err != nil {
		t.Error(err)
	}
	t.Log(get)
}

func TestRedis_HIncrByFloat(t *testing.T) {
	r := getHash()

	key := "user_2"
	r.HSetNX(key, "height", 1.65)
	r.HIncrByFloat(key, "height", 0.23)

	get, err := r.HGet(key, "height")
	if err != nil {
		t.Error(err)
	}
	t.Log(get)
}
