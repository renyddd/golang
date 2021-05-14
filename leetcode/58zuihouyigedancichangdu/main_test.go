package main

import "testing"

func TestLastWord(t *testing.T) {
	m := make(map[string]int)
	m["Hello, world"] = 5
	m["     "] = 0
	m["  "] = 0
	m[" "] = 0
	m[""] = 0
	m["helloworld"] = 10
	m["a"] = 1

	// 注意下面这种测试用例
	m["hello   "] = 5

	for k, v := range m {
		res := lengthOfLastWord(k)
		if res != v {
			t.Errorf("%v: expect %v, but %v", k, v, res)
		}
	}
}
