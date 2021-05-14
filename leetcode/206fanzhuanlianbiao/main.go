package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/reverse-linked-list
func reverseListIterate(head *ListNode) *ListNode {
	var prev *ListNode

	p := head
	for p != nil {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}

	return prev
}

func reverseList(head *ListNode) *ListNode {
	// 首先处理 head 为空的特殊情况
	// 以及处理值愿链表尾部时的递归终止情况
	if head == nil || head.Next == nil {
		return head
	}

	// @p: 注意其始终指向原链表的尾节点
	// 从当自外层返回时也可看出；其一直保留并传递着递归终止条件时的值
	p := reverseList(head.Next)

	// 接下来的两步是反转
	head.Next.Next = head
	head.Next = nil

	return p
}

/*
	上面所说的"反转"其实我觉的不是特别的恰当，	这更像是一种"插入"的过程。
	遇到递归终止条件后，p 始终指向原链表的末尾元素，也就是新链表的头节点。
	自从 origin 状态器，该链表就已被隐形的分为了两条叉，分别是 head, p

	每一个"插入"操作，都可以理解为是从"正序部分"链表中取出其尾元素，
	插入至"逆序部分"的尾部。


	origin:
	1 -> 2 -> 3 -> 4 -> 5 -> nil
					^	^
				  head	p

	first:
	1 -> 2 -> 3         5 -> 4 -> nil
              ^--------------^
              ^          ^
            head         p

	second:
	1 -> 2              5 -> 4 -> 3 -> nil
         ^------------------------^
         ^               ^
       head              p

	third:
	1      				5 -> 4 -> 3 -> 2 -> nil
    ^----------------------------------^
    ^                     ^
   head                   p

	forth:
	5 -> 4 -> 3 -> 2 -> 1 -> nil
	^
	p, head
*/
