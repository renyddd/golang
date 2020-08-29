package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := &ListNode{9, nil}
	head1 := &ListNode{6, n1}

	n2 := &ListNode{15, nil}
	head2 := &ListNode{3, n2}

	PrintList(head1)
	PrintList(head2)

	newHead := Merged(head1, head2)
	fmt.Println("newhead: ", newHead)
	PrintList(newHead)
}

func PrintList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println("list in: ", head)
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}

func Merged(pHead1 *ListNode, pHead2 *ListNode) *ListNode {

	newHead := &ListNode{}
	cur := newHead

	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		}
		cur = cur.Next
	}

	// 将剩余链表的尾部全部补全
	if pHead1 != nil {
		cur.Next = pHead1
	} else {
		cur.Next = pHead2
	}

	return newHead.Next
}
