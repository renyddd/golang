package main

import "log"

// ref: https://www.bilibili.com/video/BV1LE411Z7wn
// 将 0～n 个物体放到背包中，求所装的最大价值
// eg： 物品号  体积Vi  权重Wi
//        1      2     5
//        2      4     6
//        3      3     7

// ZeroOnePack
// @volums 物品的价值
// @weights 物品的权重
// @Vbag 背包的最大空间
func ZeroOnePack(volums, weights []int, Vbag int) int {
	n := len(volums) + 1
	dp := make([][]int, 0)
	for i := 0; i < n; i++ {
		tmp := make([]int, Vbag+1)
		dp = append(dp, tmp)
	}

	// 第一层全使用初始化的 0 值
	// 第 i 层，就表示已处理到第 i 号物品
	for i := 1; i < n; i++ {
		for j := Vbag; j >= 0; j-- {
			index := i-1
			if j-volums[index] >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-volums[index]]+weights[index])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	log.Println(dp)
	return dp[len(volums)][Vbag]
}

func max(i, j int) int {
	if i > j { return i }
	return j
}