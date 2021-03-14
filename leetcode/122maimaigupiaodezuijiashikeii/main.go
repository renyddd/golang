package main

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/solution/tan-xin-suan-fa-by-liweiwei1419-2/
// maxProfit01 使用回朔算法（暴力解法）
func maxProfit01(prices []int) int {
	res := 0
	var dfs func(prics []int, index, lenn, status, profit int)

	dfs = func(prics []int, index, lenn, status, profit int) {
		if index == lenn {
			res = max(profit, res)
			return
		}

		dfs(prices, index+1, lenn, status, profit)

		// 对应于上述情况的分之情况
		// 从持有转为未持有，从未持有转为持有
		if status == 0 {
			dfs(prices, index+1, lenn, 1, profit-prices[index])
		} else {
			dfs(prices, index+1, lenn, 0, profit+prices[index])
		}
	}

	dfs(prices, 0, len(prices), 0, 0)

	return res
}

// maxProfit 由于交易次数不受限制，即可收集每次上坡
func maxProfit(prices []int) int {
	res := 0

	for i := 0; i < len(prices)-1; i++ {
		if v := prices[i+1] - prices[i]; v > 0 {
			res += v
		}
	}

	return res
}

// maxProfit2 使用贪心算法
// 仔细理解 h1 到 h5 的绝对高度之差，就是分别每个间隔一区间的绝对高度差之和
// 因此运用贪心算法，每次只统计每个"1间隔"做出正贡献的高度差，极为最优解
func maxProfit2(prices []int) int {
	res := 0

	for i := 0; i < len(prices)-1; i++ {
		res += max(0, prices[i+1]-prices[i])
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// maxProfit3 动态规划解法中关注的股票买后某一日的"持有与否的状态"而非具体的"买卖操作"
// dp0 代表着当日末未持有股票的资金状态，即当日后未持有股票的资金
// dp1 代表着当日末持有股票的资金状态
// 因此新一轮的 dp0 状态，可由前一天的 dp0 和前一天 dp1 选择卖出的收益，的最优解确定
// 新一天 dp1 状态值，可有前一天 dp1 和前一天 dp0 买入后负债，的最优解确定
func maxProfit3(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	// 最终一日的收入状态，肯定是选择卖出的收益更高，也就是不持有股票状态
	return dp[n-1][0]
}

// maxProfit4
func maxProfit4(prices []int) int {
	dp0, dp1 := 0, -prices[0]

	for i := 1; i < len(prices); i++ {
		newdp0 := max(dp0, dp1+prices[i])
		newdp1 := max(dp1, dp0-prices[i])
		dp0, dp1 = newdp0, newdp1
	}

	return dp0
}
