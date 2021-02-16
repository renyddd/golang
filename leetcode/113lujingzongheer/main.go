package main

// question ref: https://leetcode-cn.com/problems/path-sum-ii/

// https://blog.golang.org/slices
// 2D slice
// - https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go
// - https://stackoverflow.com/questions/39561140/what-is-two-dimensional-arrays-memory-representation

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// pathSum 总的思路还是走 dfs 递归遍历存储每次判断相等的路径切片
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	// 存储整个结果
	res := make([][]int, 0)

	var DFS func(curNode *TreeNode, curSum int, pathNums []int)
	// DFS 尽在某叶子节点判断复合后进行结果的保存
	DFS = func(curNode *TreeNode, curSum int, pathNums []int) {
		if curNode == nil {
			return
		}
		// 对！如果此处不生成新的切片（该且切片指针指向新的一块内存）则变量过程中结果将紊乱
		newPath := make([]int, 0)
		newPath = append(newPath, pathNums...)
		newPath = append(newPath, curNode.Val)

		// 首先是基准条件判断
		if curNode.Left == nil && curNode.Right == nil {
			// leaf
			if curSum == curNode.Val {
				res = append(res, newPath)
			}
		}
		// 假设情况已经正确返回后的操作
		// 递归调用左右子树
		DFS(curNode.Left, curSum-curNode.Val, newPath)
		DFS(curNode.Right, curSum-curNode.Val, newPath)
	}

	// 在递归的过程成，可能是因为 append 导致的切片翻倍，才致使后续的操作没受指针的影响
	DFS(root, targetSum, make([]int, 0))
	return res
}

// 官方解，然后探究 slice 的特性
// 只是注意要关注特定语言的特性哈

// pathSumOfficial 为什么官方解这里使用切片的方式就更剩内存呢？
func pathSumOfficial(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	// curPathNum 记录当前遍历路径时的结果
	curPathNum := make([]int, 0)

	// DFS 同样是深度优先遍历，不同是在切片的应用
	var DFS func(*TreeNode, int)
	DFS = func(curNode *TreeNode, curNumLeft int) {
		if curNode == nil {
			return
		}
		// 路径判断时同样传递的是剩余值
		curNumLeft -= curNode.Val
		curPathNum = append(curPathNum, curNode.Val)

		defer func() {
			curPathNum = curPathNum[:len(curPathNum)-1]
		}()

		if curNode.Left == nil && curNode.Right == nil && curNumLeft == 0 {
			res = append(res, append(make([]int, 0), curPathNum...))
			return
		}
		DFS(curNode.Left, curNumLeft)
		DFS(curNode.Right, curNumLeft)
	}

	DFS(root, targetSum)

	return res
}
