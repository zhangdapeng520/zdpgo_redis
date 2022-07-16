package rstring

import (
	"github.com/zhangdapeng520/zdpgo_test/libs/assert"
	"testing"
)

// 测试添加
func TestString_Add(t *testing.T) {
	// 创建断言对象
	ast := assert.New(t)

	// 创建redis操作对象
	s := getString()

	// 创建表格
	var tests = []struct {
		key    string      // 输入：key
		value  interface{} // 输入：value
		except error       // 结果：错误对象
	}{
		{"a", 1, nil},
		{"b", 1.1, nil},
		{"c", true, nil},
		{"d", "abc123Abc", nil},
	}

	// 遍历表格数据
	for _, tt := range tests {
		err := s.Add(tt.key, tt.value)
		ast.Equal(tt.except, err) // 断言
	}
}

// 测试同时添加多个
func TestString_AddMany(t *testing.T) {
	// 创建断言对象
	ast := assert.New(t)

	// 创建redis操作对象
	s := getString()

	// 创建表格
	var tests = []struct {
		key1   string      // 输入：key
		value1 interface{} // 输入：value
		key2   string      // 输入：key
		value2 interface{} // 输入：value
		key3   string      // 输入：key
		value3 interface{} // 输入：value
		except error       // 结果：错误对象
	}{
		{"a", 1, "b", 1.1, "d", "abc123Abc", nil},
		{"a", 1, "b", 1.1, "d", "abc123Abc", nil},
	}

	// 遍历表格数据
	for _, tt := range tests {
		err := s.AddMany(tt.key1, tt.value1, tt.key2, tt.value2, tt.key3, tt.value3)
		ast.Equal(tt.except, err) // 断言
	}
}

// 测试修改
func TestString_Update(t *testing.T) {
	// 创建断言对象
	ast := assert.New(t)

	// 创建redis操作对象
	s := getString()

	// 创建表格
	var tests = []struct {
		key    string      // 输入：key
		value  interface{} // 输入：value
		except error       // 结果：错误对象
	}{
		{"a", 1, nil},
		{"b", 1.1, nil},
		{"c", true, nil},
		{"d", "abc123Abc", nil},
	}

	// 遍历表格数据
	for _, tt := range tests {
		err := s.Update(tt.key, tt.value)
		ast.Equal(tt.except, err) // 断言
	}
}

// 测试修改多个
func TestString_UpdateMany(t *testing.T) {
	// 创建断言对象
	ast := assert.New(t)

	// 创建redis操作对象
	s := getString()

	// 创建表格
	var tests = []struct {
		key1   string      // 输入：key
		value1 interface{} // 输入：value
		key2   string      // 输入：key
		value2 interface{} // 输入：value
		key3   string      // 输入：key
		value3 interface{} // 输入：value
		except error       // 结果：错误对象
	}{
		{"a", 1, "b", 1.1, "d", "abc123Abc", nil},
		{"a", 1, "b", 1.1, "d", "abc123Abc", nil},
	}

	// 遍历表格数据
	for _, tt := range tests {
		err := s.UpdateMany(tt.key1, tt.value1, tt.key2, tt.value2, tt.key3, tt.value3)
		ast.Equal(tt.except, err) // 断言
	}
}

// 测试查询
func TestString_Find(t *testing.T) {
	// 创建断言对象
	ast := assert.New(t)

	// 创建redis操作对象
	s := getString()

	// 创建表格
	var tests = []struct {
		key    string      // 输入：key
		value  interface{} // 输入：value
		except interface{} // 结果：错误对象
	}{
		{"a", 1, "1"},
		{"b", 1.1, "1.1"},
		{"c", true, "1"}, // 布尔值存储的值数字1
		{"d", "abc123Abc", "abc123Abc"},
	}

	// 遍历表格数据
	for _, tt := range tests {
		// 添加
		err := s.Add(tt.key, tt.value)
		ast.Equal(err, nil) // 断言

		// 查询
		value, err := s.Find(tt.key)
		ast.Equal(err, nil)
		ast.Equal(tt.except, value)
	}
}

// 测试查询多个
func TestString_FindMany(t *testing.T) {
	// 创建断言对象
	ast := assert.New(t)

	// 创建redis操作对象
	s := getString()

	// 创建表格
	var tests = []struct {
		key1    string      // 输入：key
		value1  interface{} // 输入：value
		except1 interface{} // 结果

		key2    string      // 输入：key
		value2  interface{} // 输入：value
		except2 interface{} // 结果

		key3    string      // 输入：key
		value3  interface{} // 输入：value
		except3 interface{} // 结果
	}{
		{"a", 1, "1", "b", 1.1, "1.1", "d", "abc123Abc", "abc123Abc"},
	}

	// 遍历表格数据
	for _, tt := range tests {
		// 添加
		err := s.AddMany(tt.key1, tt.value1, tt.key2, tt.value2, tt.key3, tt.value3)
		ast.Equal(err, nil) // 断言

		// 取出
		keys := []string{tt.key1, tt.key2, tt.key3}
		values, err := s.FindMany(keys...)
		ast.Equal(err, nil) // 断言
		ast.Equal(values[0], tt.except1)
		ast.Equal(values[1], tt.except2)
		ast.Equal(values[2], tt.except3)
	}
}
