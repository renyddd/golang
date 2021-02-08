package main

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

/* 与 105 同样的处理套路
   由 inorder 切片确定左右子树的长度
   由 postorder 切片确定 root 跟节点的索引
 */

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}

	rootValue := postorder[len(postorder)-1]
	rootIndex := 0
	for i := range inorder {
		if inorder[i] == rootValue {
			rootIndex = i
		}
	}

	left_len := len(inorder[:rootIndex])
	root := &TreeNode{rootValue, nil, nil}
	root.Left = buildTree(inorder[:left_len], postorder[:left_len])
	root.Right = buildTree(inorder[left_len+1:], postorder[left_len:len(postorder)-1])

	return root
}