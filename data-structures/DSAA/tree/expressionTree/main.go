package main

import "fmt"

type TreeNode struct {
	Value       rune
	Left, Right *TreeNode
}

func BuildExpressionTreeFromString(expression string) *TreeNode {
	stack := make([]*TreeNode, 0)
	for _, c := range expression {
		if c != '+' && c != '-' && c != '*' && c != '/' {
			stack = append(stack, &TreeNode{c, nil, nil})
		} else {
			l, r := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := &TreeNode{Value: c, Left: l, Right: r}
			stack = append(stack, res)
		}
	}

	// 最后的运算结果将保存在栈底
	return stack[0]
}

func PrintTree(root *TreeNode) {
	if root != nil {
		PrintTree(root.Left)
		fmt.Printf("%c", root.Value)
		PrintTree(root.Right)
	}
}

func main() {
	res := BuildExpressionTreeFromString("ab+cde+**")
	PrintTree(res)
	// TODO: 如何处理计算优先级以及括号问题？
}
