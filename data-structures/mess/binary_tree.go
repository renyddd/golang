/*
-  https://github.com/wangzheng0822/algo/blob/master/go/24_tree/TreeNode.go
-  https://github.com/wangzheng0822/algo/blob/master/go/24_tree/BinaryTree.go
-  https://github.com/wangzheng0822/algo/blob/master/go/23_binarytree/binarytree.go
-  https://github.com/renyddd/embedded-Linux-programming/blob/master/dataStructure/treeDir/BTNode.h
*/
package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func (n *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", n.data, n.left, n.right)
}

func CreateBTree(arr []int) *Node {
	tree := NewNode(arr[0]) // 树的根节点
	var tree_node, process_node *Node
	/*
		tree_node: 新生成的承载数据的待添加节点
		process_node: 二层循环中 p 的前驱
	*/

	for i := 1; i < len(arr); i++ {
		process_node = NewNode(arr[i])
		var p *Node = tree // 节点循环变量，从根开始寻找
		for p != nil {
			tree_node = p
			if process_node.data < p.data {
				p = p.left
			} else {
				p = p.right
			}
		}

		if process_node.data < tree_node.data {
			tree_node.left = process_node
		} else {
			tree_node.right = process_node
		}
	}

	return tree
}

func FOrder(root *Node) {
	if root != nil {
		fmt.Println(root.data)
		FOrder(root.left)
		FOrder(root.right)
	}
}

func InOrder(root *Node) {
	if root != nil {
		InOrder(root.left)
		fmt.Println(root.data)
		InOrder(root.right)
	}
}

func main() {
	num := []int{4, 2, 6, 3, 9, 8, 7, 5, 12, 54, 234, 65, 346, 3463}

	tree := CreateBTree(num)
	InOrder(tree)
}
