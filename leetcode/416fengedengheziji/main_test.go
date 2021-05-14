package main

import "testing"

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				nums: []int{1, 5, 11, 5},
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				[]int{1, 2, 4, 5},
			},
			want: true,
		},
		{
			name: "t3",
			args: args{
				[]int{1, 2, 3, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
