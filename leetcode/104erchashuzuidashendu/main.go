package main

import (
	"../treenode"
)

func maxDepth(root *treenode.TreeNode) int {
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