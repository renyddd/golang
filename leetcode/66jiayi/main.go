package main

import (
	"fmt"
)

func main() {
	n := []int{9,9,9}
	n=plusOneRecursion(n)
	fmt.Println(n)
}

func plusOneRecursionIn(digits []int) []int {
	l := len(digits)
	var recursion func(int, *[]int)
	recursion = func(index int, d *[]int) {
		if (*d)[index] != 9 {
			(*d)[index]++
			return
		} else if index-1 >= 0 {
			(*d)[index] = 0
			recursion(index-1, d)
		} else {
			// 处理以 9 开头的情况
			(*d)[0] = 1
			*d = append(*d, 0)
			return
		}
	}

	if l == 0 {
		return []int{1}
	} else {
		recursion(l-1, &digits)
	}
	return digits
}

// plusOneRecursion 此处的递归无外乎就是在考虑两种情况
// - 不是 9 就直接加一并返回
// - 判断前一位是不是 9 ？否：当前变零前一位加一；是：下
// - - 同时还要判断前一位是否是否到头了？是：则新加一位 1；否：递归判断
func plusOneRecursion(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return []int{1}
	} else {
		Recursion(l-1, &digits)
	}
	return digits
}

func Recursion(index int, d *[]int) {
	if (*d)[index] != 9 {
		(*d)[index]++
		return
	} else if index-1 >= 0 {
		(*d)[index] = 0
		Recursion(index-1, d)
	} else {
		// 处理以 9 开头的情况
		(*d)[0] = 1
		*d = append(*d, 0)
		return
	}
}
/*
	{1,2,3},
	{1,1,9},
	{9},
	{9,9,9},
	{0},
	{},
 */

func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return []int{1}
	}

	for i := l-1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}