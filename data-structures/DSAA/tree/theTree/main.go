package main

import "fmt"

type ElementType string

type TreeNode struct {
	Value                      string
	Element                    ElementType
	FirstChildren, NextSibling *TreeNode
}

func ListDir(D *TreeNode, Depth int) {
	// D is a legitimate entry
	if D.Element != "" {
		for i := 0; i < Depth; i++ {
			fmt.Printf("\t")
		}
		fmt.Printf("%v, %v\n", D, Depth)
		if D.Element == "dir" {
			for p := D.FirstChildren; p != nil; p = p.NextSibling {
				ListDir(p, Depth+1)
			}
		}
	}
}

func ListDirectory(D *TreeNode) {
	ListDir(D, 0)
}

func main() {
	l2n2 := &TreeNode{"l2n2", "file", nil, nil}
	l3n0 := &TreeNode{"l3n0", "file", nil, nil}
	l2n1 := &TreeNode{"l2n1", "dir", l3n0, l2n2}
	l2n0 := &TreeNode{"l2n0", "file", nil, l2n1}
	l1n3 := &TreeNode{"l1n3", "file", nil, nil}
	l1n2 := &TreeNode{"l1n2", "file", nil, l1n3}
	l1n1 := &TreeNode{"l1n1", "dir", l2n0, l1n2}
	l1n0 := &TreeNode{"l1n0", "file", nil, l1n1}
	root := &TreeNode{"root", "dir", l1n0, nil}

	ListDirectory(root)
}
