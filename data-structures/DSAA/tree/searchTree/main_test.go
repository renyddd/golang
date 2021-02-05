package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	root := &TreeNode{ElementType(rand.Intn(1000)), nil, nil}
	for i := 0; i < 1000; i++ {
		rd := rand.Intn(1000)
		root = InsertIteration(ElementType(rd), root)
	}
	PrintTree(root)
	t.Log()
	t.Log(FindMin(root))
	t.Log(FindMax(root))
	t.Log(Find(ElementType(433), root))
}

func TestDelete(t *testing.T) {
	root := &TreeNode{ElementType(9), nil, nil}
	Insert(ElementType(1), root)
	Insert(ElementType(123), root)
	Insert(ElementType(7), root)
	PrintTree(root)
	fmt.Println()

	root = Delete(ElementType(7), root)
	PrintTree(root)
}