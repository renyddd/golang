package main

import (
	"fmt"
	"sort"
)

func main() {
	merge([]int{1,2,4,6,0,0,0}, 4, []int{3,4,5,7}, 0)
}
/*
	输入：
	nums1 = [1,2,3,0,0,0], m = 3
	nums2 = [2,5,6],       n = 3

	输出：[1,2,2,3,5,6]
 */

func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums1 = append(nums1[:m], nums2...)
	fmt.Println(nums1)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}

	fmt.Println(nums1)
}

func mergeSortPackage(nums1 []int, m int, nums2 []int, n int)  {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

/**
 *
 * @param str string字符串 the string
 * @return string字符串
 */