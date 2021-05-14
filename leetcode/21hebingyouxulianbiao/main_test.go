package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestUseful(t *testing.T) {
	ss := []string{
		"hello",
		"world",
		"sixuan",
		"bob",
	}
	t.Log(ss)
	sort.Strings(ss)
	t.Log(ss)
}

func TestMoreTimes(t *testing.T) {
	MergeTimes(10)
}

func MergeTimes(n int) {

	for i := 0; i < n; i++ {

		rand.Seed(time.Now().Unix() + int64(i))
		n1 := elem{}
		n2 := elem{}

		for i := 0; i < 10; i++ {
			n1 = append(n1, rand.Intn(100))
		}
		sort.Sort(n1)
		h1 := MakeListNode(n1)
		for i := 0; i < 10; i++ {
			n2 = append(n2, rand.Intn(140))
		}
		sort.Sort(n2)
		h2 := MakeListNode(n2)

		newhead := mergeTwoLists2(h1, h2)
		newhead.PrintList()
	}
}
