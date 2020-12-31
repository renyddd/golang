// src: https://github.com/floyernick/Data-Structures-and-Algorithms/blob/master/BinaryTree/BinaryTree.go
package main

import "fmt"

type Node struct {
	data   int
	parent *Node
	left   *Node
	right  *Node
}

type BinaryTree struct {
	root *Node
}

func main() {
	root := &Node{data: 666}
	tree := &BinaryTree{root}

	tree.InsertItem(777)
	tree.InsertItem(888)
	tree.InsertItem(555)

	InOrder(tree.root)
	// 	fmt.Println(tree.root.data)
}

// try to understand callback
func InOrder(root *Node) {
	if root != nil {
		InOrder(root.left)
		fmt.Println(root.data)
		InOrder(root.right)
	}
}

func (tree *BinaryTree) InsertItem(i int) {
	if tree.root == nil {
		tree.root = &Node{data: i}
		return
	}
	currentNode := tree.root
	for {
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &Node{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *BinaryTree) SearchItem(i int) (*Node, bool) {
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

func (tree *BinaryTree) InorderTraversal(subtree *Node, callback func(int)) {
	if subtree.left != nil {
		tree.InorderTraversal(subtree.left, callback)
	}
	callback(subtree.data)
	if subtree.right != nil {
		tree.InorderTraversal(subtree.right, callback)
	}
}

func (tree *BinaryTree) PreorderTraversal(subtree *Node, callback func(int)) {
	callback(subtree.data)
	if subtree.left != nil {
		tree.PreorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PreorderTraversal(subtree.right, callback)
	}
}

func (tree *BinaryTree) PostorderTraversal(subtree *Node, callback func(int)) {
	if subtree.left != nil {
		tree.PostorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PostorderTraversal(subtree.right, callback)
	}
	callback(subtree.data)
}
