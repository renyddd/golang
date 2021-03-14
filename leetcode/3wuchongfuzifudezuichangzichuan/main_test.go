package main

import (
	"fmt"
	"testing"
)

func TestFindLongest(t *testing.T) {
	ss := []string{
		"aaaabcaa",
		"fdafijueolk",
		"ioujekdiwour",
		"dskfjasdlkfweiusdlfjahsdfjlaksdjfdshfueufdsflkasjdfkldsufidddodsjfwejfiuqweoihdncvbcxvsajdhfjlkdsufoieuwfhds",
	}

	for _, s := range ss {
		res := FindLongest(s)
		fmt.Println(res)
	}
}

func TestFindLongest2(t *testing.T) {
	ss := []string{
		"aaaabcaa",
		"fdafijueolk",
		"ioujekdiwour",
		"dskfjasdlkfweiusdlfjahsdfjlaksdjfdshfueufdsflkasjdfkldsufidddodsjfwejfiuqweoihdncvbcxvsajdhfjlkdsufoieuwfhds",
	}

	for _, s := range ss {
		res := FindLongest2(s)
		fmt.Println(string(res))
	}
}
