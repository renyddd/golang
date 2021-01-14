package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	list := preOrderTraversal(root)

	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left = nil
		prev.Right = curr
	}
}

func preOrderTraversal(root *TreeNode) []*TreeNode {
	list := make([]*TreeNode, 0)
	if root != nil {
		list = append(list, root)
		list = append(list, preOrderTraversal(root.Left)...)
		list = append(list, preOrderTraversal(root.Right)...)
	}
	return list
}

func flatten_Iterate(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	list := make([]*TreeNode, 0)
	node := root

	for root != nil || len(stack) > 0 {
		for node != nil {
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	for i := 1; i < len(stack); i++ {
		prev, curr := stack[i-1], stack[i]
		prev.Left = curr
		prev.Right = nil
	}
}

func flattenIterateVisited(root *TreeNode) {
	type Elem struct {
		Visited bool
		Node *TreeNode
	}
	var (
		stack = make([]Elem, 0)
		list = make([]*TreeNode, 0)
	)
	stack = append(stack, Elem{false, root})

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Node != nil {
			if !node.Visited {
				stack = append(stack,
					Elem{false, node.Node.Right},
					Elem{false, node.Node.Left},
					Elem{true, node.Node})
			} else {
				list = append(list, node.Node)
			}
		}
	}

	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}
