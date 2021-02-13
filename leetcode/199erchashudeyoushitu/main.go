package main

/*
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

// 这不就是层序遍历后输出每个切片的最后的元素

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

// rightSideViewNotGraceful 不优雅的处理
func rightSideViewNotGraceful(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	rr := make([]int, 0)
	res := make([][]int, 0)
	toProcessLevel := make([]*TreeNode, 0)
	toProcessLevel = append(toProcessLevel, root)

	for len(toProcessLevel) > 0 {
		tmpLevelNodes := make([]*TreeNode, 0)
		aLevelRes := make([]int, 0)

		for _, v := range toProcessLevel {
			aLevelRes = append(aLevelRes, v.Val)
			if v.Left != nil {
				tmpLevelNodes = append(tmpLevelNodes, v.Left)
			}
			if v.Right != nil {
				tmpLevelNodes = append(tmpLevelNodes, v.Right)
			}
		}
		toProcessLevel = tmpLevelNodes
		res = append(res, aLevelRes)
	}


	for i, _ := range res {
		rr = append(rr, res[i][len(res[i])-1])
	}

	return rr
}

// rightSideViewBFS
func rightSideViewBFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	toProcessNodes := make([]*TreeNode, 0)
	toProcessNodes = append(toProcessNodes, root)

	for len(toProcessNodes) > 0 {
		tmpLevelNodes := make([]*TreeNode, 0)
		tmpRes := make([]int, 0)
		for _, v := range toProcessNodes {
			tmpRes = append(tmpRes, v.Val)
			if v.Left != nil {
				tmpLevelNodes = append(tmpLevelNodes, v.Left)
			}
			if v.Right != nil {
				tmpLevelNodes = append(tmpLevelNodes, v.Right)
			}
		}

		res = append(res, tmpRes[len(tmpRes)-1])
		toProcessNodes = tmpLevelNodes
	}
	return res
}

// DFS 进行深度优先遍历（根右左）时，每层访问到的第一个节点就是需求节点； 如何处理层的深度？
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	// 层数信息预先无法得知
	visitedDepth := make(map[int]struct{}, 0)

	var DFS func(node *TreeNode, d int)
	DFS = func(node *TreeNode, d int) {
		if node == nil {
			return
		}
		// 如果该层是第一次访问那一定是右视图的元素
		// 结果元素处理
		if _, ok := visitedDepth[d]; !ok {
			visitedDepth[d] = struct{}{}
			res = append(res, node.Val)
		}

		// 根右左，非标准的后序遍历
		DFS(node.Right, d + 1)
		DFS(node.Left, d + 1)
	}

	DFS(root, 0)
	return res
}