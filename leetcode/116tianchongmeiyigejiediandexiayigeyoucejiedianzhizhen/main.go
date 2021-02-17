package main

// ref: https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val               int
	Left, Right, Next *Node
}

// connect 好像是并没有什么，只需要按照层序遍历将前一个节点的 Next 指针指向当前节点即可
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	toProcessNodesQueue := make([]*Node, 0)
	toProcessNodesQueue = append(toProcessNodesQueue, root)

	for len(toProcessNodesQueue) > 0 {
		nextNodesQueue := make([]*Node, 0)
		for i, v := range toProcessNodesQueue {
			// handle ptr next
			if i+1 < len(toProcessNodesQueue) {
				v.Next = toProcessNodesQueue[i+1]
			} else {
				v.Next = nil
			}

			// handle next traverse queue
			if l := v.Left; l != nil {
				nextNodesQueue = append(nextNodesQueue, l)
			}
			if r := v.Right; r != nil {
				nextNodesQueue = append(nextNodesQueue, r)
			}
		}
		toProcessNodesQueue = nextNodesQueue
	}

	return root
}

/* connectOfficial 官方解答，注意题目描述给定一颗**完美二叉树**
 因此每个 Node 的链接情况就可以分为两种，
  - 同一个父节点的，左二字对右儿子的链接
  - 跨父节点，前一个的右儿子对后一个的左儿子的链接
 并且遍历也可通过更新最左侧的节点来完成
 */
func connectOfficial(root *Node) *Node {
	if root == nil {
		return nil
	}

	// 先处理特殊的 root 节点，因为其不存在上一层的父节点
	root.Next = nil
	// leftMost 记为最左侧的节点，遍历其 Next 指针来
	leftMost := root
	// 跟新为其左儿子以保证永远从最左侧开始
	for ; leftMost.Left != nil; leftMost = leftMost.Left {
		// 开始以 leftMost 串起来的链表的层序遍历
		for p := leftMost; p != nil; p = p.Next {
			// 处理情况一
			p.Left.Next = p.Right

			// 处理情况二
			if p.Next != nil {
				p.Right.Next = p.Next.Left
			}
		}
	}

	return root
}

/*
递归的巧妙解答：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/solution/dong-hua-yan-shi-san-chong-shi-xian-116-tian-chong/
 */
func conn(root *Node) {
	if root == nil {
		return
	}

	nLeft, nRight := root.Left, root.Right

	for nLeft != nil {
		nLeft.Next = nRight
		nLeft = nLeft.Right
		nRight = nRight.Left
	}

	conn(root.Left)
	conn(root.Right)
}