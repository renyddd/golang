### 解题思路
这题首先想的是，如何设计 dp？
若采用归纳的方式，则应该思考的是：我现在已经有了 dp[0...i-1]，那么我该如何推导出 dp[i] 的值呢？

开始定义 dp 数组的含义：dp[i] 表示，以 nums[i] 结尾的最长递增子序列个数。

以数组 nums [1, 4, 3, 4, 2, 5] 为例，手动填充一次 dp 数组：
![WechatIMG178.jpeg](https://pic.leetcode-cn.com/1631027720-VQAMhm-WechatIMG178.jpeg)

如上图所示，以末尾的 nums[5] = 5 对应的 dp 取值为例：
1. 当取值到 nums[5] 时，你需要考虑所有 nums[0...4] 中比 5 小的，因为它们组成新最长递增子序列的长度为 dp[j]+1;
2. 比较更大值更新 dp[i]。

最后返回结果取 dp 数组中的最大值即可。

### 代码

```golang
import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	// dp[i] 对应着以 nums[i] 结尾的最长递增子序列的长度lengthOfLIS

	for i, _ := range dp {
		dp[i] = 1 // 最小长度为 1
	}

	for i, _ := range nums {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 只有在待更新的元素值大于前面的元素时，才有更长的可能性
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 1
	for _, v := range dp {
		if v > res {
			res = v
		}
	}

    fmt.Println(dp)
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
```