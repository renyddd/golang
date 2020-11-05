package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestPalindrome(t *testing.T) {

	var tests = []struct {
		input int
		want bool
	}{
		{121, true},
		{0, true},
		{-1, false},
		{-2, true},
		{33233, false},
	}

	for _, test := range tests {
		if got := isPalindrome(test.input); got != test.want {
			t.Error(test)
		}
	}
}

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 10; i++ {
		p := rng.Intn(9999999)
		t.Error(p, isPalindrome(p))
	}
}