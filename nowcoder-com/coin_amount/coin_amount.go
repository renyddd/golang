package main

import "fmt"

// 硬币问题：
// 输入不同面额的硬币与一固定金额
// 返回能够组成改金额的最小硬币数量，无法则返回 -1
// 贪婪、动态规划

func main() {
	var coins = []int{2, 3, 7, 10}

	use := do_Greed_Find(coins, 12345433245)
	fmt.Println(use)
}

// 首先使用贪婪
// 一直从总额中减去最大硬币的数值
// 直到剩下的金额直小于最大金额则尝试第二大面额
func do_Greed_Find(coins []int, amount int) int {
	total := 0

	for i := len(coins) - 1; i >= 0; i-- {
		total += amount / coins[i]
		amount = amount % coins[i]
		if amount == 0 {
			return total
		}
	}
	return -1
}

// 此种解法的限制为硬币为递增顺序
// coinds=1,5,20,25；amount=40 时无法获取到最优解
//
