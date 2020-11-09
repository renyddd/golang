package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")
}

// 如下是一个合格的二分查找，二本题的意思是要找出当数字不在时的合适索引
func bSearch(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n - 1
	var mid int
	for l <= r {
		mid = (r -l) >> 1 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n - 1
	var mid int
	for l <= r {
		mid = (r -l) >> 1 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// 好好画图理解为什么最后返回的是 l
	return l
}

func searchInsert2(nums []int, target int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return i
}