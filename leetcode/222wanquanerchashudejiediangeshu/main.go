package main

// ref: https://leetcode-cn.com/problems/count-complete-tree-nodes/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func countNodesNotGraceful(root *TreeNode) int {
	resCount := 0

	var DFS func(*TreeNode)
	DFS = func(curNode *TreeNode) {
		if curNode == nil {
			return
		}
		resCount++
		if v := curNode.Left; v != nil {
			DFS(v)
		}
		if v := curNode.Right; v != nil {
			DFS(v)
		}
	}
	DFS(root)

	return resCount
}

// 进阶解法后面再来