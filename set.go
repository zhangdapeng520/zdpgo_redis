package zdpgo_redis

import "context"

// sadd(key, member)：向名称为key的set中添加元素member
func (r *Redis) SAdd(key string, values ...interface{}) {
	r.db.SAdd(context.Background(), key, values...)
}

// srem(key, member) ：删除名称为key的set中的元素member
func (r *Redis) SRem(key string, value interface{}) {
	r.db.SRem(context.Background(), key, value)
}

// spop(key) ：随机返回并删除名称为key的set中一个元素
func (r *Redis) SPop(key string) {
	r.db.SPop(context.Background(), key)
}

// smove(srckey, dstkey, member) ：移到集合元素
func (r *Redis) SMove(srcKey, dstKey string, value interface{}) {
	r.db.SMove(context.Background(), srcKey, dstKey, value)
}

// scard(key) ：返回名称为key的set的基数
func (r *Redis) SCard(key string) (int64, error) {
	result, err := r.db.SCard(context.Background(), key).Result()
	if err != nil {
		r.log.Error("返回名称为key的set的基数失败：", err)
	}
	return result, err
}

// sismember(key, member) ：member是否是名称为key的set的元素
func (r *Redis) SIsMember(key string, value interface{}) (bool, error) {
	result, err := r.db.SIsMember(context.Background(), key, value).Result()
	if err != nil {
		r.log.Error("检查member是否是名称为key的set的元素失败：", err)
	}
	return result, err
}

// sinter(key1, key2,…key N) ：求交集
func (r *Redis) SInter(keys ...string) ([]string, error) {
	result, err := r.db.SInter(context.Background(), keys...).Result()
	if err != nil {
		r.log.Error("求交集失败：", err)
	}
	return result, err
}

// sinterstore(dstkey, (keys)) ：求交集并将交集保存到dstkey的集合
func (r *Redis) SInterStore(dstKey string, keys ...string) {
	r.db.SInterStore(context.Background(), dstKey, keys...)
}

// sunion(key1, (keys)) ：求并集
func (r *Redis) SUnion(keys ...string) ([]string, error) {
	result, err := r.db.SUnion(context.Background(), keys...).Result()
	if err != nil {
		r.log.Error("求并集失败：", err)
	}
	return result, err
}

// sunionstore(dstkey, (keys)) ：求并集并将并集保存到dstkey的集合
func (r *Redis) SUnionStore(dstKey string, keys ...string) {
	r.db.SUnionStore(context.Background(), dstKey, keys...)
}

// sdiff(key1, (keys)) ：求差集
func (r *Redis) SDiff(keys ...string) ([]string, error) {
	result, err := r.db.SDiff(context.Background(), keys...).Result()
	if err != nil {
		r.log.Error("求差集失败：", err)
	}
	return result, err
}

// sdiffstore(dstkey, (keys)) ：求差集并将差集保存到dstkey的集合
func (r *Redis) SDiffStore(dstKey string, keys ...string) {
	r.db.SDiffStore(context.Background(), dstKey, keys...)
}

// smembers(key) ：返回名称为key的set的所有元素
func (r *Redis) SMembers(key string) ([]string, error) {
	result, err := r.db.SMembers(context.Background(), key).Result()
	if err != nil {
		r.log.Error("获取名称为key的set的所有元素失败：", err)
	}
	return result, err
}

// srandmember(key) ：随机返回名称为key的set的一个元素
func (r *Redis) SRandMember(key string) (string, error) {
	result, err := r.db.SRandMember(context.Background(), key).Result()
	if err != nil {
		r.log.Error("随机返回名称为key的set的一个元素失败：", err)
	}
	return result, err
}
