package main

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	// dp[i] 对应着以 nums[i] 结尾的最长递增子序列的长度lengthOfLIS

	for i, _ := range dp {
		dp[i] = 1 // 最小长度为 1
	}

	for i, _ := range nums {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 只有在待更新的元素值大于前面的元素时，才有更长的可能性
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 1
	for _, v := range dp {
		if v > res {
			res = v
		}
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
