package main

import "testing"

func TestHeap_Insert(t *testing.T) {
	h := NewHeap(5)
	t.Log(h)

	h.Insert(5)
	h.Insert(9)
	h.Insert(3)
	h.Insert(2)
	t.Log(h)
	h.DeleteTopElem()
	t.Log(h)
}

func TestNewHeapWithSlice1(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{11, 22, 33, 4, 5, 67, 7, 8},
	}

	for _, v := range tests {
		h := NewHeapWithSlice1(v)
		t.Log(h)
	}
}

func TestNewHeapWithSlice2(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{11, 22, 33, 4, 5, 67, 7, 8},
	}

	for _, v := range tests {
		h := NewHeapWithSlice2(v)
		t.Log(h)
	}
}

func TestSort(t *testing.T) {
	tests := [][]int{
		{9, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{11, 22, 3, 4, 5, 67, 7, 8, 1},
	}

	for _, v := range tests {
		h := NewHeapWithSlice2(v)
		h.Sort()
		t.Log(h)
	}
}
