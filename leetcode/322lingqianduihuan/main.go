package main

import "log"

//ref: https://leetcode-cn.com/problems/coin-change/a
/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
你可以认为每种硬币的数量是无限的。
*/

func coinChange(coins []int, amount int) int {
	n := len(coins) // n 中硬币也就代表着 n 中可能发生状态转移的机会
	dp := make([]int, amount+1)
	// 每个元素索引代表这每种可能的存在金额
	// 而每个元素值代表着已用的硬币数量

	for i := 0; i < amount+1; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 1; i < n+1; i++ {
		index := i - 1
		for j := coins[index]; j < amount+1; j++ {
			if dp[j-coins[index]] == -1 {
				continue
			}
			if dp[j] == -1 {
				dp[j] = dp[j-coins[index]] + 1
				continue
			}
			dp[j] = min(dp[j], dp[j-coins[index]]+1)
		}
		log.Println(dp)
	}

	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
