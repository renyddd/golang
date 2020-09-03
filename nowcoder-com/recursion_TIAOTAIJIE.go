package main

import "fmt"

func main() {
	res := f(5)

	fmt.Println(res)
}

// 青蛙跳台阶问题
// 一次可以跳一节台阶或者两节台阶，问总共有多少中跳法？
// 进行递归时首先从递归终止条件开始想
// 再来描述递推公式
func f(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return f(n-1) + f(n-2)
	}
}
