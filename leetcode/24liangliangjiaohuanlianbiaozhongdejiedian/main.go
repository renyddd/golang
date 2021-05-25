package main

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs 为节点提前标记好序号是个好办法，只需注意指向的哪个节点需要递归调用
// 哪个需要返回
func swapPairs(head *ListNode) *ListNode {
	var iter func(*ListNode) *ListNode
	iter = func(p *ListNode) *ListNode {
		if p == nil || p.Next == nil {
			return p
		}

		n1, n2, n3 := p, p.Next, p.Next.Next
		n2.Next = n1
		n1.Next = swapPairs(n3)

		return n2
	}

	return iter(head)
}

func swapPairs2(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	p := dummyHead

	// todo
}
