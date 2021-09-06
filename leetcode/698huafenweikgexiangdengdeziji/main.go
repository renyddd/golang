package main

import "sort"

/* 给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。

示例 1：

输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。


提示：

1 <= k <= len(nums) <= 16
0 < nums[i] < 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

// canPartitionKSubsets1 以数字为视角会超时间限制
func canPartitionKSubsets1(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%k != 0 {
		return false
	}

	buckets := make([]int, k)
	target := sum / k

	// 加上之后直接通过！！！nb
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	return backtrack1(nums, 0, buckets, target)
}

func backtrack1(nums []int, index int, buckets []int, target int) bool {
	if index == len(nums) {
		// 检查所有桶的数字之和是否都为 target
		for _, b := range buckets {
			if b != target {
				return false
			}
		}
		return true
	}

	for i, b := range buckets {
		if b+nums[index] > target {
			continue
		}

		// 做选择
		buckets[i] += nums[index]
		if backtrack1(nums, index+1, buckets, target) {
			return true
		}
		// 撤销选择
		buckets[i] -= nums[index]
	}

	return false
}
