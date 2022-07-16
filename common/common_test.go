package common

import "testing"

func getCommon() *Common {
	return NewCommon("127.0.0.1", 6379, "", "", 0, 20)
}

// 测试获取数据库中所有的key
func TestCommon_Keys(t *testing.T) {
	c := getCommon()
	keys, err := c.Keys()
	if err != nil {
		t.Error(err)
	}
	t.Log(keys)
}
