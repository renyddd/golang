package main

import "testing"

var (
	weights = []int{5, 6, 7}
	volums = []int{2, 4, 3}
)

func TestZeroOnePack(t *testing.T) {
	t.Log(ZeroOnePack(volums, weights, 7))
}