package main

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// minDepthDFs 深度优先统一存储结果后比较返回
func minDepthDFs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	visitedDepth := make(map[int]struct{})

	var DFS func(*TreeNode, int)
	DFS = func(curNode *TreeNode, d int) {
		if curNode == nil {
			return
		}
		if curNode.Left == nil && curNode.Right == nil {
			visitedDepth[d] = struct{}{}
		}
		DFS(curNode.Left, d + 1)
		DFS(curNode.Right, d + 1)
	}

	DFS(root, 1)
	minD := math.MaxInt64
	for key, _ := range visitedDepth {
		if key < minD {
			minD = key
		}
	}

	return minD
}

// minDepthBFS 层序优先遍历遇见叶子结点就返回结果
func minDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	toProcessNodes := make([]*TreeNode, 0)
	toProcessNodes = append(toProcessNodes, root)
	depth := 0

	for len(toProcessNodes) > 0 {
		depth++
		nextLevelNodes := make([]*TreeNode, 0)
		for _, v := range toProcessNodes {
			if v.Left == nil && v.Right == nil {
				// 遇见了叶子结点即刻返回
				return depth
			}
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

/*
ref: https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
先看上面的
 */

// input: root = [3,9,20,nil,nil,15,7]
// 首先了解树的特性及到达叶子结点时的特性
// 当 i 的 2i 与 2i+1 均为 nil 时即为叶子节点

// 深度优先
// 这种递归应为更优方式，因为没有改变传入值没有带来副作用
func minDepth_old(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	// 求最小路径时记得初始化为最大值
	var minD = math.MaxInt32
	if root.Left != nil {
		minD = mathMin(minDepth_old(root.Left), minD)
	}
	if root.Right != nil {
		minD = mathMin(minDepth_old(root.Right), minD)
	}

	return minD + 1

	//
	//var minD int
	//if root.Left != nil {
	//	minD = minDepth(root.Left)
	//}
	//if root.Right != nil {
	//	minD = mathMin(minDepth(root.Right), minD)
	//}
	//return minD + 1
	// 注意上述方法的错误：当如果 root.Left 为空跳过了 minD 的赋值后，便会导致 Right 的比较
	// 是在跟初始值 0 在比较
}

func mathMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 广度优先
// 对这棵树进行广度优先遍历时，直接返回先找见的这个叶子结点的深度即为整棵树的最小深度
func minDepthBreadth_old(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	count := make([]int, 0)
	queue = append(queue, root)
	count = append(count, 1)

	/*
	相比较而言，这是队列的另外一种使用方式
	只不过当用 toProcessNodes, nextLevelNodes 时，一次循环就是一层
	 */
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth + 1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth + 1)
		}
	}
	return 0
}
