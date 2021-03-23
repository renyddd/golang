package main

import "testing"

func Test_robDfs(t *testing.T) {
	l := []int{104,209,137,52,158,67,213,86,141,110,151,127,238,147,169,138,240,185,246,225,147,203,83,83,131,227,54,78,165,180,214,151,111,161,233,147,124,143}
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{[]int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "t2",
			args: args{[]int{2, 7, 9, 3, 1}},
			want: 12,
		},
		{
			name: "bag",
			args: args{l},
			want: rob2(l),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robDfs(tt.args.nums); got != tt.want {
				t.Errorf("robDfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robDfs2(t *testing.T) {
	l := []int{104,209,137,52,158,67,213,86,141,110,151,127,238,147,169,138,240,185,246,225,147,203,83,83,131,227,54,78,165,180,214,151,111,161,233,147,124,143}
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{[]int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "t2",
			args: args{[]int{2, 7, 9, 3, 1}},
			want: 12,
		},
		{
			name: "bag",
			args: args{l},
			want: rob2(l),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robDfs2(tt.args.nums); got != tt.want {
				t.Errorf("robDfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
