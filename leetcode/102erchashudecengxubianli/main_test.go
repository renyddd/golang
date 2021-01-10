package main

import "testing"

func TestTwoDimensionalSlice(t *testing.T) {
	var s [][]int
	t.Log(s)

	data := []int{1, 2, 1, 3, 3, 1, 2, 4, 5, 6}
	ms := make(map[int][]int, 0)

	for _, v := range data {
		ms[v] = append(ms[v], v*2)
	}

	t.Log(ms)

	for _, v := range ms {
		s = append(s, v)
	}

	t.Log(s)
}
