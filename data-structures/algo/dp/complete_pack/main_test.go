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
