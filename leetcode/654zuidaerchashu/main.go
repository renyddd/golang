package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
		// 分类长度为 0 和 1 的情况
	} else if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	} else {
		maxIndex := findMax(nums)
		root := &TreeNode{nums[maxIndex], nil, nil}

		root.Left = constructMaximumBinaryTree(nums[:maxIndex])
		// 此处索引 maxIndex+1 是为了空出区间中间值（即是最大值）
		root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

		return root
	}
}

// findMax 假设传入此的长度都不为零
func findMax(nums []int) int {
	maxIndex := 0
	for i := range nums {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
