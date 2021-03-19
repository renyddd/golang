package main

/*
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum+= v
	}
	// 总和是奇数，或者个数小于两个的天然错误
	if sum % 2 > 0 || len(nums) < 2 {
		return false
	}
	// amount 与之前题目相同，代表着背包的总大小，所能装的物品总量
	amount := sum / 2
	dp := make([]bool, amount+1)
	// 特殊初始化第一个元素
	dp[0] = true

	for i := 1; i < len(nums) + 1; i++ {
		index := i - 1
		// 01 背包从后向前扫描
		for j := amount; j >= nums[index]; j-- {
			// 如果 dp[j-nums[index]] 状态能成立，则当前状态必然成立
			//
			dp[j] = dp[j] || dp[j-nums[index]]
		}
	}

	//log.Println(dp)

	return dp[amount]
}
