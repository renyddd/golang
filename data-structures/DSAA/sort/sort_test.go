package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{2, 5, 3, 7, 9, 8, 1, 6, 4, -5, 99, 0, -2}
	InsertionSort(a)
	fmt.Println(a)
}

func TestBubbleSort(t *testing.T) {
	a := []int{2, 5, 3, 7, 9, 8, 1, 6, 4, -5, 99, 0, -2}
	BubbleSort(a)
	fmt.Println(a)
}

func TestMergeSort(t *testing.T) {
	a := []int{2, 5, 3, 7, 9, 8, 1, 6, 4, -5, 99, 0, -2}
	fmt.Println("original:", a)
	MergeSort(a)
	fmt.Println(a)
}

func TestQuickSort(t *testing.T) {
	a := []int{2, 5, 3, 7, 9, 8, 1, 6, 4, -5, 99, 0, -2}
	fmt.Println("original:", a)
	QuickSort(a)
	fmt.Println(a)
}
