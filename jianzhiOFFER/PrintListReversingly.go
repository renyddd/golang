package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n3 := &ListNode{4, nil}
	n2 := &ListNode{3, n3}
	n1 := &ListNode{2, n2}
	head := &ListNode{1, n1}

	PrintList(head)
	fmt.Println("------")
	PrintListReversingly_Recursively(head)
}

// 非递归的实现
func PrintListReversingly_Iteratively(head *ListNode) {
	// 用数组代替一个栈的实现
	if head == nil {
		return
	}

	var vnum = make([]int, 0)

	for p := head; p != nil; p = p.Next {
		vnum = append(vnum, p.Val)
	}

	for i := len(vnum) - 1; i > -1; i-- {
		fmt.Println(vnum[i])
	}
}

// 递归实现
// 既然非递归的形式所使用的为栈结构，那这不久等于是在递归嘛
func PrintListReversingly_Recursively(head *ListNode) {
	if head == nil {
		return
	}

	if head == nil {
		return
	} else {
		PrintListReversingly_Recursively(head.Next)
		fmt.Println(head.Val)
	}
}

func PrintList(head *ListNode) {
	if head == nil {
		return
	}

	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
