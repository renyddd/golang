package treenode

import (
	"testing"
)

func TestBuildTree(t *testing.T) {


	t1 := BuildA163TmpTree()
	t2 := BuildA163TmpTree()
	t.Log(t1, t2)
}