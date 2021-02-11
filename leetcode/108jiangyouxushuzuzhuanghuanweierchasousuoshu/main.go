package main

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	// 此处已将 nums 切片的原长度抽象为了递归函数的左右边界
	return construct(nums, 0, len(nums) - 1)
}

func construct(nums []int, leftBoundary, rightBoundary int) *TreeNode {
	// 既是递归的终止条件，又用以构造叶子节点
	if leftBoundary > rightBoundary {
		return nil
	}
	rootIndex := (leftBoundary + rightBoundary) / 2
	root := new(TreeNode)
	root.Val = nums[rootIndex]
	root.Left = construct(nums, leftBoundary, rootIndex - 1)
	root.Right = construct(nums, rootIndex + 1, rightBoundary)

	return root
}

// TODO: 该题的扩展点还是很多的，可以探讨根据不同的选取根节点的方式而创造出不通形态的二叉树的依据是什么
