package main

import (
	"fmt"
)

func main() {
	head := MakeListNode([]int{1, 2, 3, 4, 7, 9})
	head.PrintList()

	head2 := MakeListNode([]int{7, 9, 12, 45})
	head2.PrintList()

	newHead := mergeTwoLists2(head, head2)
	newHead.PrintList()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type elem []int

func (e elem) Len() int {
	return len(e)
}

func (e elem) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e elem) Less(i, j int) bool {
	if e[i] < e[j] {
		return true
	} else {
		return false
	}
}

func (l *ListNode) PrintList() {
	p := l
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}

func (l *ListNode) AddToTail(n int) *ListNode {
	newL := &ListNode{
		Val:  n,
		Next: nil,
	}
	l.Next = newL
	return newL
}

func MakeListNode(nums []int) *ListNode {
	head := new(ListNode)
	p := head
	for i := 0; i < len(nums); i++ {
		p.Next = p.AddToTail(nums[i])
		p = p.Next
	}
	return head.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func mergeTwoLists2(l1, l2 *ListNode) *ListNode {
	preHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := preHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return preHead.Next
}
