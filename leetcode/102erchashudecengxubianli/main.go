package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	type Elem struct {
		Node    *TreeNode
		Visited bool
		Level   int
	}

	resWithLevel := make([]Elem, 0)
	stack := make([]Elem, 0)
	stack = append(stack, Elem{root, false, 0})

	for len(stack) > 0 {
		cur_node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur_node.Node != nil {
			if !cur_node.Visited {
				resWithLevel = append(resWithLevel, Elem{cur_node.Node, true, cur_node.Level})
				stack = append(stack, Elem{cur_node.Node, true, cur_node.Level})
				stack = append(stack, Elem{cur_node.Node.Right, false, cur_node.Level + 1})
				stack = append(stack, Elem{cur_node.Node.Left, false, cur_node.Level + 1})
			}
		}
	}

	ms := make(map[int][]int, 0)
	retSlice := make([][]int, 0)

	for _, v := range resWithLevel {
		l := v.Level
		ms[l] = append(ms[l], v.Node.Val)
	}

	// 因为 map 会无序输出，并且其键值最大值即为其长度减一（level）
	for i := 0; i < len(ms); i++ {
		if value, ok := ms[i]; ok {
			retSlice = append(retSlice, value)
		}
	}

	return retSlice
}
