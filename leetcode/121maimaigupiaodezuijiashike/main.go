package main

import "math"

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// maxProfit01 使用暴力破解
func maxProfit01(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if v := prices[j] - prices[i]; v > res {
				res = v
			}
		}
	}

	return res
}

// maxProfit 使用动态规划求解
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	minprice := prices[0]
	dp := make([]int, len(prices))

	for i := 1; i < len(prices); i++ {
		minprice = min(prices[i], minprice)
		dp[i] = max(dp[i-1], dp[i]-minprice)
	}

	// dp 切片是一个将最佳收益保留至最后的数组
	return dp[len(prices)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxProfit1(prices []int) int {
	minprice := math.MaxInt64
	maxprofit := 0

	for i := 0; i < len(prices); i++ {
		// 当某一天价格更小时，更新价格最小记录
		if prices[i] < minprice {
			minprice = prices[i]
		} else
		// 否则在这一天卖出，是否可以获得最大收益
		if v := prices[i] - minprice; v > maxprofit {
			// 最大收益肯定是来自于已有记录的最低点减去目前价格
			maxprofit = v
		}
	}

	return maxprofit
}
