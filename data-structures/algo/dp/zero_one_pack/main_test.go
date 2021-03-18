package main

import "testing"

var (
	weights      = []int{5, 6, 7, 9, 11}
	volums       = []int{2, 4, 3, 5, 6}
	bagMaxVloume = 9
)

func TestZeroOnePack(t *testing.T) {
	t.Log(ZeroOnePack(volums, weights, bagMaxVloume))
}

func TestZeroOnePackCommented(t *testing.T) {
	t.Log(ZeroOnePackCommented(volums, weights, bagMaxVloume))
}

func TestZeroOnePackOneSlice(t *testing.T) {
	t.Log(ZeroOnePackOneSlice(volums, weights, bagMaxVloume))
}

func BenchmarkZeroOnePack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroOnePack(volums, weights, bagMaxVloume)
	}
}

func BenchmarkZeroOnePackOneSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZeroOnePackOneSlice(volums, weights, bagMaxVloume)
	}
}

/*
BenchmarkZeroOnePack
BenchmarkZeroOnePack-12                  4812624               249 ns/op
BenchmarkZeroOnePackOneSlice
BenchmarkZeroOnePackOneSlice-12         18294130                65.9 ns/op
*/
