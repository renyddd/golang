### 解题思路
首先使用暴力递归的递归数，如下图所示：
![coins-dp.jpg](https://pic.leetcode-cn.com/1630945063-NXZELu-coins-dp.jpg)

辅助函数 find，用来根据**「目标金额 amount」**与**「所用硬币个数 count」**来持续更新最小值，由于枝数过多必然会导致超时：

```golang
func coinChange(coins []int, amount int) int {
	res := math.MaxInt64

	var find func(amount, count int)
	find = func(amount, count int) {
		if amount < 0 {
			return
		}
		if amount == 0 {
			res = min(res, count)
		}

		for _, v := range coins {
			find(amount-v, count+1)
		}
	}

	find(amount, 0)

	if res == math.MaxInt64 {
		return -1
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
```

当我们仔细思考整个递归过程，与递归数的时候，会有如下发现：
1. 我们不关心硬币的排列顺序，即不关心以什么顺序进行递归；
2. 我们最终只关心这个结果的数值；
3. 可通过 memo 来进行记忆化搜索。

重复部分如图所示：
![WechatIMG176.jpeg](https://pic.leetcode-cn.com/1630946449-DsDgQe-WechatIMG176.jpeg)

有了上面的思考，就应该清楚了 memo 代表：能组成「目标金额值」的「最小硬币个数」；
现在的问题是，把 memo 放在哪？

```golang
func coinChange(coins []int, amount int) int {
	memo := make([]int, amount) // memo 代表目标和所需的最小硬币个数

	var find func(amount int) int 
    // find 用于返回能组成「目标金额值 amount」的「最小硬币个数」
	find = func(amount int) int {
		if amount < 0 {
			return -1
		}
		if amount == 0 {
			return 0
		}

		if v := memo[amount-1]; v != 0 { // 不等于 0 即为已记录过该总和，或认其无效即负一
			return v
		}

		tmpMin := math.MaxInt64
		for _, v := range coins {
			res := find(amount - v)
			if res >= 0 && res < tmpMin {
				tmpMin = res + 1
			}
		}

		if tmpMin == math.MaxInt64 {
			tmpMin = -1
		}
		memo[amount-1] = tmpMin

		return tmpMin
	}

	return find(amount)
}
```