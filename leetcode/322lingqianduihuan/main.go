package main

import "math"

//ref: https://leetcode-cn.com/problems/coin-change/a
/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
你可以认为每种硬币的数量是无限的。
*/

// reg: https://leetcode-cn.com/problems/coin-change/solution/javadi-gui-ji-yi-hua-sou-suo-dong-tai-gui-hua-by-s/

// coinChange1 直接暴力遍历，会超时
func coinChange1(coins []int, amount int) int {
	res := math.MaxInt64

	var find func(amount, count int)
	find = func(amount, count int) {
		if amount < 0 {
			return
		}
		if amount == 0 {
			res = min(res, count)
		}

		for _, v := range coins {
			find(amount-v, count+1)
		}
	}

	find(amount, 0)

	if res == math.MaxInt64 {
		res = -1
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// coinChange2 采用记忆优化
func coinChange2(coins []int, amount int) int {
	memo := make([]int, amount) // memo 代表目标和所需的最小硬币个数

	var find func(amount int) int // 返回目标和 amount 所需要的个数
	find = func(amount int) int {
		if amount < 0 {
			return -1
		}
		if amount == 0 {
			return 0
		}

		if v := memo[amount-1]; v != 0 { // 不等于 0 即为已记录过该总和
			return v
		}

		tmpMin := math.MaxInt64
		for _, v := range coins {
			res := find(amount - v)
			if res >= 0 && res < tmpMin {
				tmpMin = res + 1
			}
		}

		if tmpMin == math.MaxInt64 {
			tmpMin = -1
		}
		memo[amount-1] = tmpMin

		return tmpMin
	}

	return find(amount)
}

// coinChange3 使用动态规划
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1) // amount 值直接作为索引
	for i, _ := range dp {
		dp[i] = amount + 1 // 默认 todo
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i > coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

/* old do not look */
// func coinChange_old(coins []int, amount int) int {
// 	n := len(coins) // n 中硬币也就代表着 n 中可能发生状态转移的机会
// 	dp := make([]int, amount+1)
// 	// 每个元素索引代表这每种可能的存在金额
// 	// 而每个元素值代表着已用的硬币数量

// 	for i := 0; i < amount+1; i++ {
// 		dp[i] = -1
// 	}
// 	dp[0] = 0

// 	for i := 1; i < n+1; i++ {
// 		index := i - 1
// 		for j := coins[index]; j < amount+1; j++ {
// 			if dp[j-coins[index]] == -1 {
// 				continue
// 			}
// 			if dp[j] == -1 {
// 				dp[j] = dp[j-coins[index]] + 1
// 				continue
// 			}
// 			dp[j] = min(dp[j], dp[j-coins[index]]+1)
// 		}
// 		log.Println(dp)
// 	}

// 	return dp[amount]
// }

// func min(i, j int) int {
// 	if i < j {
// 		return i
// 	}
// 	return j
// }
