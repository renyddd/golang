package main

import "log"

// ref: https://www.bilibili.com/video/BV1LE411Z7wn
// 将 0～n 个物体放到背包中，求所装的最大价值
// eg： 物品号  体积Vi  权重Wi
//        1      2     5
//        2      4     6
//        3      3     7

// ZeroOnePack
// @volumes 物品的价值
// @weights 物品的权重
// @Vbag 背包的最大空间
func ZeroOnePack(volumes, weights []int, Vbag int) int {
	n := len(volumes) + 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, Vbag+1)
	}

	// 第一层全使用初始化的 0 值
	// 第 i 层，就表示已处理到第 i 号物品
	for i := 1; i < n; i++ {
		for j := Vbag; j >= 0; j-- {
			index := i - 1
			if j-volumes[index] >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-volumes[index]]+weights[index])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	log.Println(dp)
	return dp[len(volumes)][Vbag]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// ZeroOnePackCommented 完全注释版
// @volums 物品的价值
// @weights 物品的权重
// @Vbag 背包的最大空间
// a[i][j] = WEIGHT 中 i 表示已处理的物品个数，也可理解为是正在处理的物品号
// j 表示背包当前的体积；WEIGHT 表示在此刻 i，j 状态下的权重值
func ZeroOnePackCommented(volumes, weights []int, bagMaxVolum int) int {
	// volumes 与 weights 的个数必须要相同，代表着待处理物品的个数
	// 因为我们所用的 dp 二维数组中，每一层的号直接代表着物品的号码
	// 而且第零层意义也即是：目前处理到了第 0 号物品，无论当前背包所剩空间还有多少，我的权重都是 0
	n := len(volumes) + 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		// dp 每一层切片长度选择为背包最大容量 +1 的原因与上 n 相同
		// 都是想用"背包当前体积数"来直接表示索引
		dp[i] = make([]int, bagMaxVolum+1)
	}

	for i := 1; i < n; i++ {
		// 从后往前遍历以确保每次在状态判断时都会用到之前的状态
		for bagLeftVolume := bagMaxVolum; bagLeftVolume >= 0; bagLeftVolume-- {
			// index 表示在 values，weights 切片中的索引值
			index := i - 1
			// 避免 v 计算出负数而导致切片索引 panic
			if v := bagLeftVolume - volumes[index]; v >= 0 {
				// 状态转移方程
				// 新状态等于，不装第 i 号物品，与装 i 号物品的最大值
				// v 是装了第 i 号物品后的空间，由当前所剩空前减去 i 号物品的空间大小
				dp[i][bagLeftVolume] = max(dp[i-1][bagLeftVolume], dp[i-1][v]+weights[index])
			} else {
				// 若装不下，则直接复制上面的状态值
				dp[i][bagLeftVolume] = dp[i-1][bagLeftVolume]
			}
		}
	}

	log.Println(dp)

	for i := bagMaxVolum; i >= 0; i-- {
		if v := dp[len(volumes)][i]; v != 0 {
			return v
		}
	}

	return -1
}

// ZeroOnePackOneSlice 使用一维数组优化空间复杂度
func ZeroOnePackOneSlice(volumes, weights []int, Vbag int) int {
	n := len(volumes)
	dp := make([]int, Vbag+1)

	for i := 1; i < n+1; i++ {
		// 如果此时选择从前向后更新，则会导致在计算更大的体积时用到了新的状态
		// 而从后向前更新，则会保证在计算新状态时，用于更新的小体积值总是来自旧状态的
		//for j := Vbag; j >= 0; j-- {
		//	index := i - 1
		//	if delta := j-volumes[index]; delta >= 0 {
		//		dp[j] = max(dp[j], dp[delta]+weights[index])
		//	}
		//	// 无需再做剩余的拷贝
		//}

		index := i - 1
		// 当剩余体积比当前物品还小时，便无需再判断
		for j := Vbag; j >= volumes[index]; j-- {
			dp[j] = max(dp[j], dp[j-volumes[index]]+weights[index])
		}

		log.Println(dp)
	}

	return dp[Vbag]
}

/*
 29:00 左右
 会涉及到背包是否有要求必须要装满，注意这种条件下的初始化策略
 */