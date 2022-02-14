package zdpgo_redis

import "context"

// HSet 根据键设置hash的值
func (r *Redis) HSet(key string, values ...interface{}) error {
	err := r.db.HSet(context.Background(), key, values...).Err()
	if err != nil {
		r.log.Error("HSet 设置hash值失败", "error", err.Error())
	}
	return err
}

// HGet 根据key和field字段，查询field字段的值
func (r *Redis) HGet(key, field string) (string, error) {
	result, err := r.db.HGet(context.Background(), key, field).Result()
	if err != nil {
		r.log.Error("HGet 返回名称为key的hash中field对应的value失败", "error", err)
	}
	return result, err
}

// HMGet 批量获取hash键对应的值
func (r *Redis) HMGet(key string, fields ...string) ([]interface{}, error) {
	result, err := r.db.HMGet(context.Background(), key, fields...).Result()
	if err != nil {
		r.log.Error("HMGet 批量获取hash的属性值失败", "error", err)
	}
	return result, err
}

// HMSet 批量添加hash键值
func (r *Redis) HMSet(key string, fieldValues ...interface{}) error {
	err := r.db.HMSet(context.Background(), key, fieldValues...).Err()
	if err != nil {
		r.log.Error("HMSet 设置hash值失败", "error", err.Error())
	}
	return err
}

// HSetNX 如果字段不存在才设置hash值
func (r *Redis) HSetNX(key string, field string, value interface{}) error {
	err := r.db.HSetNX(context.Background(), key, field, value).Err()
	if err != nil {
		r.log.Error("HMSet 设置hash值失败", "error", err.Error())
	}
	return err
}

// HIncrBy 根据key和field字段，累加数值。
func (r *Redis) HIncrBy(key string, field string, value int64) error {
	err := r.db.HIncrBy(context.Background(), key, field, value).Err()
	if err != nil {
		r.log.Error("HIncrBy 累加数值失败", "error", err.Error())
	}
	return err
}

// HIncrByFloat 增长指定的浮点数
func (r *Redis) HIncrByFloat(key string, field string, value float64) error {
	err := r.db.HIncrByFloat(context.Background(), key, field, value).Err()
	if err != nil {
		r.log.Error("HIncrByFloat 累加数值失败", "error", err.Error())
	}
	return err
}

// HIncr 让hash的指定字段自增1
func (r *Redis) HIncr(key string, field string) error {
	err := r.db.HIncrBy(context.Background(), key, field, 1).Err()
	if err != nil {
		r.log.Error("字段自增失败", "error", err.Error())
	}
	return err
}

// HExists 检测hash字段名是否存在。
func (r *Redis) HExists(key string, field string) (bool, error) {
	result, err := r.db.HExists(context.Background(), key, field).Result()
	if err != nil {
		r.log.Error("判断hash的指定字段是否存在失败：", err, key, field)
	}
	return result, err
}

// HDel 根据key和字段名，删除hash字段，支持批量删除hash字段
func (r *Redis) HDel(key string, fields ...string) error {
	err := r.db.HDel(context.Background(), key, fields...).Err()
	if err != nil {
		r.log.Error("HDel 删除数据失败", "error", err.Error())
	}
	return err
}

// HLen 获取hash的字段个数
func (r *Redis) HLen(key string) (int64, error) {
	result, err := r.db.HLen(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取hash的字段个数失败：", err, key)
	}
	return result, err
}

// HKeys 获取hash的所有键
func (r *Redis) HKeys(key string) ([]string, error) {
	result, err := r.db.HKeys(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取hash的所有键失败：", err, key)
	}
	return result, err
}

// HVals 获取hash所有的值
func (r *Redis) HVals(key string) ([]string, error) {
	result, err := r.db.HVals(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取hash所有的值失败：", err, key)
	}
	return result, err
}

// HGetAll 根据key查询所有字段和值
func (r *Redis) HGetAll(key string) (map[string]string, error) {
	result, err := r.db.HGetAll(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取hash所有的值失败：", err, key)
	}
	return result, err
}
