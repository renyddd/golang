package main

import (
	"strings"
	"testing"
)

func TestStrstr(t *testing.T) {
	s := map[string]string{
		"hello, world": "ld",
		"hello世界": "o世",
		"-3fdsik99j": "k9",
		"jfi291": "jf",
	}

	for k, v := range s {
		res := strStr(k, v)
		reference := strings.Index(k, v)
		t.Log(res, reference)
		if res != reference {
			t.Fail()
		}
	}
}