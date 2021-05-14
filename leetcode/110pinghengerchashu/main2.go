package main

func isBalanced2(root *TreeNode) bool {
	return height2(root) >= 0
}

func height2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height2(root.Left)
	rightHeitht := height2(root.Right)
	if leftHeight == -1 || rightHeitht == -1 || abs(leftHeight-rightHeitht) > 1 {
		return -1
	}
	return max(leftHeight, rightHeitht) + 1
}
