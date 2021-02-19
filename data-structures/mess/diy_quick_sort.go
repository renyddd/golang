package main

import "fmt"

// reference - https://github.com/wangzheng0822/algo/blob/master/go/12_sorts/QuickSort.go
func partition_d(a []int, start, end int) int {
	left, right := start, end
	pivot := end
	for left < right {
		if a[left] > a[pivot] {
			if a[right] < a[pivot] {
				a[left], a[right] = a[right], a[left]
			} else {
				right--
			}
		} else {
			left++
		}
	}
	if a[left] != a[pivot] {
		a[left], a[pivot] = a[pivot], a[left]
	}
	fmt.Printf("array order is: %d --> ", a)
	return left
}

func separateSort_d(a []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition_d(a, start, end)
	fmt.Println("the pivot is:", a[pivot])
	separateSort_d(a, start, pivot-1)
	separateSort_d(a, pivot+1, end)
}

func qsort(a []int) {
	separateSort_d(a, 0, len(a)-1)
}

func main() {
	a := []int{6, 2, 5, 9, 7, 0, 1, 3, 4, 8}
	//a := []int{3, 9, 5, 3, 2, 1, 7}
	fmt.Println("Original: ", a)
	qsort(a)
	fmt.Println("Complete! ", a)
}
