package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
// ref: https://leetcode-cn.com/problems/binary-tree-tilt/

//findTilt 完全可以递归进行处理
func findTilt(root *TreeNode) int {
	sum := 0

	var doFind func(root *TreeNode) int
	// 返回值为当前节点的 tilt 值
	doFind = func(root *TreeNode) int {
		// 处理空节点
		if root == nil {
			return 0
		}
		le := doFind(root.Left)
		ri := doFind(root.Right)

		if le - ri < 0 {
			sum += ri - le
		} else {
			sum += le - ri
		}

		// 还没搞明白此处递归的逻辑
		return le + ri + root.Val
	}

	doFind(root)

	return sum
}