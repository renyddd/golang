package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		data      []int
		key, want int
	}{
		{
			data: []int{1, 1, 1},
			key:  2,
			want: 2,
		},
		{
			data: []int{1, 1, 1, 1},
			key:  2,
			want: 3,
		},
		{
			data: []int{1, -1, 0},
			key:  0,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Log("res: ", subarraySum(tt.data, tt.key))
		t.Log("want: ", tt.want)
	}
}
