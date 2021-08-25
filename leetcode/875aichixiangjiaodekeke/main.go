package main

/* package
珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。

珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。



示例 1：

输入: piles = [3,6,7,11], H = 8
输出: 4
示例 2：

输入: piles = [30,11,23,4,20], H = 5
输出: 30
示例 3：

输入: piles = [30,11,23,4,20], H = 6
输出: 23


提示：

1 <= piles.length <= 10^4
piles.length <= H <= 10^9
1 <= piles[i] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/koko-eating-bananas
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minEatingSpeed(piles []int, h int) int {
	left, right := 0, 1000000000+1

	for left < right {
		mid := left + (right-left)/2
		if mid == 0 { //  提前处理发生除以 0 的情况
			// 当 h，即 x, f(x) 中的 x 取值为 0 时，代表着一种不可能有结果的情况：
			// 即吃香蕉的速度为零，这是对应的取值本应该为无穷大，也可以表示为，此种情况需要向右去逼近
			left = mid + 1
		} else if f(piles, mid) > h {
			// 这里的是否 +1，需要理解清楚循环中是否有等于号，即该范围是左闭右开还是双闭区间
			left = mid + 1
		} else if f(piles, mid) < h {
			// 循环中的右区间本身就为右闭，新词使 right = mid 不会致使重复判断的发生
			right = mid
		} else if f(piles, mid) == h {
			// 持续向左逼近
			right = mid
		}
	}

	return left
}

func f(piles []int, x int) int {
	hours := 0
	for _, v := range piles {
		hours += v / x
		if v%x != 0 {
			hours++
		}
	}

	return hours
}
