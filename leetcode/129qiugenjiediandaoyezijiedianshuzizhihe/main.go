package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// ref: https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

// 其实感觉与 113 题的框架很像，只是在中间的操作有所改变

func sumNumbers(root *TreeNode) int {
	var res int
	pathNums := make([]int, 0)

	var DFS func(*TreeNode)
	DFS = func(curNode *TreeNode) {
		if curNode == nil {
			return
		}

		pathNums = append(pathNums, curNode.Val)

		if curNode.Left == nil && curNode.Right == nil {
			res += doSum(pathNums)
		}

		defer func() {
			pathNums = pathNums[:len(pathNums)-1]
		}()

		DFS(curNode.Left)
		DFS(curNode.Right)
	}

	DFS(root)
	return res
}

func doSum(nums []int) int {
	tmp := 0
	for _, v := range nums {
		tmp *= 10
		tmp += v
	}

	return tmp
}

/*
   official solve
*/

func sumNumbersOfficial(root *TreeNode) int {
	return dfsOfficial(root, 0)
}

func dfsOfficial(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}

	sum := prevSum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return sum
	}

	return dfsOfficial(root.Left, sum) + dfsOfficial(root.Right, sum)
}

// bfs

type Pair struct {
	Node *TreeNode
	Num int
}

func sumNumberOfficialBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int
	nodesQueue := make([]Pair, 0)
	nodesQueue = append(nodesQueue, Pair{root, root.Val})

	for len(nodesQueue) > 0 {
		curNode := nodesQueue[0]
		nodesQueue = nodesQueue[1:]

		tmpSum := curNode.Num
		if curNode.Node.Left == nil && curNode.Node.Right == nil {
			res += tmpSum
		} else {
			if v := curNode.Node.Left; v != nil {
				nodesQueue = append(nodesQueue, Pair{v, tmpSum*10 + v.Val})
			}
			if v := curNode.Node.Right; v != nil {
				nodesQueue = append(nodesQueue, Pair{v, tmpSum*10 + v.Val})
			}
		}
	}

	return res
}