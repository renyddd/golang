package main

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	node := root

	for node != nil {
		// 此处相等的条件判断没必要的！否则就会创建出一条单练支出来
		if val == node.Val {
			break
		} else if val < node.Val {
			if node.Left != nil {
				node = node.Left
			} else {
				node.Left = &TreeNode{val, nil, nil}
			}
		} else {
			if node.Right != nil {
				node = node.Right
			} else {
				node.Right = &TreeNode{val, nil, nil}
			}
		}
	}
	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if val == root.Val {
		return root
	} else if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{val, nil, nil}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{val, nil, nil}
		} else {
			insertIntoBST(root.Right, val)
		}
	}
	return root
}