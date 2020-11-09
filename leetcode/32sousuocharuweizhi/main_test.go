package main

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1,3,5,6}
	n := map[int]int{
		5: 2,
		2: 1,
	}

	for k, expect := range n {
		res := searchInsert(nums, k)
		if expect != res {
			t.Errorf("%d: expect %d, but %d", k, expect, res)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1,3,5,6,7,8,10}
	n := map[int]int{
		5: 2,
		9: -1,
		2: -1,
		10: 6,
		1: 0,
		0: -1,
	}

	for k, expect := range n {
		res := bSearch(nums, k)
		if expect != res {
			t.Errorf("%d: expect %d, but %d", k, expect, res)
		}
	}
}