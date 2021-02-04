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

func BuildTree1() {
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

type TreeNodeWithWeight struct {
	Name string
	Value int
	// 完全可以通过 FirstChild 的值是否为空来判断该节点的类型，目录还是文件
	FirstChild, NextSibling *TreeNodeWithWeight
}

func SizeDirectory(D *TreeNodeWithWeight) int {
	var TotalSize int
	if D != nil {
		TotalSize += D.Value
		if D.FirstChild != nil {
			for p := D.FirstChild; p != nil; p = p.NextSibling {
				TotalSize += SizeDirectory(p)
			}
		}
	}
	return TotalSize
}

func BuildTree2() {
	l3n0 := &TreeNodeWithWeight{"l3n0", 10, nil, nil}
	l1n2 := &TreeNodeWithWeight{"l1n2", 10, nil,nil}
	l1n1 := &TreeNodeWithWeight{"l1n1", 6, nil, l1n2}
	l2n0 := &TreeNodeWithWeight{"l2n0", 5, l3n0, nil}
	l1n0 := &TreeNodeWithWeight{"l1n0", 1, l2n0, l1n1}
	root := &TreeNodeWithWeight{"root", 1, l1n0, nil}

	fmt.Println(SizeDirectory(root))
}

func main() {
	BuildTree2()
}