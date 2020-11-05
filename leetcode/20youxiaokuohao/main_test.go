package main

import "testing"

func TestValid(t *testing.T) {
	s := []string{
		"(){}[]",
		"{()[]}",
		"[}[]]",
		"[",
		"((",
		"[[",
		"]]",
	}
	for _, v := range s {
		t.Log(isValid(v))
	}
}
