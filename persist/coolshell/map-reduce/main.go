package main

import (
	"fmt"
	"strings"
)

// https://coolshell.cn/articles/21164.html

func MapStr2Str(arr []string, fn func(s string) string) []string {
	var newArr []string
	for _, it := range arr {
		newArr = append(newArr, fn(it))
	}
	return newArr
}

func MapStr2Int(arr []string, fn func(s string) int) []int {
	var newArr []int
	for _, it := range arr {
		newArr = append(newArr, fn(it))
	}
	return newArr
}

var list = []string{"Hello", "world", "!", "How", "are", "you"}

func UseMap() {
	//x := MapStr2Str(list, strings.ToUpper)
	x := MapStr2Str(list, func(s string) string {
		return strings.ToUpper(s)
	})
	y := MapStr2Int(list, func(s string) int {
		return len(s)
	})
	fmt.Println(x)
	fmt.Println(y)
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

func UseReduce() {
	x := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Println(x)
}

func Filter(arr []int, fn func(i int) bool) []int {
	var newArr []int
	for _, it := range arr {
		if fn(it) {
			newArr = append(newArr, it)
		}
	}
	return newArr
}

func UseFilter() {
	nums := []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11}
	x := Filter(nums, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(x)
}

func main() {
	UseMap()
	UseReduce()
	UseFilter()
}
