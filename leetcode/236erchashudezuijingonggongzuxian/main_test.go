package main

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	left := &TreeNode{2, nil, nil}
	right := &TreeNode{3, nil, nil}
	root := &TreeNode{1, left, right}
	LowestCommonAncestor(root, left, right)
}