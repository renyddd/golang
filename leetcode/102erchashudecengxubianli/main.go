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

// levelOrderWithQueue 使用队列完成层序遍历
func levelOrderWithQueue(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	// nextLevelNodes 其实就代表着下一次待处理的队列
	nextLevelNodes := make([]*TreeNode, 0)
	nextLevelNodes = append(nextLevelNodes, root)

	for len(nextLevelNodes) > 0 {
		tmpLevelNodes := make([]*TreeNode, 0)
		aLevelRes := make([]int, 0)

		for _, v := range nextLevelNodes {
			aLevelRes = append(aLevelRes, v.Val)
			if v.Left != nil {
				tmpLevelNodes = append(tmpLevelNodes, v.Left)
			}
			if v.Right != nil {
				tmpLevelNodes = append(tmpLevelNodes, v.Right)
			}
		}

		// TODO 需要思考，go gc 会怎么处理对这两个切片的回收？
		// TODO 为什么使用 copy 就不行呢？
		nextLevelNodes = tmpLevelNodes

		res = append(res, aLevelRes)
	}

	return res
}

// levelOrderWithQueueByChan channel implement queue
func levelOrderWithQueueByChan(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	toProcessLevelNodes := make(chan *TreeNode,800)
	toProcessLevelNodes <- root

	for len(toProcessLevelNodes) > 0 {
		tmpLevelNodes := make(chan *TreeNode, 100)
		aLevelRes := make([]int, 0)

		var v *TreeNode
		for len(toProcessLevelNodes) > 0 {
			v = <- toProcessLevelNodes
			aLevelRes = append(aLevelRes, v.Val)
			if v.Left != nil {
				tmpLevelNodes <- v.Left
			}
			if v.Right != nil {
				tmpLevelNodes <- v.Right
			}
		}

		for len(tmpLevelNodes) > 0 {
			toProcessLevelNodes <- <-tmpLevelNodes
		}

		res = append(res, aLevelRes)
	}

	return res
}