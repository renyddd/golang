package main

import "testing"

func TestDoSum(t *testing.T) {
	nums := []int{1, 2, 3}
	res := doSum(nums)

	t.Log(res)
}
