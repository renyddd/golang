package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// bst 的查询，只需要在左右子树中做一次选择

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

func searchBST2(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil && node.Val != val {
		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node
}
