package zdpgo_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// Z 表示已排序的集合成员
type Z struct {
	Score  float64     // 分数
	Member interface{} // 元素名
}

func z2Zs(members []*Z) []*redis.Z {
	var membersNew []*redis.Z
	for _, v := range members {
		t := &redis.Z{
			Score:  v.Score,
			Member: v.Member,
		}
		membersNew = append(membersNew, t)
	}
	return membersNew
}

func z2Z(member *Z) *redis.Z {
	t := &redis.Z{
		Score:  member.Score,
		Member: member.Member,
	}
	return t
}

// ZAdd 添加一个或者多个元素到集合，如果元素已经存在则更新分数
func (r *Redis) ZAdd(key string, members ...*Z) error {
	membersNew := z2Zs(members)
	err := r.db.ZAdd(context.Background(), key, membersNew...).Err()
	return err
}

// ZAddX 存在才添加
func (r *Redis) ZAddX(key string, members ...*Z) error {
	membersNew := z2Zs(members)
	err := r.db.ZAddXX(context.Background(), key, membersNew...).Err()
	return err
}

// ZCard 查看集合元素的个数
func (r *Redis) ZCard(key string) (int64, error) {
	size, err := r.db.ZCard(context.Background(), key).Result()
	return size, err
}

// ZRange 取出一段分数
func (r *Redis) ZRange(key string, start int64, end int64) ([]string, error) {
	result, err := r.db.ZRange(context.Background(), key, start, end).Result()
	return result, err
}

// ZIncrBy 增加分数
func (r *Redis) ZIncrBy(key string, score float64, member string) (float64, error) {
	result, err := r.db.ZIncrBy(context.Background(), key, score, member).Result()
	return result, err
}

// ZCount 统计某个分数段内的元素个数
func (r *Redis) ZCount(key string, min, max string) (int64, error) {
	result, err := r.db.ZCount(context.Background(), key, min, max).Result()
	return result, err
}

// ZScore 查询集合元素的分数
func (r *Redis) ZScore(key string, member string) (float64, error) {
	result, err := r.db.ZScore(context.Background(), key, member).Result()
	return result, err
}

// ZRank 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
func (r *Redis) ZRank(key string, member string) (int64, error) {
	result, err := r.db.ZRank(context.Background(), key, member).Result()
	return result, err
}

// ZRem 删除集合元素
func (r *Redis) ZRem(key string, members ...interface{}) (int64, error) {
	result, err := r.db.ZRem(context.Background(), key, members...).Result()
	return result, err
}

// ZRemRangeByRank 根据排名删除分数，支持按负数排名
func (r *Redis) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	result, err := r.db.ZRemRangeByRank(context.Background(), key, start, stop).Result()
	return result, err
}
