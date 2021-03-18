package main

import "log"

// ref: https://www.bilibili.com/video/BV1nE411D759?p=2
// 完全背包每种物品都能用无限次

// CompletePack
func CompletePack(volumes, weights []int, Vbag int) int {
	n := len(volumes)
	dp := make([]int, Vbag+1)

	for i := 1; i < n+1; i++ {
		index := i - 1
		for j := volumes[index]; j < Vbag+1; j++ {
			dp[j] = max(dp[j], dp[j-volumes[index]]+weights[index])
		}

		log.Println(dp)
	}

	return dp[Vbag]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
