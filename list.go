package zdpgo_redis

import (
	"context"
)

// rpush(key, value)：在名称为key的list尾添加一个值为value的元素
// 从右边追加元素
func (r *Redis) RPush(key string, values ...interface{}) {
	r.db.RPush(context.Background(), key, values...)
}

// lpush(key, value)：在名称为key的list头添加一个值为value的 元素
// 从左边追加元素
func (r *Redis) LPush(key string, values ...interface{}) {
	r.db.LPush(context.Background(), key, values...)
}

// llen(key)：返回名称为key的list的长度
func (r *Redis) LLen(key string) (int64, error) {
	result, err := r.db.LLen(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取列表长度失败：", result)
	}
	return result, err
}

// lrange(key, start, end)：返回名称为key的list中start至end之间的元素
func (r *Redis) LRange(key string, start, stop int64) ([]string, error) {
	result, err := r.db.LRange(context.Background(), key, start, stop).Result()
	if err != nil {
		r.log.Error("批量获取列表元素失败：", err)
	}
	return result, err
}

// ltrim(key, start, end)：截取名称为key的list
func (r *Redis) LTrim(key string, start, stop int64) {
	r.db.LTrim(context.Background(), key, start, stop)
}

// lindex(key, index)：返回名称为key的list中index位置的元素
func (r *Redis) LIndex(key string, index int64) (string, error) {
	result, err := r.db.LIndex(context.Background(), key, index).Result()
	if err != nil {
		r.log.Error("根据索引获取值失败：", err)
	}
	return result, err
}

// lset(key, index, value)：给名称为key的list中index位置的元素赋值
func (r *Redis) LSet(key string, index int64, value interface{}) {
	r.db.LSet(context.Background(), key, index, value)
}

// lrem(key, count, value)：删除count个key的list中值为value的元素
func (r *Redis) LRem(key string, count int64, value interface{}) {
	r.db.LRem(context.Background(), key, count, value)
}

// lpop(key)：返回并删除名称为key的list中的首元素
func (r *Redis) LPop(key string) (string, error) {
	result, err := r.db.LPop(context.Background(), key).Result()
	if err != nil {
		r.log.Error("删除列表首位元素失败：", err)
	}
	return result, err
}

// rpop(key)：返回并删除名称为key的list中的尾元素
func (r *Redis) RPop(key string) (string, error) {
	result, err := r.db.RPop(context.Background(), key).Result()
	if err != nil {
		r.log.Error("删除列表末尾元素失败：", err)
	}
	return result, err
}

// rpoplpush(srckey, dstkey)：返回并删除名称为srckey的list的尾元素，并将该元素添加到名称为dstkey的list的头部
func (r *Redis) RPopLPush(srcKey, destKey string) (string, error) {
	result, err := r.db.RPopLPush(context.Background(), srcKey, destKey).Result()
	if err != nil {
		r.log.Error("转移列表元素失败：", err)
	}
	return result, err
}
