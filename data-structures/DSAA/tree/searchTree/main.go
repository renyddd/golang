package main

import "fmt"

type ElementType int

type TreeNode struct {
	Element     ElementType
	Left, Right *TreeNode
}

type SearchTree *TreeNode
type Position *TreeNode

func MakeEmpty(T SearchTree) SearchTree {
	if T != nil {
		MakeEmpty(T.Left)
		MakeEmpty(T.Right)
		T = nil
	}
	return nil
}

func Find(X ElementType, T SearchTree) Position {
	if T == nil {
		return nil
	}
	if X < T.Element {
		return Find(X, T.Left)
	} else if X > T.Element {
		return Find(X, T.Right)
	} else {
		return Position(T)
	}
}

func FindMin(T SearchTree) Position {
	if T == nil {
		return nil
	}
	if T.Left == nil {
		return Position(T)
	} else {
		return FindMin(T.Left)
	}
}

func FindMax(T SearchTree) Position {
	if T != nil {
		p := T
		for ; p.Right != nil; p = p.Right {
		}
		return Position(p)
	} else {
		return nil
	}
}

func Insert(X ElementType, T SearchTree) SearchTree {
	if T == nil {
		return &TreeNode{X, nil, nil}
	}
	p := T
	for p != nil {
		if p.Element == X {
			break
		} else if X < p.Element {
			if p.Left != nil {
				p = p.Left
			} else {
				p.Left = &TreeNode{X, nil, nil}
			}
		} else if X > p.Element {
			if p.Right != nil {
				p = p.Right
			} else {
				p.Right = &TreeNode{X, nil, nil}
			}
		}
	}
	return T
}

func InsertIteration(X ElementType, T SearchTree) SearchTree {
	if T == nil {
		return &TreeNode{X, nil, nil}
	}
	if X == T.Element {
		return T
	} else if X < T.Element {
		T.Left =  InsertIteration(X, T.Left)
	} else {
		T.Right = InsertIteration(X, T.Right)
	}
	return T
}

func Delete(X ElementType, T SearchTree) SearchTree {
	if T == nil {
		return nil
	} else if X < T.Element {
		// go left
		Delete(X, T.Left)
	} else if X > T.Element {
		// go right
		Delete(X, T.Right)
	} else // Found
	if T.Left != nil && T.Right != nil { // two children
		// replace with the smallest in the right subtree
		tmp := FindMin(T.Right)
		T.Element = tmp.Element
		Delete(X, T.Right)
	} else { // one or zero child
		if T.Left == nil {
			T = T.Right
		} else if T.Right == nil {
			T = T.Left
		}
	}

	return T
}

func PrintTree(root SearchTree) {
	if root != nil {
		PrintTree(root.Left)
		fmt.Printf("%v\t", root.Element)
		PrintTree(root.Right)
	}
}

func LookSearchTree1() {
	root := &TreeNode{4, nil, nil}
	newT := Insert(5, root)
	newT = Insert(1, newT)
	newT = Insert(2, newT)
	newT = Insert(1, newT)
	newT = Insert(3, newT)
	newT = Insert(8, newT)
	newT = Insert(7, newT)
	newT = Insert(6, newT)

	PrintTree(newT)
}

func LookSearchTree2() {
	root := &TreeNode{4, nil, nil}
	newT := InsertIteration(5, root)
	newT = InsertIteration(1, newT)
	newT = InsertIteration(2, newT)
	newT = InsertIteration(9, newT)
	newT = InsertIteration(3, newT)
	newT = InsertIteration(8, newT)
	newT = InsertIteration(7, newT)
	newT = InsertIteration(6, newT)

	PrintTree(newT)
}

func main() {
	LookSearchTree2()
}