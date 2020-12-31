// 堆：
// - 堆是一颗完全二叉树（除了最后一层，其他层的节点个数都是满的，最后一层的节点都靠左排列
// - 堆中的每一个节点的值都必须大于或等于其子树中每个节点的值
package main

import "fmt"

type Heap struct {
	data     []int
	capacity int // 堆可存储的最大数据个数
	count    int // 堆中已存在的数据个数
}

func NewHeap(capacity int) *Heap {
	// 注意堆的节点运算方式 left = 2*i, right = 2*i+1
	// 越过 index 为 0 的下标以避免数组越界
	data := make([]int, capacity+1)
	return &Heap{
		data:     data,
		capacity: capacity,
		count:    0,
	}
}

func (h *Heap) String() string {
	res := make([]int, 0)
	for i := 0; i < h.count+1; i++ {
		res = append(res, h.data[i])
	}
	return fmt.Sprint(res, h.capacity, h.count)
}

func (h *Heap) Insert(item int) {
	if h.count >= h.capacity {
		return
	}
	h.count++
	h.data[h.count] = item

	// heapiyf 从下往上的堆化；让新插入尾部的节点与其父节点进行比较使其满足子节点小于父节点的关系
	for i := h.count; i/2 > 0 && h.data[i] > h.data[i/2]; {
		// i / 2 > 0 表示不能为数组的 0 索引赋数据值
		h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
		i /= 2
	}
}

func (h *Heap) DeleteTopElem() {
	if h.count == 0 {
		return
	}
	h.data[1] = h.data[h.count]
	h.count--
	heapify(h.data, h.count, 1)
}

func heapify(a []int, n, i int) {
	for {
		maxPos := i
		if i*2 <= n && a[i] < a[i*2] {
			maxPos = i * 2
		}
		if i*2+1 <= n && a[maxPos] < a[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		a[i], a[maxPos] = a[maxPos], a[i]
		i = maxPos
	}
}

func NewHeapWithSlice1(a []int) *Heap {
	capacity := len(a)
	heap := NewHeap(capacity)
	for i := 0; i < capacity; i++ {
		heap.Insert(a[i])
	}
	return heap
}

func NewHeapWithSlice2(a []int) *Heap {
	n := len(a)
	heap := NewHeap(n)
	copy(heap.data[1:], a)
	//fmt.Printf("------ %d times heapify: %v\n", 0, heap.data)
	heap.count = n
	for i := n / 2; i >= 1; i-- {
		heapify(heap.data, n, i)
		//fmt.Printf("------ %d times heapify: %v\n", i, heap.data)
	}
	return heap
}

func (h *Heap) Sort() {
	k := h.capacity
	for k > 1 {
		h.data[1], h.data[k] = h.data[k], h.data[1]
		k--
		heapify(h.data, k, 1)
	}
}
