package main

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max (l, r int) int {
	if l > r {
		return l
	}
	return r
}

// maxDepth2Recursion dfs 求解，明知不是最优解，只是为了锻炼一种思路
func maxDepth2Recursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	visitedDepth :=  make(map[int]struct{})

	var PreOrder func(*TreeNode, int)
	PreOrder = func(curNode *TreeNode, d int) {
		if _, ok := visitedDepth[d]; !ok {
			visitedDepth[d] = struct{}{}
		}
		if curNode.Left != nil {
			PreOrder(curNode.Left, d + 1)
		}
		if curNode.Right != nil {
			PreOrder(curNode.Right, d + 1)
		}
	}
	// root 已算作一层，并且叶子节点并不会导致加一
	PreOrder(root, 1)

	maxD := 0
	for key, _ := range visitedDepth {
		if key > maxD {
			maxD = key
		}
	}
	return maxD
}

// maxDepthBFS 层序遍历同样可以完成，明知不是最优解，只是为了锻炼一种思路
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// todo 如果此处 make 是 1，append 后会导致第一个元素为空吗？
	toProcessNodes := make([]*TreeNode, 0)
	toProcessNodes = append(toProcessNodes, root)
	// depth 同样也可以理解为遍历过的层次数
	depth := 0

	for len(toProcessNodes) > 0 {
		depth++
		nextLevelNodes := make([]*TreeNode, 0)
		// 这个先写对了，在优化为该层有数据后就进行下一层；
		// 每层必须要遍历完，又不是一个左沉树
		for _, v := range toProcessNodes {
			if p := v.Left; p != nil {
				nextLevelNodes = append(nextLevelNodes, p)
			}
			if p := v.Right; p != nil {
				nextLevelNodes = append(nextLevelNodes, p)
			}
		}
		toProcessNodes = nextLevelNodes
	}

	return depth
}
