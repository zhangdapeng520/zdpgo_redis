package rstring

import (
	"context"
)

// Add 根据键设置值，过期时间默认为30天
func (r *String) Add(key string, value interface{}) error {
	return r.Set(key, value)
}

// AddMany 同时添加多个Key、Value键值对
func (r *String) AddMany(kvs ...interface{}) error {
	return r.MSet(kvs...)
}

// Delete 删除key
func (r *String) Delete(key string) error {
	return r.common.Del(key)
}

// DeleteMany 同时删除多个key
func (r *String) DeleteMany(keys ...string) error {
	return r.common.Del(keys...)
}

// Update 修改key对应的值
func (r *String) Update(key string, value interface{}) error {
	return r.Set(key, value)
}

// UpdateMany 同时修改多个key对应的value
func (r *String) UpdateMany(kvs ...interface{}) error {
	return r.MSet(kvs...)
}

// Find 根据键获取值
func (r *String) Find(key string) (value string, err error) {
	value, err = r.db.Get(context.Background(), key).Result()
	return
}

// FindMany 同时获取多个键对应的值
func (r *String) FindMany(keys ...string) ([]interface{}, error) {
	return r.MGet(keys...)
}
