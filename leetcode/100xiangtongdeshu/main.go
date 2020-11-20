package main

import (
	"../treenode"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTreeDepthFirst
func isSameTreeDepthFirst(p *treenode.TreeNode, q *treenode.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		// 此种情况为真说明 p,q 有其一不为空
		return false
	} else if p.Val != q.Val {
		return false
	} else {
		return isSameTreeDepthFirst(p.Left, q.Left) && isSameTreeDepthFirst(p.Right, q.Right)
	}
}

// 题解中的判断使用了深度优先遍历，那可否用前中序遍历的结果来判断呢？
// 此处为什么用前序遍历写出来的逻辑很相似呢？

// isSameTreeBreadthFirst
func isSameTreeBreadthFirst(p *treenode.TreeNode, q *treenode.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	queue1, queue2 := []*treenode.TreeNode{p}, []*treenode.TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		node1, node2 := queue1[0], queue2[0]
		if node1.Val != node2.Val {
			return false
		}
		queue1, queue2 = queue1[1:], queue2[1:]
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if (left1 == nil && left2 != nil) || (left1 != nil && left2 != nil) {
			return false
		}
		if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
			return false
		}
		if left1 != nil {
			queue1 = append(queue1, left1)
		}
		if right1 != nil {
			queue1 = append(queue1, right1)
		}
		if left2 != nil {
			queue2 = append(queue2, left2)
		}
		if right2 != nil {
			queue2 = append(queue2, right2)
		}
	}
	return (len(queue1) == 0) && (len(queue2) == 0)
}
//isSameTreeBreadthFirstRebuild 用作错误记录
//func isSameTreeBreadthFirstRebuild() {
//		appendNotNil := func(slice []*treenode.TreeNode, elem *treenode.TreeNode) {
//			if elem == nil {
//				return
//			}
//			slice = append(slice, elem)
//		}
//		// 若以这种方式的话会存在当切片的底层数组需要扩容时地址会发生变化
//		appendNotNil(queue1, left1)
//		appendNotNil(queue1, right1)
//		appendNotNil(queue2, left2)
//		appendNotNil(queue2, right2)
//}