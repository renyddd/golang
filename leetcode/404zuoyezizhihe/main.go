package main

/*
计算给定二叉树的所有左叶子之和。
https://leetcode-cn.com/problems/sum-of-left-leaves/
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var res int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			if root.Left != nil {
				if IsLeafNode(root.Left) {
					res += root.Left.Val
				}
				dfs(root.Left)
			}
			if root.Right != nil {
				dfs(root.Right)
			}
		}
	}

	dfs(root)

	return res
}

func IsLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func sumOfLeftLeaves_Official(root *TreeNode) int {
	var dfs func(root *TreeNode) (ans int)
	dfs = func(root *TreeNode) (ans int) {
		if root.Left != nil {
			if IsLeafNode(root.Left) {
				// 到达叶子节点，自然结束循环
				ans += root.Left.Val
			} else {
				ans += dfs(root.Left)
			}
		}
		if root.Right != nil {
			ans += dfs(root.Right)
		}

		return
	}

	return dfs(root)
}
