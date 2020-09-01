package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	rear := &ListNode{9, nil}
	c := &ListNode{666, rear}
	l12 := &ListNode{4, c}
	h1 := &ListNode{1, l12}
	h2 := &ListNode{1, c}

	PrintList(h1)
	PrintList(h2)
	fmt.Println("========")

	fmt.Println(FindFirstCommonNode(h1, h2))
}

func PrintList(head *ListNode) {
	if head == nil {
		return
	}
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}

func FindFirstCommonNode(pHead1, pHead2 *ListNode) *ListNode {
	len1 := ListLen(pHead1)
	len2 := ListLen(pHead2)
	var k int

	if k = len1 - len2; k < 0 {
		k = k * -1
		for ; k > 0; k-- {
			pHead2 = pHead2.Next
		}
	} else {
		for ; k > 0; k-- {
			pHead1 = pHead1.Next
		}
	}

	for pHead1 != nil || pHead2 != nil {
		pHead1 = pHead1.Next
		pHead2 = pHead2.Next

		if pHead1 == pHead2 {
			return pHead1
		}
	}

	return nil
}

func ListLen(head *ListNode) int {
	var num int

	for ; head != nil; head = head.Next {
		num += 1
	}

	return num
}
