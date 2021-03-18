package main

import "testing"

func TestC(t *testing.T) {
	amount := 7
	coins := []int{1, 2, 5}
	t.Log(change(amount, coins))
}
