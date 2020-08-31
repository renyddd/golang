package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l3 := &ListNode{3, nil}
	l2 := &ListNode{5, l3}
	l1 := &ListNode{99, l2}
	head := &ListNode{0, l1}

	l3.Next = head

	if flag := hasCycle(head); flag != true {
		PrintList(head)
	} else {
		fmt.Println("Has Cycle!")
	}
}

func PrintList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Next)
	}
}

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	pFast, pSlow := head, head

	for pFast != nil && pFast.Next != nil {
		pSlow = pSlow.Next
		pFast = pFast.Next.Next

		if pFast == pSlow {
			return true
		}
	}

	return false
}
