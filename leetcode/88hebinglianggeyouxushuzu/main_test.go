package main

import (
	"fmt"
	"testing"
)

type ns struct {
	n1, n2 []int
}

func TestMerge(t *testing.T) {
	nn := []ns{
		ns{[]int{1,2,4,5}, []int{4,5,6}},
		ns{[]int{1,4,5}, []int{4,5,6}},
	}

	for _, v := range nn {
		t.Log(v.n1[:3])
	}
}

func TestDiySort( t *testing.T) {
	nn := [][]int{
		[]int{3,4,6,1,3,4,61,45},
		[]int{3,4,5,145,161,45,321,31,231,1,1,2},
		[]int{99,8,7,6,5,4,3,2,1},
	}

	dsort := func(nums1 []int) {
		for i := 0; i < len(nums1); i++ {
			for j := 0; j < len(nums1)-i-1; j++ {
				if nums1[j] > nums1[j+1] {
					nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
					fmt.Printf("switch %d & %d\t", nums1[j], nums1[j+1])
				}
			}
		}
		fmt.Println(nums1)
	}

	for _, v := range nn {
		dsort(v)
	}
}