package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/*
	注意题解中的求和方式
*/
//ref: https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var DFS func(root *TreeNode)
	DFS = func(root *TreeNode) {
		if root == nil {
			return
		}
		DFS(root.Right)
		sum += root.Val
		root.Val = sum
		DFS(root.Left)
	}

	DFS(root)

	return root
}
