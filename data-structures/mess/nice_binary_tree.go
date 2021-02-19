// src: https://github.com/floyernick/Data-Structures-and-Algorithms/blob/master/BinaryTree/BinaryTree.go
package main

import "fmt"

type RNode struct {
	data   int
	parent *RNode
	left   *RNode
	right  *RNode
}

type BinaryTree struct {
	root *RNode
}

func main() {
	root := &RNode{data: 666}
	tree := &BinaryTree{root}

	tree.InsertItem(777)
	tree.InsertItem(888)
	tree.InsertItem(555)

	InOrderR(tree.root)
	// 	fmt.Println(tree.root.data)
}

// try to understand callback
func InOrderR(root *RNode) {
	if root != nil {
		InOrderR(root.left)
		fmt.Println(root.data)
		InOrderR(root.right)
	}
}

func (tree *BinaryTree) InsertItem(i int) {
	if tree.root == nil {
		tree.root = &RNode{data: i}
		return
	}
	currentNode := tree.root
	for {
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &RNode{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &RNode{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *BinaryTree) SearchItem(i int) (*RNode, bool) {
	if tree.root == nil {
		return nil, false
	}
	currentNode := tree.root
	for currentNode != nil {
		if i == currentNode.data {
			return currentNode, true
		} else if i > currentNode.data {
			currentNode = currentNode.right
		} else if i < currentNode.data {
			currentNode = currentNode.left
		}
	}
	return nil, false
}

func (tree *BinaryTree) InorderTraversal(subtree *RNode, callback func(int)) {
	if subtree.left != nil {
		tree.InorderTraversal(subtree.left, callback)
	}
	callback(subtree.data)
	if subtree.right != nil {
		tree.InorderTraversal(subtree.right, callback)
	}
}

func (tree *BinaryTree) PreorderTraversal(subtree *RNode, callback func(int)) {
	callback(subtree.data)
	if subtree.left != nil {
		tree.PreorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PreorderTraversal(subtree.right, callback)
	}
}

func (tree *BinaryTree) PostorderTraversal(subtree *RNode, callback func(int)) {
	if subtree.left != nil {
		tree.PostorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PostorderTraversal(subtree.right, callback)
	}
	callback(subtree.data)
}
