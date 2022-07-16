package rlist

import (
	"fmt"
	"testing"
)

func getList() *List {
	return NewList("10.1.3.52", 6379, "", "", 0, 20)
}
func TestRedis_LPush(t *testing.T) {
	r := getList()
	r.LPush("arr", 1, 2, 3)
	lLen, err := r.LLen("arr")
	if err != nil {
		t.Error("LPush Error", err)
		return
	}
	t.Log("arr length", lLen)
}

func TestRedis_LPushX(t *testing.T) {
	r := getList()
	r.LPushX("arr1", 1, 2, 3)
	lLen, _ := r.LLen("arr1")
	if lLen > 0 {
		t.Error("data not exists")
		return
	}
	t.Log("arr length", lLen)
}

func TestRedis_RPush(t *testing.T) {
	r := getList()
	r.RPush("arr", 1, 2, 3)
	lLen, err := r.LLen("arr")
	if err != nil {
		t.Error("RPush Error", err)
		return
	}
	t.Log("arr length", lLen)
}

func TestRedis_RPushX(t *testing.T) {
	r := getList()
	r.RPushX("arr3", 1, 2, 3)
	lLen, err := r.LLen("arr3")
	if err != nil {
		t.Error("RPush Error", err)
		return
	}
	if lLen > 0 {
		t.Error("data not exists")
	}
	t.Log("arr length", lLen)
}

func TestRedis_LRange(t *testing.T) {
	r := getList()
	key := "arr4"
	r.RPush(key, 1, 2, 3)
	lRange, err := r.LRange(key, 0, -1)
	if err != nil {
		t.Error(err)
	}
	t.Log(lRange)
}

func TestRedis_LIndex(t *testing.T) {
	r := getList()
	key := "arr5"
	r.RPush(key, 1, 2, 3)
	value, err := r.LIndex(key, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(value)
}

func TestRedis_LTrim(t *testing.T) {
	r := getList()
	key := "arr6"
	r.RPush(key, 1, 2, 3, 4, 5, 6)
	value, err := r.LTrim(key, 0, 3)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(value)

	lRange, err := r.LRange(key, 0, 6)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lRange)

	lLen, err := r.LLen(key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lLen)

	if lLen > (3 - 0 + 1) {
		t.Error("LTrim Error")
	}
}

func TestRedis_LSet(t *testing.T) {
	r := getList()
	key := "arr7"
	r.RPush(key, 1, 2, 3, 4, 5, 6)

	err := r.LSet(key, 1, "abc")
	if err != nil {
		t.Error(err)
	}

	index, err := r.LIndex(key, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(index)
}

func TestRedis_LInsertBefore(t *testing.T) {
	r := getList()
	key := "arr9"
	r.RPush(key, 1, 2, 3, 4, 5, 6)

	err := r.LInsertBefore(key, 3, "333")
	if err != nil {
		t.Error(err)
	}

	value, err := r.LIndex(key, 2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(value)
	if value != "333" {
		t.Error("insert error")
	}

	value, err = r.LIndex(key, 3)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(value)
	if value != "3" {
		t.Error("insert error")
	}
}

func TestRedis_LInsertAfter(t *testing.T) {
	r := getList()
	key := "arr9"
	r.RPush(key, 1, 2, 3, 4, 5, 6)

	err := r.LInsertAfter(key, 3, "333")
	if err != nil {
		t.Error(err)
	}

	value, err := r.LIndex(key, 3)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(value)
	if value != "333" {
		t.Error("insert error")
	}

	value, err = r.LIndex(key, 2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(value)
	if value != "3" {
		t.Error("insert error")
	}
}

func TestRedis_LPop(t *testing.T) {
	r := getList()
	key := "arr"
	r.RPush(key, "a", "b", "c")

	pop, err := r.LPop(key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(pop)

	if pop != "a" {
		t.Error("LPop error")
	}
}

func TestRedis_RPop(t *testing.T) {
	r := getList()
	key := "arr"
	r.RPush(key, "a", "b", "c")

	pop, err := r.RPop(key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(pop)

	if pop != "c" {
		t.Error("RPop error")
	}
}

func TestRedis_LRem(t *testing.T) {
	r := getList()
	key := "arr"
	r.RPush(key, "a", "b", "c", "c", "c", "c", "c", "c", "c")

	err := r.LRem(key, -7, "b")
	if err == nil {
		t.Error("count不能是负数")
	}

	err = r.LRem(key, 7, "b")
	if err != nil {
		t.Error(err)
	}

	err = r.LRem(key, 7, "c")
	if err != nil {
		t.Error(err)
	}

	lRange, err := r.LRange(key, 0, -1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lRange)
}
