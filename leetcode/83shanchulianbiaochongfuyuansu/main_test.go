package main

import "testing"

func TestDeleta(t *testing.T) {
	ll := []*ListNode{
		MakeListNode([]int{1,1,1,1}),
		MakeListNode([]int{1,2,3,3,3}),
		MakeListNode([]int{1,3,4,5,6}),
	}

	for _, l := range ll {
		deleteDuplicates(l).PrintList()
	}
}