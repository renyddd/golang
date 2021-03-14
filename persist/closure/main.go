package main

import (
	"fmt"
)

// todo 观测闭包中父函数变量的垃圾回收机制

// Generator 闭包是如何保存赴父函数的值？
func Generator(tail string) func(string) string {
	return func(s string) string {
		return s + tail
	}
}

// childFunc 闭包的判断形式为何？是由父函数定义中的返回类型而定的吗？
// 还是由 return 语句后紧跟的函数定义而判定的？
func childFunc(s, tail  string) string {
	return s + tail
}

func AddTail() func(string, string) string {
	return childFunc
}

func AddTailWithoutCall(s, tail string) string {
	return childFunc(s, tail)
}

func main() {
	f := Generator("tail1")
	fmt.Println(f("hello, "))

	fmt.Println(AddTail()("hello, ", "tail2"))
	fmt.Println(AddTailWithoutCall("hello, ", "tail3"))
}