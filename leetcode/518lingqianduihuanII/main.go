package main

import "log"

/*
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
示例 1:

输入: amount = 5, coins = [1, 2, 5]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change-2
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// ref: https://www.bilibili.com/video/BV1TE411d7eF?t=52

// change
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i < n+1; i++ {
		index := i - 1
		for j := coins[index]; j < amount+1; j++ {
			dp[j] += dp[j-coins[index]]
		}

		log.Println(dp)
	}

	return dp[amount]
}
