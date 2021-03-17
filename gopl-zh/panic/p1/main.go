package main

import (
	"fmt"
)

// https://books.studygolang.com/gopl-zh/ch5/ch5-10.html
// 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

func S() {

}

func C() {

}

func NonReturn(v interface{}) (result interface{}) {
	defer func() {
		if p := recover(); p != nil {
			result = p
		}
	}()

	panic(v)
}

func main1() {
	a := NonReturn(true)
	fmt.Println(a)
}
