package main

import "testing"

func TestStringIndex(t *testing.T) {
	s := "abcdcba"

	v := s[3]
	vs := s[:]

	t.Log(string(v))
	t.Log(vs)
}

func TestLongestPalindrome(t *testing.T) {
	s := "abcdcba"
	t.Log(longestPalindrome(s))
}
