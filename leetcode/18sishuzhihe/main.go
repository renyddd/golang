package main

import "sort"

/*
给定一个包含n 个整数的数nums和一个目标值target，判断nums中是否存在四个元素 a，b，c和 d，使得a + b + c + d的值与target相等？找出所有满足条件且不重复的四元组。
注意：答案中不可以包含重复的四元组。

示例 1：
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：
输入：nums = [], target = 0
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)

	for s1 := 0; s1 < len(nums); s1++ {
		if s1 > 0 && nums[s1] == nums[s1-1] {
			continue
		}
		for s2 := s1+1; s2 < len(nums); s2++ {
			if s2 != s1+1 && nums[s2] == nums[s2-1] {
				continue
			}

			left, right := s2+1, len(nums)-1
			for left < right {
				sum := nums[s1] + nums[s2] + nums[left] + nums[right]
				if sum < target {
					left++
				} else  if sum > target {
					right--
				} else {
					res = append(res, []int{nums[s1], nums[s2], nums[left], nums[right]})

					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
			}
		}
	}

	return res
}