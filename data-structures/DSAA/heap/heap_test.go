package heap

import (
	"fmt"
	"testing"
)

/*
	一个来自官方实现的 intHeap 的例子
	https://golang.org/pkg/container/heap/
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push 与 Pop 使用指针是因为切片长度的可能会导致更新底层数组
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5, 7, 4, 9, 8, 3, 6}
	Init(h)
	Push(h, 10)
	Push(h, 0)
	fmt.Println("minimum:", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", Pop(h))
	}
}

func BenchmarkInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := &IntHeap{2, 1, 5, 7, 4, 9, 8, 3, 6}
		Init(h)
	}
}
