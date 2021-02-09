package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
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
