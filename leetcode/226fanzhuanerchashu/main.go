package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	doInvertTree_recurse(root)
	return root
}

func doInvertTree_recurse(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	doInvertTree_recurse(root.Left)
	doInvertTree_recurse(root.Right)
}

func doInvertTree_iteration(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	type Elem struct {
		Visited bool
		Node    *TreeNode
	}
	stack := make([]Elem, 1)
	stack = append(stack, Elem{
		Visited: false,
		Node:    root,
	})

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Node != nil {
			if !node.Visited {
				stack = append(stack, Elem{true, node.Node})
				stack = append(stack, Elem{false, node.Node.Right})
				stack = append(stack, Elem{false, node.Node.Left})
			} else {
				node.Node.Left, node.Node.Right = node.Node.Right, node.Node.Left
			}
		}
	}

	return root
}
