package rlist

import (
	"context"
	"errors"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_redis/libs/redis"
)

type List struct {
	db *redis.Client // redis连接对象
}

func NewList(host string, port int, username, password string, db, poolSize int) *List {
	s := List{}
	address := fmt.Sprintf("%s:%d", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,  // 连接地址
		Username: username, // 用户名
		Password: password, // 密码
		DB:       db,       // 数据库
		PoolSize: poolSize, // 连接池中的连接个数
	})
	s.db = rdb
	return &s
}

// RPush 在名称为key的list尾添加一个值为value的元素
// 从右边追加元素
func (r *List) RPush(key string, values ...interface{}) error {
	err := r.db.RPush(context.Background(), key, values...).Err()
	return err
}

// RPushX 仅当key存在的时候，从右边追加数据
func (r *List) RPushX(key string, values ...interface{}) error {
	err := r.db.RPushX(context.Background(), key, values...).Err()
	return err
}

// LPush 在名称为key的list头添加一个值为value的 元素
// 从左边追加元素
func (r *List) LPush(key string, values ...interface{}) error {
	err := r.db.LPush(context.Background(), key, values...).Err()
	return err
}

// LPushX 仅当数据存在的时候才能插入
func (r *List) LPushX(key string, values ...interface{}) error {
	err := r.db.LPushX(context.Background(), key, values...).Err()
	return err
}

// LLen 返回名称为key的list的长度
func (r *List) LLen(key string) (int64, error) {
	result, err := r.db.LLen(context.Background(), key).Result()
	return result, err
}

// LRange 返回名称为key的list中start至end之间的元素
func (r *List) LRange(key string, start, stop int64) ([]string, error) {
	result, err := r.db.LRange(context.Background(), key, start, stop).Result()
	return result, err
}

// LTrim 截取名称为key的list，只保留start到top之间的数据
func (r *List) LTrim(key string, start, stop int64) (string, error) {
	result, err := r.db.LTrim(context.Background(), key, start, stop).Result()
	return result, err
}

// LIndex 返回名称为key的list中index位置的元素
func (r *List) LIndex(key string, index int64) (string, error) {
	result, err := r.db.LIndex(context.Background(), key, index).Result()
	return result, err
}

// LSet 给名称为key的list中index位置的元素赋值，把原来的数据覆盖
func (r *List) LSet(key string, index int64, value interface{}) error {
	err := r.db.LSet(context.Background(), key, index, value).Err()
	return err
}

// LInsertBefore 在指定元素的前面插入多个元素
func (r *List) LInsertBefore(key string, target, value interface{}) error {
	err := r.db.LInsertBefore(context.Background(), key, target, value).Err()
	return err
}

// LInsertAfter 在指定元素的后面插入多个元素
func (r *List) LInsertAfter(key string, target, value interface{}) error {
	err := r.db.LInsertAfter(context.Background(), key, target, value).Err()
	return err
}

// LRem 删除列表中的数据，指定删除count个value数据
func (r *List) LRem(key string, count int64, value interface{}) error {
	if count <= 0 {
		return errors.New("count个数不能是负数")
	}
	err := r.db.LRem(context.Background(), key, count, value).Err()
	return err
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func (r *List) LPop(key string) (string, error) {
	result, err := r.db.LPop(context.Background(), key).Result()
	return result, err
}

// RPop 返回并删除名称为key的list中的尾元素
func (r *List) RPop(key string) (string, error) {
	result, err := r.db.RPop(context.Background(), key).Result()
	return result, err
}

// RPopLPush 返回并删除名称为srckey的list的尾元素，并将该元素添加到名称为dstkey的list的头部
func (r *List) RPopLPush(srcKey, destKey string) (string, error) {
	result, err := r.db.RPopLPush(context.Background(), srcKey, destKey).Result()
	return result, err
}
