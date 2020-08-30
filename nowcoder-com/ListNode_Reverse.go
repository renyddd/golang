package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l3 := &ListNode{6, nil}
	l2 := &ListNode{3, l3}
	l1 := &ListNode{2, l2}
	head := &ListNode{1, l1}

	newHead := ReverseList(head)

	PrintList(newHead)
}

// 空间复杂度为 O(1)
// 由 cur 开始遍历整个链表
// nex 预存储下次处理的节点
// pre 为改变的新节点的存储目标

func ReverseList(pHead *ListNode) *ListNode {

	var pre, nex *ListNode
	cur := pHead

	for cur != nil {
		nex = cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}

func PrintList(Head *ListNode) {
	// 注意没有表头节点的链表的便利方式
	for p := Head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}

func ReverseList_Not_Nice(pHead *ListNode) *ListNode {

	var Vnum []int

	for p := pHead; p != nil; p = p.Next {
		Vnum = append(Vnum, p.Val)
	}

	fmt.Println(Vnum)

	// 此处注意 i 的取值范围，应为长度减一
	// 注意 for 中多个变量赋值的问题
	// 为什么 for 语句中不能有逗号存在？
	for p, i := pHead, len(Vnum)-1; p != nil; p = p.Next {
		p.Val = Vnum[i]
		i--
	}

	return pHead
}
