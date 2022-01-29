package zdpgo_redis

import (
	"fmt"
	"testing"
)

func TestRedis_SetGet(t *testing.T) {
	r := prepareRedis()
	r.Set("username", "张大鹏")
	fmt.Println(r.Get("username"))
}

func TestRedis_MGet(t *testing.T) {
	r := prepareRedis()
	r.Set("username", "张大鹏")
	r.Set("age", 22)
	r.Set("gender", "male")
	fmt.Println(r.MGet("username", "age", "gender"))
}

func TestRedis_AddSub(t *testing.T) {
	r := prepareRedis()
	r.Set("age", 22)
	fmt.Println(r.Get("age"))
	r.Add1("age")
	fmt.Println(r.Get("age"))
	r.AddN("age", 3)
	fmt.Println(r.Get("age"))
	r.Sub1("age")
	fmt.Println(r.Get("age"))
	r.SubN("age", 3)
	fmt.Println(r.Get("age"))
}

func TestRedis_AppendSubstr(t *testing.T) {
	r := prepareRedis()
	r.Set("test", "a")
	r.Append("test", "b")
	r.Append("test", "c")
	r.Append("test", "d")
	r.Append("test", "e")
	fmt.Println(r.Get("test"))
	fmt.Println(r.Substr("test", 1, 3))
	fmt.Println(r.Substr("test", 0, 3))
	fmt.Println(r.Substr("test", 0, -1))
	fmt.Println(r.Substr("test", 0, 33))
	fmt.Println(r.Substr("test", 0, -33))
}
