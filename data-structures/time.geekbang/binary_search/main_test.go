package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 7, 10, 14, 16, 21, 24, 33, 44, 55, 66, 78}

	res := BinarySearch(a, 55)

	fmt.Println(res)
}

func TestBinarySearch2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 7, 10, 14, 16, 21, 24, 33, 44, 55, 66, 78}

	res := BinarySearchRecusion(a, 55)

	fmt.Println(res)
}