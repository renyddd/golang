package main

// 迭代的方法 -> 用栈在显式的模拟递归
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	// node 为待处理节点
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			// 首先处理跟节点的值信息
			res = append(res, node.Val)
			// 再将该节点入栈 为处理其右子树做准备
			stack = append(stack, node)
			// 将待处理节点变为该节点的右子树的跟节点
			node = node.Left
		}
		// 类比第一次递归返回时的状态 -> 将右子树的跟节点成为待处理节点
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
		// 该待处理节点进入循环体
	}
	return res
}

func preorderTraversal_Iterate(root *TreeNode) []int {
	res := make([]int, 0)
	var iterate func(*TreeNode)
	iterate = func(root *TreeNode) {
		if root != nil {
			res = append(res, root.Val)
			iterate(root.Left)
			iterate(root.Right)
		}
	}

	return res
}

func preorderTraversal_visited(root *TreeNode) []int {
	type Elem struct {
		Visited bool
		Node    *TreeNode
	}
	var (
		res   = make([]int, 0)
		stack = make([]Elem, 0)
	)
	stack = append(stack, Elem{false, root})

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Node != nil {
			if !node.Visited {
				stack = append(stack, Elem{
					Visited: false,
					Node:    node.Node.Right,
				}, Elem{
					Visited: false,
					Node:    node.Node.Left,
				}, Elem{
					Visited: true,
					Node:    node.Node,
				})
			} else {
				res = append(res, node.Node.Val)
			}
		}
	}

	return res
}
