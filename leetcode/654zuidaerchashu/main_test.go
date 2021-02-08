package main

import "testing"

func TestCopy(t *testing.T) {
	nums := []int{1,2,3,4,5,6,7}

	left := make([]int, len(nums))
	right := make([]int,len(nums))
	index := 3

	l_l := copy(left, nums[:index])
	l_r := copy(right, nums[index:])

	t.Log(left, l_l)
	t.Log(right, l_r)
}

func TestContruct(t *testing.T) {
	ns := []int{4,2,62,5,7,8,35,6}
	root := constructMaximumBinaryTree(ns)
	t.Log(root)
}

func TestIndex(t *testing.T) {
	nums := []int{10, 20, 30, 40}
	nums = []int{10}
	index := 1
	t.Log(nums[:index])
	t.Log(nums[index+1:])
}