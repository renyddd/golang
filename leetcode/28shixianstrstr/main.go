package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello, world"
	needle := "orl"

	fmt.Println(strStr(s, needle))
	fmt.Println(strings.Index(s, needle))
}

func strStr(haystack string, needle string) int {
	L, l := len(haystack), len(needle)

	if l == 0 {
		return 0
	}
	if L == 0 || L < l {
		return -1
	}

	for i := 0; i < L-l+1; i++ {
		if haystack[i:i+l] == needle {
			return i
		}
	}
	return -1
}
