package main

import "fmt"

func main() {
	fmt.Println(climbStairs(44))
	fmt.Println(DataOut(44))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return climbStairs(n-1) + climbStairs(n-2)
	}
}

func DataOut(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
