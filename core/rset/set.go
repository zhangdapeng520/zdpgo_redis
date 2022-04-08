package rset

import (
	"context"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_redis/libs/redis"
)

type Set struct {
	db *redis.Client // redis连接对象
}

func NewSet(host string, port int, username, password string, db, poolSize int) *Set {
	s := Set{}
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

// SAdd 向名称为key的set中添加元素member
func (r *Set) SAdd(key string, values ...interface{}) error {
	err := r.db.SAdd(context.Background(), key, values...).Err()
	return err
}

// SRem 删除名称为 key 的 set 中的元素 member,并返回删除的元素个数
func (r *Set) SRem(key string, value interface{}) (int64, error) {
	result, err := r.db.SRem(context.Background(), key, value).Result()
	return result, err
}

// SPop 随机返回集合中的一个元素，并且删除这个元素
func (r *Set) SPop(key string) (string, error) {
	result, err := r.db.SPop(context.Background(), key).Result()
	return result, err
}

// SPopN 随机返回集合中的count个元素，并且删除这些元素
func (r *Set) SPopN(key string, count int64) ([]string, error) {
	result, err := r.db.SPopN(context.Background(), key, count).Result()
	return result, err
}

// SMove 移动集合source中的一个member元素到集合destination中去
func (r *Set) SMove(srcKey, dstKey string, value interface{}) (bool, error) {
	result, err := r.db.SMove(context.Background(), srcKey, dstKey, value).Result()
	return result, err
}

// SCard 获取集合set元素个数
func (r *Set) SCard(key string) (int64, error) {
	result, err := r.db.SCard(context.Background(), key).Result()
	return result, err
}

// SIsMember 判断元素member是否在集合set中
func (r *Set) SIsMember(key string, value interface{}) (bool, error) {
	result, err := r.db.SIsMember(context.Background(), key, value).Result()
	return result, err
}

// SInter 求交集
func (r *Set) SInter(keys ...string) ([]string, error) {
	result, err := r.db.SInter(context.Background(), keys...).Result()
	return result, err
}

// SInterStore 求交集并将交集保存到dstkey的集合
func (r *Set) SInterStore(dstKey string, keys ...string) error {
	err := r.db.SInterStore(context.Background(), dstKey, keys...).Err()
	return err
}

// SUnion 求并集
func (r *Set) SUnion(keys ...string) ([]string, error) {
	result, err := r.db.SUnion(context.Background(), keys...).Result()
	return result, err
}

// SUnionStore 求并集并将并集保存到dstkey的集合
func (r *Set) SUnionStore(dstKey string, keys ...string) error {
	err := r.db.SUnionStore(context.Background(), dstKey, keys...).Err()
	return err
}

// SDiff 求差集
func (r *Set) SDiff(keys ...string) ([]string, error) {
	result, err := r.db.SDiff(context.Background(), keys...).Result()
	return result, err
}

// SDiffStore 求差集并将差集保存到dstkey的集合
func (r *Set) SDiffStore(dstKey string, keys ...string) error {
	err := r.db.SDiffStore(context.Background(), dstKey, keys...).Err()
	return err
}

// SMembers 返回名称为key的set的所有元素
func (r *Set) SMembers(key string) ([]string, error) {
	result, err := r.db.SMembers(context.Background(), key).Result()
	return result, err
}

// SRandMember 随机返回名称为 key 的 set 的一个元素
func (r *Set) SRandMember(key string) (string, error) {
	result, err := r.db.SRandMember(context.Background(), key).Result()
	return result, err
}

// SRandMemberN 随机返回名称为 key 的 set 的count个元素
func (r *Set) SRandMemberN(key string, count int64) ([]string, error) {
	result, err := r.db.SRandMemberN(context.Background(), key, count).Result()
	return result, err
}

// SMembersMap 把集合里的元素转换成map的key
func (r *Set) SMembersMap(key string) (map[string]struct{}, error) {
	result, err := r.db.SMembersMap(context.Background(), key).Result()
	return result, err
}
