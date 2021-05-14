package bs

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
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
	nums := []int{1, 3, 5, 6, 7, 8, 10}
	n := map[int]int{
		5:  2,
		9:  -1,
		2:  -1,
		10: 6,
		1:  0,
		0:  -1,
	}

	for k, expect := range n {
		res := bSearch(nums, k)
		if expect != res {
			t.Errorf("%d: expect %d, but %d", k, expect, res)
		}
	}
}

func TestBsearchFirst(t *testing.T) {
	nums := []int{1, 2, 6, 6, 8, 9, 89}
	ref := map[int]int{
		1: 0,
		6: 2,
		9: 5,
	}

	for k, except := range ref {
		res := bSearchFirst(nums, k)
		if res != except {
			t.Errorf("%d: except %d, but %d", k, except, res)
		}
	}
}

func TestBsearchLast(t *testing.T) {
	nums := []int{1, 2, 6, 7, 7, 7, 8, 8, 9, 9}
	ref := map[int]int{
		2: 1,
		6: 2,
		7: 5,
		8: 7,
		9: 9,
	}

	for k, expect := range ref {
		res := bSearchLast(nums, k)
		if res != expect {
			t.Errorf("%d: except %d, but %d", k, expect, res)
		}
	}
}

func TestBsearchFBigger(t *testing.T) {
	nums := []int{1, 2, 6, 7, 8, 9, 13, 14}
	ref := map[int]int{
		2:  1,
		5:  2,
		19: -1,
		10: 6,
	}

	for k, expect := range ref {
		res := bSearchFBigger(nums, k)
		if res != expect {
			t.Errorf("%d: except %d, but %d", k, expect, res)
		}
	}
}

func TestBsearchLSmaller(t *testing.T) {
	nums := []int{1, 2, 6, 7, 8, 9, 13, 14, 20}
	ref := map[int]int{
		1:   0,
		5:   1,
		15:  7,
		203: 8,
		0:   -1,
		-3:  -1,
	}

	for k, expect := range ref {
		res := bSearchLSmaller(nums, k)
		if res != expect {
			t.Errorf("%d: expect %d, but %d", k, expect, res)
		}
	}
}
