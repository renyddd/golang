package main

import "testing"

//func Test_findTargetSumWays(t *testing.T) {
//	type args struct {
//		nums []int
//		S    int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//		{
//			name: "t1",
//			args: args{[]int{1, 1, 1, 1, 1}, 3},
//			want: 5,
//		},
//		{
//			name: "t2",
//			args: args{[]int{1}, 2},
//			want: 0,
//		},
//		{
//			name: "t3",
//			args: args{[]int{1}, 1},
//			want: 1,
//		},
//		{
//			name: "t4",
//			args: args{[]int{1, 0}, 1},
//			want: 2,
//		},
//		{
//			name: "t5",
//			args: args{[]int{1000}, 1000},
//			want: 1,
//		},
//		{
//			name: "t6",
//			args: args{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1},
//			want: 256,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := findTargetSumWaysFirstTy(tt.args.nums, tt.args.S); got != tt.want {
//				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_findTargetSumWaysBackTrack(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{[]int{1, 1, 1, 1, 1}, 3},
			want: 5,
		},
		{
			name: "t2",
			args: args{[]int{1}, 2},
			want: 0,
		},
		{
			name: "t3",
			args: args{[]int{1}, 1},
			want: 1,
		},
		{
			name: "t4",
			args: args{[]int{1, 0}, 1},
			want: 2,
		},
		{
			name: "t5",
			args: args{[]int{1000}, 1000},
			want: 1,
		},
		{
			name: "t6",
			args: args{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1},
			want: 256,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWaysEnumerate(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWaysBackTrack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCertainNums(t *testing.T) {
	nums := []int{1, 2, 1, 3, 1}
	S := 5
	t.Log(findTargetSumWaysEnumerate(nums, S))
}

func Test_findTargetSecondTry(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{[]int{1, 1, 1, 1, 1}, 3},
			want: 5,
		},
		{
			name: "t2",
			args: args{[]int{1}, 2},
			want: 0,
		},
		{
			name: "t3",
			args: args{[]int{1}, 1},
			want: 1,
		},
		{
			name: "t4",
			args: args{[]int{1, 0}, 1},
			want: 2,
		},
		{
			name: "t5",
			args: args{[]int{1000}, 1000},
			want: 1,
		},
		{
			name: "t6",
			args: args{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1},
			want: 256,
		},
		{
			name: "t7",
			args: args{[]int{1000}, -1000},
			want: 1,
		},
		{
			name:"t8",
			args: args{[]int{1}, -1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWaysSecondTry(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
