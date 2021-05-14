package main

type Node struct {
	Val      int
	Children []*Node
}

var res []int

func preorder(root *Node) []int {
	res = make([]int, 0)
	addElem(root)
	return res
}

func addElem(root *Node) {
	if root != nil {
		res = append(res, root.Val)
		for _, p := range root.Children {
			addElem(p)
		}
	}
}

func preorder_elem(root *Node) []int {
	type Elem struct {
		Visited bool
		ENode   *Node
	}

	stack := make([]Elem, 0)
	stack = append(stack, Elem{false, root})
	res := make([]int, 0)

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if n.ENode != nil {
			if n.Visited == false {
				for i := len(n.ENode.Children) - 1; i >= 0; i-- {
					stack = append(stack, Elem{false, n.ENode.Children[i]})
				}
				stack = append(stack, Elem{true, n.ENode})
			} else {
				res = append(res, n.ENode.Val)
			}
		}
	}
	return res
}
