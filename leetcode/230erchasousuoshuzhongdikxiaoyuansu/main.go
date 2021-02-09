package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthSmallestIterate(root *TreeNode, k int) int {
	nums := make([]int, 0)
	var inorderTraversal func(root *TreeNode)
	inorderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraversal(root.Left)
		nums = append(nums, root.Val)
		inorderTraversal(root.Right)
	}

	inorderTraversal(root)

	return nums[k-1]
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil || k == 0 {
		return 0
	}
	times := 0
	stack := make([]*TreeNode, 0)
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		times++
		if k == times {
			return node.Val
		}
		node = node.Right
	}
	return 0
}

func kthSmallest2(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	nums := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nums = append(nums, node.Val)
		node = node.Right
	}

	return nums[k-1]
}