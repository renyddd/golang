package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(21412))
}

func isPalindrome(x int) bool {
	// 负数与个位是零的都不可能是回文数
	if g := x % 10; g == 0 && x != 0 || x < 0 {
		return false
	}

	var rd, rev int

	// 直接这样反转会导致溢出，若只反转一半又省空间又不会溢出
	for rev < x {
		rd = x % 10
		rev = rev*10 + rd
		x = x / 10
	}

	return x == rev || x == rev/10
}
