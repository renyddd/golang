package main

import (
	"fmt"
)

func main() {
	head := MakeListNode([]int{1, 1, 2, 3, 3, 4, 7, 9, 9})
	head.PrintList()

	l := deleteDuplicates(head)
	l.PrintList()
}

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			// 搬移尾指针
			p.Next = p.Next.Next
		} else {
			// 如果不将 p 后移放在这里将会导致 eg：1，1，1 错过中间 1 的判断的情况
			p = p.Next
		}
	}
	return head
}

func deleteRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteRecursion(head.Next)
	if head.Val == head.Next.Val {
		head.Next = head.Next.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) PrintList() {
	p := l
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}

func (l *ListNode) AddToTail(n int) {
	newL := &ListNode{
		Val:  n,
		Next: nil,
	}
	l.Next = newL
}

func MakeListNode(nums []int) *ListNode {
	head := new(ListNode)
	p := head
	for i := 0; i < len(nums); i++ {
		p.AddToTail(nums[i])
		p = p.Next
	}
	return head.Next
}
