package heap

import (
	"fmt"
	"testing"
)

// 堆排序，利用堆的删除最小元素的性质
// 对待排序数组初始化堆，每次删除最小元素即可

type HInt []int

func (h HInt) Len() int           { return len(h) }
func (h HInt) Less(i, j int) bool { return h[i] < h[j] }
func (h HInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TestHeapSort(t *testing.T) {
	//heap := &HInt{4, 2, 7, 5, 3, 1, 9, 8, 6}
	heap := &HInt{2, 1, 5, 7, 4, 9, 8, 3, 6}
	Init(heap)

	sorted := make([]int, 0)
	fmt.Printf("minimum: %d\n", (*heap)[0])
	for heap.Len() > 0 {
		v := Pop(heap).(int)
		fmt.Println(v)
		sorted = append(sorted, v)
	}

	fmt.Println(sorted)
}
