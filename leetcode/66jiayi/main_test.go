package main

import "testing"

func TestPlusOne(t *testing.T) {
	nn := [][]int{
		{1,2,3},
		{1,1,9},
		{9},
		{9,9,9},
		{0},
		{},
	}

	for _, v := range nn {
		t.Log(plusOne(v))
	}
}

func TestPlusOneRecursion(t *testing.T) {
	nn := [][]int{
		{1,2,3},
		{1,1,9},
		{9},
		{9,9,9},
		{0},
		{},
	}

	for _, v := range nn {
		t.Log(plusOneRecursion(v))
	}
}