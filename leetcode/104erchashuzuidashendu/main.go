package main

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max (l, r int) int {
	if l > r {
		return l
	}
	return r
}