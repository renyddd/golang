package main

import "testing"

func TestRemove(t *testing.T) {
	nums := [][]int{
		{1,2,3,4,3,2,3,2,3},
		{1,4,52,5,6,3,4,56,7},
		{3,4,2,5,13,6,7,32,12,3,4,3},
	}

	for _, v := range nums {
		n := removeElement(v, 3)
		v = v[:n]
		t.Log(v)
	}
}