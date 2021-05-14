package main

/* 给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和
https://leetcode-cn.com/problems/deepest-leaves-sum/
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	maxDepth := -1
	res := 0

	var dfsIter func(root *TreeNode, d int)
	dfsIter = func(root *TreeNode, d int) {
		if root == nil {
			return
		}

		if d > maxDepth {
			maxDepth = d
			res = root.Val
		} else if d == maxDepth {
			res += root.Val
		}

		dfsIter(root.Left, d+1)
		dfsIter(root.Right, d+1)
	}

	dfsIter(root, 0)

	return res
}
