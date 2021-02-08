package main

/*
前序遍历 preorder = [ 3,         9,      20,          15,         7]
	             preRoot    preLeft   preRight

中序遍历 inorder =  [ 9,         3,        15,        20,         7]
                 inLeft      inRoot     inRight
    3
   / \
  9  20
    /  \
   15   7

preorder { [ root ]  [ == left == ] [ == right == ] }

inorder  { [ == left == ]  [ root ] [ == right == ] }
                             i
 */

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

// 假设没有重复元素

func buildTree(preorder []int, inorder []int) *TreeNode {
	tlen := len(preorder)
	if tlen == 0 {
		return nil
	}

	root_value := preorder[0]
	root_inorder_index := 0
	for i := range inorder {
		if inorder[i] == root_value {
			root_inorder_index = i
		}
	}

	// 有了 inorder 的左子树长度即可
	left_len := len(inorder[:root_inorder_index])
	root := &TreeNode{inorder[root_inorder_index], nil, nil}

	// 记得避开 preorder 中的 0 索引即跟节点；
	root.Left = buildTree(preorder[1:1+left_len], inorder[:root_inorder_index])
	root.Right = buildTree(preorder[1+left_len:], inorder[root_inorder_index+1:])

	return root
}