package main

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()

	// 还是无法避免这样的创建
	s := &single{}
	fmt.Println(s)
}

func TestGetInstance2(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstance2()
	}
	fmt.Scanln()
}
