package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 注意审题：给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// 判断当前节点是否为叶子节点，如果是则直接比较判断即可
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
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
