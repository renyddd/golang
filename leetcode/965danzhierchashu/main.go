package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 1
func isUnivalTreeIterate(root *TreeNode) bool {
	flag := root.Val
	stack := make([]*TreeNode, 1)
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			if node.Val != flag {
				return false
			}

			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}
	}

	return true
}

// 2
var res []int

func isUnivalTree_StoreRes(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res = make([]int, 0)

	internalRecurse(root)
	for i, _ := range res {
		if res[i] != root.Val {
			return false
		}
	}
	return true
}

func internalRecurse(root *TreeNode) {
	if root != nil {
		internalRecurse(root.Left)
		res = append(res, root.Val)
		internalRecurse(root.Right)
	}
}

// 3
func isUnivalTree(root *TreeNode) bool {
	leftCorrect := root.Left == nil ||
		(root.Val == root.Left.Val && isUnivalTree(root.Left))
	rightCorrect := root.Right == nil ||
		(root.Val == root.Right.Val && isUnivalTree(root.Right))

	return leftCorrect && rightCorrect
}
