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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree(preorder[1:i+1], inorder[0:i-1])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}