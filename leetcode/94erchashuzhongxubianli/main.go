package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 这里的迭代遍历操作也就是在模仿一次函数的递归，通过恰当的数据结构来保存递归时的现场
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	// 首先处理跟节点
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			// 模拟左子树的递归操作
			stack = append(stack, node)
			// 利用循环迭代时的条件来判断左子树是否为空
			node = node.Left
		}
		// 类比中序遍历时-》遍历完左子树时的 访值时的情景
		node = stack[len(stack)-1]
		// 弹出
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		// 开始遍历右子树
		node = node.Right
	}
	return res
}

// ref: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/yan-se-biao-ji-fa-yi-chong-tong-yong-qie-jian-ming/
func inorderTraversal_2(root *TreeNode) []int {
	type Elem struct {
		Visited bool
		Node    *TreeNode
	}
	stack := make([]Elem, 0)
	stack = append(stack, Elem{false, root})
	res := make([]int, 0)

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if n.Node != nil {
			if n.Visited == false {
				stack = append(stack, Elem{false, n.Node.Right})
				stack = append(stack, Elem{true, n.Node})
				stack = append(stack, Elem{false, n.Node.Left})
			} else {
				res = append(res, n.Node.Val)
			}
		}
	}
	return res
}
