package main

import "testing"

var (
	Vomules = []int{2, 3, 5}
	Weights = []int{3, 5, 6}
	MaxBag  = 12
)

func TestCompletePack(t *testing.T) {
	t.Log(CompletePack(Vomules, Weights, MaxBag))
}

func TestCompletePackBackTrack(t *testing.T) {
	type args struct {
		V    []int
		W    []int
		Vbag int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{V: []int{2, 4, 3, 4}, W: []int{5, 6, 7, 7}, Vbag: 7},
			want: CompletePack([]int{2, 4, 3, 4}, []int{5, 6, 7, 7}, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompletePackBackTrack(tt.args.V, tt.args.W, tt.args.Vbag); got != tt.want {
				t.Errorf("CompletePackBackTrack() = %v, want %v", got, tt.want)
			}
		})
	}
}
