package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	m1 := make(map[int]int)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 999}
	for i, _ := range nums {
		m1[nums[i]] = DataOut(nums[i])
	}

	for k := range m1 {
		if res := climbStairs(k); m1[k] != res {
			t.Errorf("except %d, but %d\n", m1[k], res)
		}
	}

}
