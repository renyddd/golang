package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 注意审题：给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和
/*
ref: https://leetcode-cn.com/problems/path-sum/
 */

// hasPathSum 递归
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return true
		} else { return false }
	} else {
		return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
	}
}

// hasPathSumBFS 用广度优先遍历，当进行至叶子结点的时候就判断整条路径和是否满足目标值
func hasPathSumBFS(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	toProcessNodes := make([]*TreeNode, 0)
	// 要从最开始的给队列的命名处就开始想清楚，你要的是队列还是栈
	valuesQueue := make([]int, 0)
	toProcessNodes = append(toProcessNodes, root)
	// 一个实时跟进的保存计算结果的队列
	valuesQueue = append(valuesQueue, root.Val)

	for len(toProcessNodes) > 0 {
		// 队列的操作同样也可以用 buffered channel 实现
		curNode := toProcessNodes[0]
		curValue := valuesQueue[0]
		toProcessNodes = toProcessNodes[1:]
		valuesQueue = valuesQueue[1:]

		if curNode.Left == nil && curNode.Right == nil {
			if curValue == sum {
				return true
			}
		}

		if curNode.Left != nil {
			toProcessNodes = append(toProcessNodes, curNode.Left)
			valuesQueue = append(valuesQueue, curValue + curNode.Left.Val)
		}

		if curNode.Right != nil {
			toProcessNodes = append(toProcessNodes, curNode.Right)
			valuesQueue = append(valuesQueue, curValue + curNode.Right.Val)
		}
	}

	return false
}

// hasPathSumDFS 迭代实现的深度优先遍历还是以前序遍历为主
func hasPathSumDFS(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	toProcessNodesStack := make([]*TreeNode, 0)
	realTimeValueStack := make([]int, 0)
	// diff
	curNode, curValue := root, 0

	// diff 【要注意的细节太多了】
	/*
	 - 避免空指针的解引用
	 - value stack 的赋值为增量的赋值
	 - for 循环的条件判断需要带上当前节点不为空的判断，因为待处理栈很容易为空
	 */
	for curNode != nil || len(toProcessNodesStack) > 0 {
		for curNode != nil {
			toProcessNodesStack = append(toProcessNodesStack, curNode)
			// 注意传递的是增量值
			curValue += curNode.Val
			realTimeValueStack = append(realTimeValueStack, curValue)
			curNode = curNode.Left
		}

		curNode = toProcessNodesStack[len(toProcessNodesStack)-1]
		toProcessNodesStack = toProcessNodesStack[:len(toProcessNodesStack)-1]
		curValue = realTimeValueStack[len(realTimeValueStack)-1]
		realTimeValueStack = realTimeValueStack[:len(realTimeValueStack)-1]

		if curNode.Left == nil && curNode.Right == nil {
			if curValue == sum {
				return true
			}
		}

		if v := curNode.Right; v != nil {
			toProcessNodesStack = append(toProcessNodesStack, v)
			// curValue 这里同样传递的为增量值
			curValue += v.Val
			realTimeValueStack = append(realTimeValueStack, curValue)
		}
		curNode = curNode.Right
	}

	return false
	/* 以下用以分析 DFS 与 BFS 首元素赋值方式的差异

	toProcessStack := make([]*TreeNode, 0)
	toProcessStack = append(toProcessStack, root)
	// 一个保存实时遍历到的节点的计算和的栈
	realTimeValueStack := make([]int, 0)
	realTimeValueStack = append(realTimeValueStack, root.Val)

	var (
		curNode *TreeNode
		curValue int
	)
	for len(toProcessStack) > 0 {
		l := len(toProcessStack)
		curNode = toProcessStack[l-1]
		toProcessStack = toProcessStack[:l-1]
		curValue = realTimeValueStack[l-1]
		realTimeValueStack = realTimeValueStack[:l-1]

		// todo

		// 模拟持续的递归左子树
		for curNode.Left != nil {
			curNode = curNode.Left
			if curNode != nil {
				toProcessStack = append(toProcessStack, curNode)
				realTimeValueStack = append(realTimeValueStack, curValue + curNode.Val)
			}
		}

		if curNode.Right != nil {
			toProcessStack = append(toProcessStack, root.Right)
		}
	}
	 */
}



/*
 -------------------------------------------------
 */

//func hasPathSum_old(root *TreeNode, sum int) bool {
//	if root == nil {
//		return false
//	}
//	// 判断当前节点是否为叶子节点，如果是则直接比较判断即可
//	if root.Left == nil && root.Right == nil {
//		return sum == root.Val
//	}
//	// 上面当找到不满足的叶子节点后不返回 false 的会会徒多一层对叶子节点的递归调用
//	return hasPathSum_old(root.Left, sum-root.Val) || hasPathSum_old(root.Right, sum-root.Val)
//}

// hasPathSum_TwoQueues 使用两个队列分别存储处理节点与运算和结果
// 是否与迭代的层序遍历有点像？对，这就是层序遍历
func hasPathSum_TwoQueues(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	nodeQueue := make([]*TreeNode, 0)
	valQueue := make([]int, 0)
	nodeQueue = append(nodeQueue, root)
	valQueue = append(valQueue, root.Val)

	for len(nodeQueue) > 0 {
		cur_node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		cur_val := valQueue[0]
		valQueue = valQueue[1:]

		// 判断该叶子节点累加和是否满足
		if cur_node.Left == nil && cur_node.Right == nil {
			if cur_val == sum {
				return true
			}
		}

		// 继续处理非叶子节点
		if cur_node.Left != nil {
			nodeQueue = append(nodeQueue, cur_node.Left)
			valQueue = append(valQueue, cur_val + cur_node.Left.Val)
		}
		if cur_node.Right != nil {
			nodeQueue = append(nodeQueue, cur_node.Right)
			valQueue = append(valQueue, cur_val + cur_node.Right.Val)
		}
	}
	return false
}

func hasPathSum_iterate(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	type Elem struct {
		Visited bool
		Node    *TreeNode
		LeafSum int
	}
	stack := make([]Elem, 0)
	stack = append(stack, Elem{false, root, root.Val})

	for len(stack) > 0 {
		cur_node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur_node.Node != nil {
			if !cur_node.Visited {
				if cur_node.Node.Right != nil {
					stack = append(stack, Elem{false, cur_node.Node.Right, cur_node.LeafSum + cur_node.Node.Right.Val})
				} else {
					// 处理空指针解引用的情况
					stack = append(stack, Elem{false, cur_node.Node.Right, cur_node.LeafSum})
				}

				stack = append(stack, Elem{true, cur_node.Node, cur_node.LeafSum})

				if cur_node.Node.Left != nil {
					stack = append(stack, Elem{false, cur_node.Node.Left, cur_node.LeafSum + cur_node.Node.Left.Val})
				} else {
					stack = append(stack, Elem{false, cur_node.Node.Left, cur_node.LeafSum})
				}
			} else {
				if cur_node.Node.Left == nil && cur_node.Node.Right == nil {
					if cur_node.LeafSum == sum {
						return true
					}
				}
			}
		}
	}
	return false
}
