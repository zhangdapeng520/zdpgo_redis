package rset

import (
	"fmt"
	"testing"
)

func getSet() *Set {
	return NewSet("10.1.3.52", 6379, "", "", 0, 20)
}
func TestRedis_SAdd(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	members, err := r.SMembers(key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(members)
}

func TestRedis_SCard(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	card, err := r.SCard(key)
	if err != nil {
		t.Error(err)
	}
	t.Log(card)
}

func TestRedis_SIsMember(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	member, err := r.SIsMember(key, "a")
	if err != nil {
		t.Error(err)
	}
	t.Log(member)
}

func TestRedis_SInter(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	key2 := "set2"
	r.SAdd(key2, "a1")
	r.SAdd(key2, "b")
	r.SAdd(key2, "c1")

	inter, err := r.SInter(key, key2)
	if err != nil {
		t.Error(err)
	}
	t.Log(inter)
}

func TestRedis_SInterStore(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	key2 := "set2"
	r.SAdd(key2, "a1")
	r.SAdd(key2, "b")
	r.SAdd(key2, "c1")

	dstKey := "set3"
	err := r.SInterStore(dstKey, key, key2)
	if err != nil {
		t.Error(err)
	}

	members, err := r.SMembers(dstKey)
	if err != nil {
		t.Error(err)
	}
	t.Log(members)
}

func TestRedis_SUnion(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	key2 := "set2"
	r.SAdd(key2, "a1")
	r.SAdd(key2, "b")
	r.SAdd(key2, "c1")

	union, err := r.SUnion(key, key2)
	if err != nil {
		t.Error(err)
	}
	t.Log(union)
}

func TestRedis_SDiff(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	key2 := "set2"
	r.SAdd(key2, "a1")
	r.SAdd(key2, "b")
	r.SAdd(key2, "c1")

	union, err := r.SDiff(key, key2)
	if err != nil {
		t.Error(err)
	}
	t.Log(union)
}

func TestRedis_SDiffStore(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	key2 := "set2"
	r.SAdd(key2, "a1")
	r.SAdd(key2, "b")
	r.SAdd(key2, "c1")

	dstKey := "set3"
	err := r.SDiffStore(dstKey, key, key2)
	if err != nil {
		t.Error(err)
	}

	members, err := r.SMembers(dstKey)
	if err != nil {
		t.Error(err)
	}
	t.Log(members)
}

func TestRedis_SPop(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	pop, err := r.SPop(key)
	if err != nil {
		t.Error(err)
	}
	t.Log(pop)
}

func TestRedis_SRem(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	rem, err := r.SRem(key, "a")
	if err != nil {
		t.Error(err)
	}
	t.Log(rem)
}

func TestRedis_SRandMember(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	rem, err := r.SRandMember(key)
	if err != nil {
		t.Error(err)
	}
	t.Log(rem)
}

func TestRedis_SRandMemberN(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	rem, err := r.SRandMemberN(key, 2)
	if err != nil {
		t.Error(err)
	}
	t.Log(rem)
}

func TestRedis_SMembersMap(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	rem, err := r.SMembersMap(key)
	if err != nil {
		t.Error(err)
	}
	t.Log(rem)
}

func TestRedis_SMove(t *testing.T) {
	r := getSet()
	key := "set1"
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "a")
	r.SAdd(key, "b")
	r.SAdd(key, "c")

	rem, err := r.SMove(key, "set2", "a")
	if err != nil {
		t.Error(err)
	}
	t.Log(rem)

	members, err := r.SMembers(key)
	if err != nil {
		t.Error(err)
	}
	t.Log(members)

	sMembers, err := r.SMembers("set2")
	if err != nil {
		t.Error(err)
	}
	t.Log(sMembers)
}
