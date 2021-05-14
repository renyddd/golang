package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	// tMap 用以保存序列化结果是否出现
	tMap := make(map[string]int)
	// result 保存 root 节点下的重复情况
	result := make([]*TreeNode, 0)

	var Traverse func(node *TreeNode) string
	Traverse = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		serialization := fmt.Sprintf("%d,%s,%s", node.Val, Traverse(node.Left), Traverse(node.Right))

		tMap[serialization]++
		if v, ok := tMap[serialization]; ok && v == 2 {
			result = append(result, node)
		}

		return serialization
	}

	Traverse(root)

	return result
}
