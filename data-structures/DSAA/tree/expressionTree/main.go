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

func BuildExpressionResult(expression string) int32 {
	stack := make([]*TreeNode, 0)
	for _, c := range expression {
		if c != '+' && c != '-' && c != '*' && c != '/' {
			stack = append(stack, &TreeNode{c, nil, nil})
		} else {
			var resVal int32
			l, r := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch c {
			case '+':
				resVal = l.Value + r.Value
			case '-':
				resVal = l.Value + r.Value
			case '*':
				resVal = l.Value + r.Value
			case '/':
				resVal = l.Value + r.Value
			}
			stack = append(stack, &TreeNode{resVal, nil, nil})
		}
	}

	return stack[0].Value
}

func printTree(root *TreeNode) {
	if root != nil {
		printTree(root.Left)
		fmt.Printf("%c", root.Value)
		printTree(root.Right)
	}
}

func PrintTree(root *TreeNode) {
	printTree(root)
	fmt.Println()
}

func main() {
	expression1 := "ab+cde+**"
	res := BuildExpressionTreeFromString(expression1)
	PrintTree(res)

	expression2 := "11+2++++++" // must panic
	expression2 = "11+1+1+"
	// res: 196, '1' (int32) for 49
	fmt.Println("res:", BuildExpressionResult(expression2))
	// TODO: 如何处理计算优先级以及括号问题？
}
