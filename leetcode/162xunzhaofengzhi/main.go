package main

import "math"

func findPeakElement(nums []int) int {
	max := math.MinInt64
	maxIndex := 0

	for i, v := range nums {
		if v > max {
			max = v
			maxIndex = i
		}
	}

	return maxIndex
}
