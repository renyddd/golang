package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 平衡二叉树（Balanced BinaryTree）又被称为AVL树。它具有以下性质：
// 它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 判断左右两子树高度差是否小于一
	return abs(height(root.Left)-height(root.Right)) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
