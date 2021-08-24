package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums         []int
		target, want int
	}{
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			nums:   []int{3, 6, 9, 11, 12, 13, 14, 16, 21},
			target: 3,
			want:   0,
		},
		{
			nums:   []int{3, 6, 9, 11, 12, 13, 14, 16, 21},
			target: 13,
			want:   5,
		},
	}

	for _, tt := range tests {
		if v := search(tt.nums, tt.target); tt.want != v {
			t.Log("failure ", v, ", but want ", tt.want)
		} else {
			t.Log("success ", v, ", and want ", tt.want)
		}
	}
}
