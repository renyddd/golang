package main

/*
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :

输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
说明 :

数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 首先得出前序和，再根据结果数组寻找相差为 k 的次数

func subarraySum(nums []int, k int) int {
	tmp := make([]int, len(nums)+1) // 前序和数组
	cnt := 0

	for i, v := range nums {
		tmp[i+1] = tmp[i] + v
	}

	for i := 0; i < len(tmp)-1; i++ {
		for j := i + 1; j < len(tmp); j++ {
			v := tmp[j] - tmp[i]

			if v == k {
				cnt++
			}
		}
	}

	return cnt
}
