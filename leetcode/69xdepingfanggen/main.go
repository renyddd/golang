package main

//ref https://leetcode-cn.com/problems/sqrtx/

// mySqrt binary search
func mySqrt(x int) int {
	var res int = -1
	low, high := 0, x

	for low <= high {
		mid := low + ((high - low) >> 1)

		if mid*mid > x {
			high = mid - 1
		} else {
			res = mid
			low = mid + 1
		}
	}

	return res
}

// mySqrt2 顺序查找
func mySqrt2(x int) int {
	i := 0
	for ; i*i < x; i++ {
	}

	if i*i > x {
		return i - 1
	}

	return i
}
