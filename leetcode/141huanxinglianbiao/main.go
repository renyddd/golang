package main

import (
	"../listnode"
	"fmt"
)

func main() {
	head := listnode.MakeListNode([]int{1,2,3,4,5})
	head.PrintList()

	fmt.Println(HasCycle(head))
	BuildRingList(head, 4)

	head.PrintList()
}

func HasCycle(head *listnode.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// BuildRingList 将尾节点连接到指定为之
func BuildRingList(head *listnode.ListNode, pos int) {
	tail := head
	for ; tail.Next != nil; tail = tail.Next {
	}
	fmt.Println(tail, pos)

	var i int = 0
	for p := head; p.Next != nil; p = p.Next {
		i++
		if i == pos {
			//fmt.Println(p)
			tail.Next = p
			break
		}
	}
}