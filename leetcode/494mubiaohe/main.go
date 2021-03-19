package main

/*
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 - 中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

示例：

输入：nums: [1, 1, 1, 1, 1], S: 3
输出：5
解释：
-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

一共有5种方法让最终目标和为3。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 写的确实有问题
// findTargetSumWays 同样是采用动态规划的方法，关注点是：
// 1 如何初始化首行元素
// 2 如何找到递推方程
// 3 画状态转移数组
// 4 写代码
//func findTargetSumWaysFirstTry(nums []int, S int) int {
//	//sum := 0
//	//for _, v := range nums {
//	//	sum += v
//	//}
//	//// 题目比较坑，会给出不可能出现的 S 值，将会导致索引越界
//	//if abs(S) > abs(sum) {
//	//	return 0
//	//}
//	//
//	//// amount 代表着 nums 中全部使用减号或全部使用加号后的值上下限，也就是所会出现的值的数量
//	//// 后面的加一代表着例如 [-5,5] 中的零值
//	//amount := 2*sum + 1
//	//dp := make([][]int, len(nums)+1)
//	//for i := 0; i < len(nums)+1; i++ {
//	//	// 这里 amount 的加一，表示想用计算值来直接作为可索引号，因此需要向后直接偏移 1，让出 0
//	//	dp[i] = make([]int, amount+1)
//	//}
//	//// 如何选择初始化行？这里依旧保留 0 行的目的是方便行的索引
//	//// 如何确定索引的方法？因为数组中无法使用负值来索引
//	//// 如果用计算的结果值来作为索引的话，则有
//	//// dp[1][-1]=-1,  dp[1][1]=1
//	///* 例如如下 dp 数组（其中 -x 表示占位，不会出现的值）
//	//会出现的值：-x,-5,-4,-3,-2,-1,0,1,2,3,4, 5
//	//数组中索引：0,  1, 2, 3, 4, 5,6,7,8,9,10,11
//	//----
//	//数组中索引 - offset = 可能的计算值
//	//eg: 11 - offset = 5
//	//*/
//	//offset := amount - sum
//	//dp[1][offset+nums[0]], dp[1][offset-nums[0]] = 1, 1
//	//
//	//// 开始动态规划
//	//for i := 2; i < len(nums)+1; i++ {
//	//	// 特殊赋值最右侧元素，因为 amount+1 会导致索引越界
//	//	nindex := i - 1
//	//	for j := amount; j > nums[nindex]; j-- {
//	//		// dp[i][j] = dp[i-1][j-nums[nindex] + dp[i-1][j+nums[nidex]
//	//		var left, right int
//	//		if j-nums[nindex] >= 0 {
//	//			left = dp[i-1][j-nums[nindex]]
//	//		}
//	//		// 避免右边界索引溢出
//	//		if j+nums[nindex] <= amount {
//	//			right = dp[i-1][j+nums[nindex]]
//	//		}
//	//		dp[i][j] = left + right
//	//	}
//	//}
//	//
//	//log.Println(dp)
//	//
//	//log.Println(dp[len(nums)][S+offset])
//	//
//	//return dp[len(nums)][S+offset]
//}

func findTargetSumWaysSecondTry(nums []int, S int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// 提前处理不可能满足的情况
	if abs(S) > abs(sum) {
		return 0
	}
	// [-sum, sum]，其中 +1 代表算上 0 的个数
	maxIndex := 2*sum + 1
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, maxIndex)
	}
	// sum 的索引代表着 0 数值的情况
	dp[0][sum] = 1

	for i := 1; i < len(nums)+1; i++ {
		// 注意 j 必要可取值至 0，否则将会永远少计算左边界的情况
		// 因为 dp[i][j] 也依赖着上层右边的状态
		for j := maxIndex - 1; j >= 0; j-- {
			rightDp, leftDp := 0, 0
			// 处理索引越界的情况
			if j+nums[i-1] < maxIndex {
				rightDp = dp[i-1][j+nums[i-1]]
			}
			if j-nums[i-1] >= 0 {
				leftDp = dp[i-1][j-nums[i-1]]
			}
			dp[i][j] = leftDp + rightDp
		}
	}

	// sum 就是从可能的计算值到数组索引的偏移量
	return dp[len(nums)][S+sum]
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// findTargetSumWaysEnumerate 采用暴力破解
func findTargetSumWaysEnumerate(nums []int, S int) int {
	res := 0
	length := len(nums)
	var dofind func(int, int)
	dofind = func(index, tmpSum int) {
		if index == length {
			if tmpSum == S {
				res++
			}
		} else {
			dofind(index+1, tmpSum+nums[index])
			dofind(index+1, tmpSum-nums[index])
		}
	}

	dofind(0, 0)

	return res
}
