package main

import "math"

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Traverse(root, math.MinInt64, math.MaxInt64)
}

func Traverse(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	// 要记得添加 false 的条件
	if root.Val <= minVal || root.Val >= maxVal {
		return false
	}
	// 对于左子树的处理，需要更新其判断的最大值边界，因为其值要小于根节点
	// 对于柚子树的处理，需要更新其判断的最小值边界，因为其值要大于根节点
	return Traverse(root.Left, minVal, root.Val) && Traverse(root.Right, root.Val, maxVal)
}

/*
   排序后验证
 */

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	nums := make([]int, 0)
	var Iterate func(node *TreeNode)
	// 之一匿名函数中变量名的一致！
	Iterate = func(node *TreeNode) {
		if node == nil {
			return
		}
		Iterate(node.Left)
		nums = append(nums, node.Val)
		Iterate(node.Right)
	}

	Iterate(root)

	// 判断条件 i >= 0 不可少，处理 [1，1] 的情况
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] - nums[i+1] > 0 {
			return false
		}
	}

	return true
}