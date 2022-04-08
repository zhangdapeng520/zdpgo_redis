package rhash

import (
	"context"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_redis/libs/redis"
)

type Hash struct {
	db *redis.Client // redis连接对象
}

func NewHash(host string, port int, username, password string, db, poolSize int) *Hash {
	s := Hash{}
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

// HSet 根据键设置hash的值
func (r *Hash) HSet(key string, values ...interface{}) error {
	err := r.db.HSet(context.Background(), key, values...).Err()
	return err
}

// HGet 根据key和field字段，查询field字段的值
func (r *Hash) HGet(key, field string) (string, error) {
	result, err := r.db.HGet(context.Background(), key, field).Result()
	return result, err
}

// HMGet 批量获取hash键对应的值
func (r *Hash) HMGet(key string, fields ...string) ([]interface{}, error) {
	result, err := r.db.HMGet(context.Background(), key, fields...).Result()
	return result, err
}

// HMSet 批量添加hash键值
func (r *Hash) HMSet(key string, fieldValues ...interface{}) error {
	err := r.db.HMSet(context.Background(), key, fieldValues...).Err()
	return err
}

// HSetNX 如果字段不存在才设置hash值
func (r *Hash) HSetNX(key string, field string, value interface{}) error {
	err := r.db.HSetNX(context.Background(), key, field, value).Err()
	return err
}

// HIncrBy 根据key和field字段，累加数值。
func (r *Hash) HIncrBy(key string, field string, value int64) error {
	err := r.db.HIncrBy(context.Background(), key, field, value).Err()
	return err
}

// HIncrByFloat 增长指定的浮点数
func (r *Hash) HIncrByFloat(key string, field string, value float64) error {
	err := r.db.HIncrByFloat(context.Background(), key, field, value).Err()
	return err
}

// HIncr 让hash的指定字段自增1
func (r *Hash) HIncr(key string, field string) error {
	err := r.db.HIncrBy(context.Background(), key, field, 1).Err()
	return err
}

// HExists 检测hash字段名是否存在。
func (r *Hash) HExists(key string, field string) (bool, error) {
	result, err := r.db.HExists(context.Background(), key, field).Result()
	return result, err
}

// HDel 根据key和字段名，删除hash字段，支持批量删除hash字段
func (r *Hash) HDel(key string, fields ...string) error {
	err := r.db.HDel(context.Background(), key, fields...).Err()
	return err
}

// HLen 获取hash的字段个数
func (r *Hash) HLen(key string) (int64, error) {
	result, err := r.db.HLen(context.Background(), key).Result()
	return result, err
}

// HKeys 获取hash的所有键
func (r *Hash) HKeys(key string) ([]string, error) {
	result, err := r.db.HKeys(context.Background(), key).Result()
	return result, err
}

// HVals 获取hash所有的值
func (r *Hash) HVals(key string) ([]string, error) {
	result, err := r.db.HVals(context.Background(), key).Result()
	return result, err
}

// HGetAll 根据key查询所有字段和值
func (r *Hash) HGetAll(key string) (map[string]string, error) {
	result, err := r.db.HGetAll(context.Background(), key).Result()
	return result, err
}
