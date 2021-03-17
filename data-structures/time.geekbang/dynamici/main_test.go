package main

import "testing"

var weight = []int{2, 2, 4, 6, 3, 7, 45, 7, 6, 9, 2, 3, 1, 4, 5, 7, 8, 6, 7, 5, 4, 5, 6, 2}
var BAGMAXWEIGHT = 55

func TestBacktracking(t *testing.T) {
	res := Backtracking(weight, BAGMAXWEIGHT)
	t.Log(res)
}

func BenchmarkBacktracking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Backtracking(weight, BAGMAXWEIGHT)
	}
}

func TestBacktracking2(t *testing.T) {
	res := Backtracking2(weight, BAGMAXWEIGHT)
	t.Log(res)
}

func BenchmarkBacktracking2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Backtracking2(weight, BAGMAXWEIGHT)
	}
}

//func TestDynamicprogramming(t *testing.T) {
//	res := Dynamicprogramming(weight, BAGMAXWEIGHT)
//	t.Log(res)
//}
//
//func BenchmarkDynamicprogramming(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Dynamicprogramming(weight, BAGMAXWEIGHT)
//	}
//}
