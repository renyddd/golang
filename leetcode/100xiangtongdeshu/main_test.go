package main

import (
	"../treenode"
	"testing"
)

//func TestIsSameTree(t *testing.T) {
//	t1 := treenode.BuildA163TmpTree()
//	t2 := treenode.BuildA163TmpTree()
//
//	t.Log(isSameTreeDepthFirst(t1, t2))
//}

func TestAppendNil(t *testing.T) {
	n := []*treenode.TreeNode{
		new(treenode.TreeNode),
	}
	t.Log(n)
	n = append(n, &treenode.TreeNode{})
	t.Log(n)
	t.Log(len(n))
	n = append(n, nil)
	t.Log(n)
	t.Log(len(n))

}
