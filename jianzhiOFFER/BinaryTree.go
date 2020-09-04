package main

import "fmt"

// 前序遍历：根 - 左 - 右
// 中序遍历：左 - 根 - 右
// 后序遍历：左 - 右 - 根

// 题目描述：输入某二叉树的前序和中序遍历结果，请重建该二叉树
// 假设前、中序遍历中都不含有重复数字

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	前序遍历中，第一个数字总是树的根节点的值
	中序遍历中，根节点的值在输入数值中介，并且其左右为左右子树的值
	因此再扫描中序遍历结果时，就可知根节点的位置
	将中序遍历所得的根节点左右子节点的数量带入前序遍历中，就可找到相应的子序列
	再用同样的方法来构建左右子树，即递归
*/

func main() {
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	in := []int{4, 7, 2, 1, 5, 3, 6, 8}

	root := ReConstructBinaryTree(pre, in)

	PrintInOrder(root)
}

// golang 实现代码参考：
// https://github.com/DinghaoLI/Coding-Interviews-Golang/blob/master/006-%E9%87%8D%E5%BB%BA%E4%BA%8C%E5%8F%89%E6%A0%91/problem006.go

func ReConstructBinaryTree(pre []int, in []int) *TreeNode {
	// 两个遍历节点数不相同直接返回空
	if len(pre) != len(in) || len(pre) == 0 {
		return nil
	}

	// 根节点为前序遍历的第一个
	rootVal := pre[0]
	rootIndex := 0

	// 在中序遍历中确定根节点的位置
	for i := 0; i < len(in); i++ {
		if in[i] == rootVal {
			rootIndex = i
		}
	}

	// 此时的 rootIndex 对于中序遍历来说，可表示其根节点的位置索引
	// 并可直接在中序遍历中用来分割构造子树
	inL, inR := in[:rootIndex], in[rootIndex+1:]

	// 而对于前序遍历来说 rootindex 仅仅是
	// 根节点亦为第一个元素后所代表的偏移量
	preL, preR := pre[1:rootIndex+1], pre[rootIndex+1:]

	// fmt.Println("inL: ", inL, " -- inR: ", inR)
	// fmt.Println("preL: ", preL, " -- preR: ", preR)

	left := ReConstructBinaryTree(preL, inL)
	right := ReConstructBinaryTree(preR, inR)

	return &TreeNode{Val: rootVal, Left: left, Right: right}
}

func PrintPreOrder(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		PrintPreOrder(root.Left)
		PrintPreOrder(root.Right)
	}
}

func PrintInOrder(root *TreeNode) {
	if root != nil {
		PrintInOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		PrintInOrder(root.Right)
	}
}
