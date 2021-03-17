package main

import "math"

//ref: https://time.geekbang.org/column/article/74788
// 0-1 背包问题
// 对于一组不同重量、不可分割的物品，哦我买呢需要选择一些装入背包，在满足背包最大重量限制的前提下，背包中物品的总重量的最大值是多少？

// Backtracking 回朔
// @weight 表示各物品重量
// @bagMaxWeight 表示背包限制的重量
func Backtracking(weight []int, bagMaxWeight int) int {
	res := math.MinInt64 // 结果最大值
	n := len(weight)     // 物品数量

	var f func(i, carryWeight int)
	f = func(i, carryWeight int) {
		// 当携带的物品总重超过，或数量到达上限时即结束
		if carryWeight >= bagMaxWeight || i >= n {
			if carryWeight > res {
				res = carryWeight
			}
			return
		}
		// 装第 i 号物品
		f(i+1, carryWeight)

		if carryWeight+weight[i] <= bagMaxWeight {
			// 不选择装第 i 号物品
			f(i+1, carryWeight+weight[i])
		}
	}

	// 从 0 号物品 0 重量开始递归调用
	f(0, 0)

	return res
}

// Backtracking2 采用备忘录模式来减少
func Backtracking2(weight []int, bagMaxWeight int) int {
	res := math.MinInt64
	n := len(weight)
	// 二维数组长度应比午评总重量稍长 以避免 index 超出

	existState := make([][]bool, 0)
	for i := 0; i < n; i++ {
		tmpArr := make([]bool, bagMaxWeight+3)
		existState = append(existState, tmpArr)
	}

	var f func(i, carryWeight int)
	f = func(i, carryWeight int) {
		if carryWeight >= bagMaxWeight || i >= n {
			if carryWeight > res {
				res = carryWeight
			}
			return
		}

		// 遇见重复状态则直接放回
		if existState[i][carryWeight] {
			return
		}
		// 此处的重复状态对应的是
		// 不论之前已经装了的物品是什么，我只关心此时进行至第 i 号物品时的总重量，是否出现过
		// 我关心的是当前*状态*对下一次选择的影响，而不用关心是如何到达这个状态的
		existState[i][carryWeight] = true

		f(i+1, carryWeight)
		if carryWeight+weight[i] <= bagMaxWeight {
			f(i+1, carryWeight+weight[i])
		}
	}

	f(0, 0)

	return res
}
