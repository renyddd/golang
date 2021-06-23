package main

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap(t *testing.T) {
	h := &intHeap{}
	heap.Init(h)

	heap.Push(h, 3)
	heap.Push(h, 7)
	heap.Push(h, 8)
	heap.Push(h, 6)
	heap.Push(h, 4)

	fmt.Printf("minimum: %d\n", (*h)[0])
	fmt.Printf("struct: %v\n", h)

	fmt.Printf("start: %d, end: %d\n", (*h)[0], (*h)[h.Len()-1])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

}
