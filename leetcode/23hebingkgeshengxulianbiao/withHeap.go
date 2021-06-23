package main

import (
	"container/heap"
)

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &intHeap{}

	for _, l := range lists {
		for p := l; p != nil; p = p.Next {
			// heap.Push(h, p.Val)
			// 之前插入元素时，无需始终保持堆序性质，使用直接向尾部添加的方式即可
			h.Push(p.Val)
		}
	}

	heap.Init(h)

	tail := &ListNode{Val: -1}
	dummyHead := tail

	for h.Len() > 0 {
		node := heap.Pop(h)
		if val, ok := node.(int); ok {
			tail.Next = &ListNode{Val: val}
			tail = tail.Next
		}
	}

	return dummyHead.Next
}
