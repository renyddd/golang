package main

import (
	"testing"
	"unsafe"

	"github.com/renyddd/golang/leetcode/listnode"
)

func TestGetKthFromEnd(t *testing.T) {
	head := listnode.MakeListNode([]int{1, 2, 3, 4, 5, 6})
	head.PrintList()

	h := (*ListNode)(unsafe.Pointer(head))
	t.Log(h)

}
