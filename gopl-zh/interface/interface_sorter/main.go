package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2, 44, 5, 1, 4, 56, 1, 44, 1}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
