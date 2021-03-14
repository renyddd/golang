package main

import (
	"fmt"
	"testing"
)

var n = 169999

func TestF(t *testing.T) {
	res := mySqrt(n)
	fmt.Println(res)
}

func TestF2(t *testing.T) {
	res := mySqrt2(n)
	fmt.Println(res)
}
