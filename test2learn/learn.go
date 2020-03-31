package main

import (
	"fmt"
	"os"
)

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func main() {
	for i := 1; i < 5; i = i + 1 {
		fmt.Printf("learn go.\n")
	}

	var timeZone = map[string]int{
		"UTC": 111,
		"EST": 222,
	}
	fmt.Printf("%s\n", timeZone["UTC"])

	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	var (
		home = os.Getenv("GOPATH")
	)
	fmt.Println(home)
}
