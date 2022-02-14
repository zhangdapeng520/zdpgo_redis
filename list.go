package zdpgo_redis

import (
	"context"
	"errors"
)

// RPush 在名称为key的list尾添加一个值为value的元素
// 从右边追加元素
func (r *Redis) RPush(key string, values ...interface{}) error {
	err := r.db.RPush(context.Background(), key, values...).Err()
	if err != nil {
		r.log.Error("RPush 失败", "error", err.Error())
	}
	return err
}

// RPushX 仅当key存在的时候，从右边追加数据
func (r *Redis) RPushX(key string, values ...interface{}) error {
	err := r.db.RPushX(context.Background(), key, values...).Err()
	if err != nil {
		r.log.Error("RPushX 失败", "error", err.Error())
	}
	return err
}

// LPush 在名称为key的list头添加一个值为value的 元素
// 从左边追加元素
func (r *Redis) LPush(key string, values ...interface{}) error {
	err := r.db.LPush(context.Background(), key, values...).Err()
	if err != nil {
		r.log.Error("LPush 失败", "error", err.Error())
	}
	return err
}

// LPushX 仅当数据存在的时候才能插入
func (r *Redis) LPushX(key string, values ...interface{}) error {
	err := r.db.LPushX(context.Background(), key, values...).Err()
	if err != nil {
		r.log.Error("LPushX 失败", "error", err.Error())
	}
	return err
}

// LLen 返回名称为key的list的长度
func (r *Redis) LLen(key string) (int64, error) {
	result, err := r.db.LLen(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取列表长度失败：", result)
	}
	return result, err
}

// LRange 返回名称为key的list中start至end之间的元素
func (r *Redis) LRange(key string, start, stop int64) ([]string, error) {
	result, err := r.db.LRange(context.Background(), key, start, stop).Result()
	if err != nil {
		r.log.Error("批量获取列表元素失败：", err)
	}
	return result, err
}

// LTrim 截取名称为key的list，只保留start到top之间的数据
func (r *Redis) LTrim(key string, start, stop int64) (string, error) {
	result, err := r.db.LTrim(context.Background(), key, start, stop).Result()
	if err != nil {
		r.log.Error("LTrim 截取数据失败", "error", err.Error())
	}
	return result, err
}

// LIndex 返回名称为key的list中index位置的元素
func (r *Redis) LIndex(key string, index int64) (string, error) {
	result, err := r.db.LIndex(context.Background(), key, index).Result()
	if err != nil {
		r.log.Error("根据索引获取值失败：", err)
	}
	return result, err
}

// LSet 给名称为key的list中index位置的元素赋值，把原来的数据覆盖
func (r *Redis) LSet(key string, index int64, value interface{}) error {
	err := r.db.LSet(context.Background(), key, index, value).Err()
	if err != nil {
		r.log.Error("LSet 设置数据失败", "error", err.Error())
	}
	return err
}

// LInsertBefore 在指定元素的前面插入多个元素
func (r *Redis) LInsertBefore(key string, target, value interface{}) error {
	err := r.db.LInsertBefore(context.Background(), key, target, value).Err()
	if err != nil {
		r.log.Error("LInsertBefore 在指定元素的前面插入多个元素失败", "error", err.Error())
	}
	return err
}

// LInsertAfter 在指定元素的后面插入多个元素
func (r *Redis) LInsertAfter(key string, target, value interface{}) error {
	err := r.db.LInsertAfter(context.Background(), key, target, value).Err()
	if err != nil {
		r.log.Error("LInsertAfter 在指定元素的后面插入多个元素", "error", err.Error())
	}
	return err
}

// LRem 删除列表中的数据，指定删除count个value数据
func (r *Redis) LRem(key string, count int64, value interface{}) error {
	if count <= 0 {
		return errors.New("count个数不能是负数")
	}

	err := r.db.LRem(context.Background(), key, count, value).Err()
	if err != nil {
		r.log.Error("LRem 删除数据失败", "error", err.Error())
	}
	return err
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func (r *Redis) LPop(key string) (string, error) {
	result, err := r.db.LPop(context.Background(), key).Result()
	if err != nil {
		r.log.Error("LPop 删除列表首位元素失败", "error", err)
	}
	return result, err
}

// RPop 返回并删除名称为key的list中的尾元素
func (r *Redis) RPop(key string) (string, error) {
	result, err := r.db.RPop(context.Background(), key).Result()
	if err != nil {
		r.log.Error("RPop 删除列表末尾元素失败：", err)
	}
	return result, err
}

// RPopLPush 返回并删除名称为srckey的list的尾元素，并将该元素添加到名称为dstkey的list的头部
func (r *Redis) RPopLPush(srcKey, destKey string) (string, error) {
	result, err := r.db.RPopLPush(context.Background(), srcKey, destKey).Result()
	if err != nil {
		r.log.Error("转移列表元素失败：", err)
	}
	return result, err
}
