package zdpgo_redis

import (
	"fmt"
	"testing"
)

func TestRedis_ZCard(t *testing.T) {
	r := prepareRedis()
	err := r.ZAdd("key", &Z{2.5, "tizi"})
	if err != nil {
		t.Error(err)
	}

	size, err := r.ZCard("key")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(size)
}

func TestRedis_ZRange(t *testing.T) {
	r := prepareRedis()
	key := "s1"
	err := r.ZAdd(key,
		&Z{2.5, "a"},
		&Z{2.1, "b"},
		&Z{3.1, "c"},
	)
	if err != nil {
		t.Error(err)
	}

	zRange, err := r.ZRange(key, 0, -1) // 默认从小到大排序
	if err != nil {
		t.Error(err)
	}
	t.Log(zRange)
}

func TestRedis_ZAddX(t *testing.T) {
	r := prepareRedis()
	key := "s1"
	err := r.ZAddX(key,
		&Z{2.5, "a"},
		&Z{2.1, "b"},
		&Z{3.1, "c"},
	)
	if err != nil {
		t.Error(err)
	}

	err = r.ZAddX(key,
		&Z{2.5, "aa"},
		&Z{2.1, "bb"},
		&Z{3.1, "cc"},
	)
	if err != nil {
		t.Error(err)
	}

	zRange, err := r.ZRange(key, 0, -1) // 默认从小到大排序
	if err != nil {
		t.Error(err)
	}
	t.Log(zRange)
}

func TestRedis_ZIncrBy(t *testing.T) {
	r := prepareRedis()
	//统计开发语言排行榜
	zsetKey := "language_rank"
	languages := []*Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 92.0, Member: "C/C++"},
	}
	r.ZAdd(zsetKey, languages...)

	//给元素Python加上8分
	r.ZIncrBy(zsetKey, 8, "Python")

	zRange, err := r.ZRange(zsetKey, 0, -1) // 默认从小到大排序
	if err != nil {
		t.Error(err)
	}
	t.Log(zRange)
}

func TestRedis_ZCount(t *testing.T) {
	r := prepareRedis()
	//统计开发语言排行榜
	zsetKey := "language_rank"
	r.Del(zsetKey)
	languages := []*Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 92.0, Member: "C/C++"},
	}
	r.ZAdd(zsetKey, languages...)

	// 添加一个元素到集合
	r.ZAdd(zsetKey, &Z{Score: 87, Member: "Vue"})

	//给元素Vue加上8分
	r.ZIncrBy(zsetKey, 8, "Vue")

	zRange, err := r.ZRange(zsetKey, 0, -1) // 默认从小到大排序
	if err != nil {
		t.Error(err)
	}
	t.Log(zRange)

	// 统计某个分数段内的元素个数，这里是查询的95<分数<100的元素个数
	count, err := r.ZCount(zsetKey, "95", "100")
	if err != nil {
		t.Error(err)
	}
	t.Log(count)
}

func TestRedis_ZScore(t *testing.T) {
	r := prepareRedis()
	//统计开发语言排行榜
	zsetKey := "language_rank"
	r.Del(zsetKey)
	languages := []*Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 92.0, Member: "C/C++"},
	}
	r.ZAdd(zsetKey, languages...)

	// 添加一个元素到集合
	r.ZAdd(zsetKey, &Z{Score: 87, Member: "Vue"})

	//给元素Vue加上8分
	r.ZIncrBy(zsetKey, 8, "Vue")

	score, err := r.ZScore(zsetKey, "Golang")
	if err != nil {
		t.Error(err)
	}
	t.Log(score)
}

func TestRedis_ZRank(t *testing.T) {
	r := prepareRedis()

	//统计开发语言排行榜
	zsetKey := "language_rank"
	r.Del(zsetKey)
	languages := []*Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 92.0, Member: "C/C++"},
	}
	r.ZAdd(zsetKey, languages...)

	// 添加一个元素到集合
	r.ZAdd(zsetKey, &Z{Score: 87, Member: "Vue"})

	// 给元素Vue加上8分
	r.ZIncrBy(zsetKey, 8, "Vue")

	score, err := r.ZScore(zsetKey, "Golang")
	if err != nil {
		t.Error(err)
	}
	t.Log(score)
}

func TestRedis_ZRem(t *testing.T) {
	r := prepareRedis()

	//统计开发语言排行榜
	zsetKey := "language_rank"
	r.Del(zsetKey)
	languages := []*Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 92.0, Member: "C/C++"},
	}
	r.ZAdd(zsetKey, languages...)

	// 添加一个元素到集合
	r.ZAdd(zsetKey, &Z{Score: 87, Member: "Vue"})

	// 给元素Vue加上8分
	r.ZRem(zsetKey, "Golang", "Java")

	zRange, err := r.ZRange(zsetKey, 0, -1)
	if err != nil {
		t.Error(err)
	}
	t.Log(zRange)
}

func TestRedis_ZRemRangeByRank(t *testing.T) {
	r := prepareRedis()

	//统计开发语言排行榜
	zsetKey := "language_rank"
	r.Del(zsetKey)
	languages := []*Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 92.0, Member: "C/C++"},
	}
	r.ZAdd(zsetKey, languages...)

	// 添加一个元素到集合
	r.ZAdd(zsetKey, &Z{Score: 87, Member: "Vue"})

	zRange, err := r.ZRange(zsetKey, 0, -1)
	if err != nil {
		t.Error(err)
	}
	t.Log(zRange)

	// 正排名删除
	rank, err := r.ZRemRangeByRank(zsetKey, 0, 2)
	if err != nil {
		t.Error(err)
	}
	t.Log(rank)

	// 负排名删除
	byRank, err := r.ZRemRangeByRank(zsetKey, -2, -1)
	if err != nil {
		t.Error(err)
	}
	t.Log(byRank)

	zRange, err = r.ZRange(zsetKey, 0, -1)
	if err != nil {
		t.Error(err)
	}
	t.Log(zRange)
}
