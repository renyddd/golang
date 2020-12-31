package main

import "fmt"

// https://github.com/wangzheng0822/algo/blob/master/go/12_sorts/QuickSort.go
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	fmt.Println("when pivot is: ", i, " --> ", arr)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func main() {
	a := []int{3, 2, 5, 8, 4, 7, 6, 9, 1, 0}
	fmt.Println("original: ", a)
	QuickSort(a)
	fmt.Println("result: ", a)
}
