package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var TraverseIterator func(curNode *TreeNode, onePath string)
	TraverseIterator = func(curNode *TreeNode, onePath string) {
		if curNode == nil {
			return
		}
		onePath += fmt.Sprint(curNode.Val)
		if curNode.Left == nil && curNode.Right == nil {
			res = append(res, onePath)
			return
		} else {
			onePath += "->"
			TraverseIterator(curNode.Left, onePath)
			TraverseIterator(curNode.Right, onePath)
		}
	}

	TraverseIterator(root, "")

	return res
}
