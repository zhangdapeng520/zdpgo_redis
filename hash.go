package zdpgo_redis

import "context"

// 根据键设置hash的值
func (r *Redis) HSet(key, field string, value interface{}) {
	r.db.HSet(context.Background(), key, field, value)
}

// 根据键获取hash的值
func (r *Redis) HGet(key, field string) (string, error) {
	result, err := r.db.HGet(context.Background(), key, field).Result()
	if err != nil {
		r.log.Error("返回名称为key的hash中field对应的value失败：", err)
	}
	return result, err
}

// 批量获取hash键对应的值
func (r *Redis) HMGet(key string, fields ...string) ([]interface{}, error) {
	result, err := r.db.HMGet(context.Background(), key, fields...).Result()
	if err != nil {
		r.log.Error("批量获取hash的属性值失败：", err)
	}
	return result, err
}

// 批量添加hash键值
func (r *Redis) HMSet(key string, fieldValues ...interface{}) {
	r.db.HMSet(context.Background(), key, fieldValues...)
}

// 让hash的指定字段自增n
func (r *Redis) HIncrBy(key string, field string, value int64) {
	r.db.HIncrBy(context.Background(), key, field, value)
}

// 让hash的指定字段自增1
func (r *Redis) HIncr(key string, field string) {
	r.db.HIncrBy(context.Background(), key, field, 1)
}

// 判断hash的指定字段是否存在
func (r *Redis) HExists(key string, field string) (bool, error) {
	result, err := r.db.HExists(context.Background(), key, field).Result()
	if err != nil {
		r.log.Error("判断hash的指定字段是否存在失败：", err, key, field)
	}
	return result, err
}

// 删除hash的指定字段
func (r *Redis) HDel(key string, field string) {
	r.db.HDel(context.Background(), key, field)
}

// 获取hash的字段个数
func (r *Redis) HLen(key string) (int64, error) {
	result, err := r.db.HLen(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取hash的字段个数失败：", err, key)
	}
	return result, err
}

// 获取hash的所有键
func (r *Redis) HKeys(key string) ([]string, error) {
	result, err := r.db.HKeys(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取hash的所有键失败：", err, key)
	}
	return result, err
}

// 获取hash所有的值
func (r *Redis) HVals(key string) ([]string, error) {
	result, err := r.db.HVals(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取hash所有的值失败：", err, key)
	}
	return result, err
}

// 获取hash的所有键值对
func (r *Redis) HGetAll(key string) (map[string]string, error) {
	result, err := r.db.HGetAll(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取hash所有的值失败：", err, key)
	}
	return result, err
}
