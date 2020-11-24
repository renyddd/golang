package main

import (
	"../treenode"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	root := treenode.BuildA163TmpTree()
	t.Log(isSymmetric(root))

	t.Log(isSymmetric(treenode.BuildSymmetric811Tree()))
}