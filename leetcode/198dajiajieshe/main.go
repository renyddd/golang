package main

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 1：

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
	偷窃到的最高金额 = 1 + 3 = 4 。
示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
    偷窃到的最高金额 = 2 + 9 + 1 = 12 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// rob
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	// 不用死记 01 背包从后向前，完全背后从前向后
	// 只用理解：计算下一个状态值的时候，是否需要复用前面状态的新最佳值
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(dp)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	prew, cur := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		prew, cur = cur, max(cur, prew+nums[i])
	}

	return cur
}

// robDfs 如采用 dfs 完成时需要额外考虑什么
func robDfs(nums []int) int {
	res := 0
	n := len(nums)

	var dfs func(i, sum int, toRob bool)
	dfs = func(i, sum int, toRob bool) {
		// 判断最值要放在返回之前
		// 否则少判断最后一种情况
		// 避免后续索引越界
		if sum > res {
			res = sum
		}

		if i >= n {
			return
		}

		if toRob {
			// 如果当前选择了打劫，则下一个必须不能打
			dfs(i+1, sum+nums[i], false)
		} else {
			// 如果未打劫这一家，则无需更新 sum 值
			// 下一家的打劫与否的有两种选择
			dfs(i+1, sum, true)
			dfs(i+1, sum, false)
		}
	}

	dfs(-1, 0, false)

	return res
}

func robDfs2(nums []int) int {
	res := 0
	n := len(nums)
	exist := make(map[[3]int]bool, 0)

	var dfs func(i, sum int, toRob bool)
	dfs = func(i, sum int, toRob bool) {
		if sum > res {
			res = sum
		}
		if i >= n {
			return
		}
		ntorob := 0
		if toRob {
			ntorob = 1
		}
		pair := [3]int{i, sum, ntorob}
		if _, ok := exist[pair]; ok {
			return
		}
		exist[pair] = true

		if toRob {
			dfs(i+1, sum+nums[i], false)
		} else {
			dfs(i+1, sum, true)
			dfs(i+1, sum, false)
		}
	}

	dfs(-1, 0, false)

	return res
}
