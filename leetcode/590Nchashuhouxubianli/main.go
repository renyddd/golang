package main

type Node struct {
	Val int
	Children []*Node
}
func postorder(root *Node) []int {
	type Elem struct {
		visited bool
		TNode *Node
	}

	stack := make([]Elem, 0)
	res := make([]int, 0)
	stack = append(stack, Elem{
		visited: false,
		TNode: root,
	})

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.TNode != nil {
			if !node.visited {
				stack = append(stack, Elem{true,  node.TNode})
				for i := len(node.TNode.Children); i >= 0; i-- {
					stack = append(stack, Elem{false, node.TNode.Children[i]})
				}
			} else {
				res = append(res, node.TNode.Val)
			}
		}
	}

	return res
}