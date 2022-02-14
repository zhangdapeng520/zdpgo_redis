package zdpgo_redis

import "context"

// SAdd 向名称为key的set中添加元素member
func (r *Redis) SAdd(key string, values ...interface{}) error {
	err := r.db.SAdd(context.Background(), key, values...).Err()
	if err != nil {
		r.log.Error("SAdd 添加数据失败", "error", err.Error())
	}
	return err
}

// SRem 删除名称为 key 的 set 中的元素 member,并返回删除的元素个数
func (r *Redis) SRem(key string, value interface{}) (int64, error) {
	result, err := r.db.SRem(context.Background(), key, value).Result()
	if err != nil {
		r.log.Error("SRem 删除名称为 key 的 set 中的元素 member,并返回删除的元素个数", "error", err.Error())
	}
	return result, err
}

// SPop 随机返回集合中的一个元素，并且删除这个元素
func (r *Redis) SPop(key string) (string, error) {
	result, err := r.db.SPop(context.Background(), key).Result()
	if err != nil {
		r.log.Error("SPop 随机返回集合中的一个元素失败", "error", err.Error())
	}
	return result, err
}

// SPopN 随机返回集合中的count个元素，并且删除这些元素
func (r *Redis) SPopN(key string, count int64) ([]string, error) {
	result, err := r.db.SPopN(context.Background(), key, count).Result()
	if err != nil {
		r.log.Error("SPopN 随机返回集合中的count个元素失败", "error", err.Error())
	}
	return result, err
}

// SMove 移动集合source中的一个member元素到集合destination中去
func (r *Redis) SMove(srcKey, dstKey string, value interface{}) (bool, error) {
	result, err := r.db.SMove(context.Background(), srcKey, dstKey, value).Result()
	if err != nil {
		r.log.Error("移动集合source中的一个member元素到集合destination中去", "error", err)
	}
	return result, err
}

// SCard 获取集合set元素个数
func (r *Redis) SCard(key string) (int64, error) {
	result, err := r.db.SCard(context.Background(), key).Result()
	if err != nil {
		r.log.Error("返回名称为key的set的基数失败", "error", err)
	}
	return result, err
}

// SIsMember 判断元素member是否在集合set中
func (r *Redis) SIsMember(key string, value interface{}) (bool, error) {
	result, err := r.db.SIsMember(context.Background(), key, value).Result()
	if err != nil {
		r.log.Error("检查member是否是名称为key的set的元素失败", "error", err)
	}
	return result, err
}

// SInter 求交集
func (r *Redis) SInter(keys ...string) ([]string, error) {
	result, err := r.db.SInter(context.Background(), keys...).Result()
	if err != nil {
		r.log.Error("求交集失败", "error", err)
	}
	return result, err
}

// SInterStore 求交集并将交集保存到dstkey的集合
func (r *Redis) SInterStore(dstKey string, keys ...string) error {
	err := r.db.SInterStore(context.Background(), dstKey, keys...).Err()
	if err != nil {
		r.log.Error("求交集失败", "error", err.Error())
	}
	return err
}

// SUnion 求并集
func (r *Redis) SUnion(keys ...string) ([]string, error) {
	result, err := r.db.SUnion(context.Background(), keys...).Result()
	if err != nil {
		r.log.Error("求并集失败：", err)
	}
	return result, err
}

// SUnionStore 求并集并将并集保存到dstkey的集合
func (r *Redis) SUnionStore(dstKey string, keys ...string) error {
	err := r.db.SUnionStore(context.Background(), dstKey, keys...).Err()
	if err != nil {
		r.log.Error("求交集失败")
	}
	return err
}

// SDiff 求差集
func (r *Redis) SDiff(keys ...string) ([]string, error) {
	result, err := r.db.SDiff(context.Background(), keys...).Result()
	if err != nil {
		r.log.Error("求差集失败", "error", err.Error())
	}
	return result, err
}

// SDiffStore 求差集并将差集保存到dstkey的集合
func (r *Redis) SDiffStore(dstKey string, keys ...string) error {
	err := r.db.SDiffStore(context.Background(), dstKey, keys...).Err()
	if err != nil {
		r.log.Error("求差集失败", "error", err.Error())
	}
	return err
}

// SMembers 返回名称为key的set的所有元素
func (r *Redis) SMembers(key string) ([]string, error) {
	result, err := r.db.SMembers(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取名称为key的set的所有元素失败：", err)
	}
	return result, err
}

// SRandMember 随机返回名称为 key 的 set 的一个元素
func (r *Redis) SRandMember(key string) (string, error) {
	result, err := r.db.SRandMember(context.Background(), key).Result()
	if err != nil {
		r.log.Error("SRandMember 随机返回名称为 key 的 set 的一个元素失败", "error", err)
	}
	return result, err
}

// SRandMemberN 随机返回名称为 key 的 set 的count个元素
func (r *Redis) SRandMemberN(key string, count int64) ([]string, error) {
	result, err := r.db.SRandMemberN(context.Background(), key, count).Result()
	if err != nil {
		r.log.Error("SRandMemberN 随机返回名称为 key 的 set 的count个元素失败", "error", err)
	}
	return result, err
}

// SMembersMap 把集合里的元素转换成map的key
func (r *Redis) SMembersMap(key string) (map[string]struct{}, error) {
	result, err := r.db.SMembersMap(context.Background(), key).Result()
	if err != nil {
		r.log.Error("SMembersMap 把集合里的元素转换成map的key失败", "error", err)
	}
	return result, err
}
