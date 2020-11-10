// Learn binary search
package bs

// 如下是一个合格的二分查找，二本题的意思是要找出当数字不在时的合适索引
func bSearch(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	var mid int
	for l <= r {
		mid = (r-l)>>1 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// bSearchFirst 有序集合存在重复数据 查找第一个等于给定值的数据
func bSearchFirst(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := ((r - l) >> 1) + l
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else if m == 0 || nums[m-1] != target {
			return m
		} else {
			r = m - 1
		}
	}
	return -1
}

// bSearchLast 有序集合存在重复数据 查找最后一个等于给定值的数据
func bSearchLast(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := ((r - l) >> 1) + l
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else if m == n-1 || nums[m+1] != target {
			return m
		} else {
			l = m + 1
		}
	}
	return -1
}

// bSearchFBigger 有序集合查找第一个大于等于给定值的数据
func bSearchFBigger(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := ((r - l) >> 1) + l
		if nums[m] >= target {
			if m == 0 || nums[m-1] < target {
				return m
			} else {
				r = m - 1
			}
		} else {
			l = m + 1
		}
	}
	return -1
}

// bSearchLSmaller 有序集合查找最后一个小于给定值的数据
func bSearchLSmaller(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n - 1
	for l <= r {
		m := ((r - l) >> 1) + l
		if nums[m] > target {
			r = m - 1
		} else if m == n-1 || nums[m+1] > target {
			return m
		} else {
			l = m + 1
		}
	}
	return -1
}



// leetcode 题目要求的解答
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。无重复。
func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	var mid int
	for l <= r {
		mid = (r-l)>>1 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// 好好画图理解为什么最后返回的是 l
	return l
}

func searchInsert2(nums []int, target int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return i
}
