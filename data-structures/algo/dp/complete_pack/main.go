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

func CompletePackBackTrack(V, W []int, Vbag int) int {
	thingsNum := len(V)
	maxValue := 0
	// 注意三个值才对应一种状态
	existMap := make(map[[3]int]bool, 0)

	var dfs func(nowValue, nowVolume, thingsIndex int)

	dfs = func(nowValue, nowVolume, thingsIndex int) {
		if thingsIndex >= thingsNum || nowVolume > Vbag {
			return
		}
		if nowValue > maxValue {
			maxValue = nowValue
		}
		pair := [3]int{nowValue, nowVolume, thingsIndex}
		if v, _ := existMap[pair]; v {
			return
		}
		existMap[pair] = true

		log.Println(pair)

		// 选了当前物品后，进行判断下一号物品
		dfs(nowValue+W[thingsIndex], nowVolume+V[thingsIndex], thingsIndex+1)
		// 再选一次当前物品进行判断
		dfs(nowValue+W[thingsIndex], nowVolume+V[thingsIndex], thingsIndex)
		// 跳过当前物品
		dfs(nowValue, nowVolume, thingsIndex+1)
	}

	dfs(0, 0, 0)

	return maxValue
}
