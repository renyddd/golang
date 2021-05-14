package main

import "math"

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	maxSum, thisSum := math.MinInt64, 0
	for i := 0; i < len(nums); i++ {
		thisSum += nums[i]
		if thisSum > maxSum {
			maxSum = thisSum
		} else if thisSum < 0 {
			thisSum = 0
		}
	}
	return maxSum
}
