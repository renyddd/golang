package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// ref: https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
// 层序遍历练手

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	curNodes := []*TreeNode{root}
	res := make([]float64, 0)

	for len(curNodes) > 0 {
		toProcessNodes := make([]*TreeNode, 0)
		var sum float64 = 0
		for _, v := range curNodes {
			sum += float64(v.Val)
			if v.Left != nil {
				toProcessNodes = append(toProcessNodes, v.Left)
			}
			if v.Right != nil {
				toProcessNodes = append(toProcessNodes, v.Right)
			}
		}
		res = append(res, sum/float64(len(curNodes)))
		curNodes = toProcessNodes
	}

	return res
}
