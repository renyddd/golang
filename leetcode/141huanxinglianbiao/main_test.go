package main

import (
	"../listnode"
	"testing"
)

func TestHasCycle(t *testing.T) {
	n1 := []int{1,2,3,4,5,6,7}

	h1 := listnode.MakeListNode(n1)
	h2 := listnode.MakeListNode(n1)
	h3 := listnode.MakeListNode(n1)
	h4 := listnode.MakeListNode(n1)

	BuildRingList(h1, 0)
	BuildRingList(h2, 4)
	BuildRingList(h3, 5)

	t.Log(HasCycle(h1))
	t.Log(HasCycle(h2))
	t.Log(HasCycle(h3))
	t.Log(HasCycle(h4))
}