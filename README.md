# zdpgo_redis

使用Golang操作Redis的快捷工具库

功能列表：

- 基于redis的分布式锁

## 版本历史

- v0.1.1 2022/04/16 新增：string类型的crud操作
- v0.1.2 2022/07/16 优化：代码结构优化

## 使用示例

### string类型的增删改查

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_redis"
)

func main() {
	r := zdpgo_redis.New(zdpgo_redis.Config{
		Host:     "localhost",
		Port:     6379,
		Database: 0,
	})

	// 增加
	r.String.Add("a", 1)
	r.String.AddMany("b", 1.1, "c", true, "d", "abc")

	// 查询
	aValue, err := r.String.Find("a")
	if err != nil {
		return
	}
	fmt.Println(aValue)

	manyValues, err := r.String.FindMany("b", "c", "d")
	if err != nil {
		return
	}
	fmt.Println(manyValues)

	// 修改
	r.String.Update("a", 1)
	r.String.UpdateMany("b", 1.1, "c", true, "d", "abc")

	keys, err := r.Common.Keys()
	if err != nil {
		return
	}
	fmt.Println("删除之前：", keys)

	// 删除
	r.String.Delete("a")
	r.String.DeleteMany("b", "c", "d")

	keys, err = r.Common.Keys()
	if err != nil {
		return
	}
	fmt.Println("删除之后：", keys)
}
```