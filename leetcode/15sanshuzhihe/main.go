package main

import "sort"

/*给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：
输入：nums = []
输出：[]
示例 3：
输入：nums = [0]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// threeSum1 . 方法超时
func threeSum1(nums []int) [][]int {
	res := make([][]int, 0)

	// 仅仅是基于乱序的数组来跑三重循环，会导致重复元素的出现
	// 通过事先排序，与进阶遍历时使用不相同的元素即可避免组合的重复

	/*
		输入
		[-1,-1,0,0,1,1]
		输出
		[[-1,0,1],[-1,0,1],[-1,0,1],[-1,0,1],[-1,0,1],[-1,0,1],[-1,0,1],[-1,0,1]]
		预期结果
		[[-1,0,1]]
	*/

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 仅允许在首次时出现重复元素
		if i == 0 || nums[i] != nums[i-1] {
			for j := i + 1; j < len(nums); j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					for k := j + 1; k < len(nums); k++ {
						// 在最后判断时依旧需要判断是否重复
						if k == j+1 || nums[k] != nums[k-1] {
							if vi, vj, vk := nums[i], nums[j], nums[k]; vi+vj+vk == 0 {
								res = append(res, []int{vi, vj, vk})
							}
						}
					}
				}
			}
		}
	}

	return res
}

// threeSum2 第三次的循环中采用双指针，但暂时还未搞清楚双指针的时间复杂度为何较低？
// 单纯的想用 eg: -3,-3,-2,-2,-1,-1,-1,-1,1,2,2,2,3,3,4,5,6,7,8 在纸上手动模拟双指针的遍历顺序，就会发现
// 在第一次确定 standard 之后，确实只需要遍历一次即可；
// 双指针对此的优化，一定是建立在"有序"之上的；即双指针通过有序的性质与夹逼的过程，有效的利用了截止点，
// 正是该有效的截止，才使其时间复杂度低于了顺序二重遍历。
// 参考题解：https://leetcode-cn.com/problems/3sum/solution/0015san-shu-zhi-he-by-jasonchiucc-a-fp9n/
func threeSum2(nums []int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)

	// 每次循环都以最左边为标准
	for standard := 0; standard < len(nums); standard++ {
		// 当最小值已大于 0 时，则无需继续
		if nums[standard] > 0 {
			break
		}
		// 避免重复
		if standard > 0 && nums[standard] == nums[standard-1] {
			continue
		}

		left, right := standard+1, len(nums)-1
		for left < right {
			sum := nums[standard] + nums[left] + nums[right]
			// 三数之和过小则，为了让结果更大，则只能将 left 右移
			if sum < 0 {
				left++
			} else if sum > 0 {
				// 三数之和过大，为了让结果更小，则只能将 right 左移
				right--
			} else {
				// 命中
				res = append(res, []int{nums[standard], nums[left], nums[right]})

				// 结果存有唯一性，故可同时向内进行
				left++
				right--
				// 分别搬移至不相同的元素
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return res
}
