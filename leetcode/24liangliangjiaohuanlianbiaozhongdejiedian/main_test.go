package main

import "testing"

func TestS(t *testing.T) {
	head := &ListNode{1, nil}
	n4 := &ListNode{4, nil}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	head.Next = n2

	t.Log(head)
	swapPairs(head)
}
