package treenode

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Tree_traversal#Pre-order_(NLR)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InOrderTraversal
//Traverse the left subtree by recursively calling the in-order function.
//Access the data part of the current node.
//Traverse the right subtree by recursively calling the in-order function.
func (t *TreeNode) InOrderTraversal() {
	if t != nil {
		t.Left.InOrderTraversal()
		fmt.Println(t.Val)
		t.Right.InOrderTraversal()
	}
}

// PreOrderTraversal
func (t *TreeNode) PreOrderTraversal() {
	if t != nil {
		fmt.Println(t.Val)
		t.Left.PreOrderTraversal()
		t.Right.PreOrderTraversal()
	}
}

func BuildA163TmpTree() *TreeNode {
	l1 := &TreeNode{1, nil, nil}
	r1 := &TreeNode{3, nil, nil}
	head := &TreeNode{
		Val:   6,
		Left:  l1,
		Right: r1,
	}
	return head
}

func BuildA213TmpTree() *TreeNode {
	l1 := &TreeNode{2, nil, nil}
	r1 := &TreeNode{3, nil, nil}
	head := &TreeNode{
		Val:   1,
		Left:  l1,
		Right: r1,
	}
	return head
}

func BuildSymmetric811Tree() *TreeNode {
	l1 := &TreeNode{1, nil, nil}
	r1 := &TreeNode{1, nil, nil}
	head := &TreeNode{8, l1, r1}
	return head
}
