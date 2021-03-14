package main

import "testing"

func TestTraverse(t *testing.T) {
	type args struct {
		a [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				a: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
