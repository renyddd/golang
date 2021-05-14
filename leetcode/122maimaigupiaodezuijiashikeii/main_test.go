package main

import (
	"fmt"
	"testing"
)

func TestMax01(t *testing.T) {
	p := []int{1, 5, 3, 6, 8, 3, 4, 5, 9}
	fmt.Println(maxProfit(p))
	fmt.Println(maxProfit01(p))
}
