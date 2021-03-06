package listnode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// PrintList .
func (l *ListNode) PrintList() {
	p := l
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}

func (l *ListNode) addToTail(n int) {
	newL := &ListNode{
		Val:  n,
		Next: nil,
	}
	l.Next = newL
}

// MakeListNode usage:
// head := listnode.MakeListNode([]int{1, 2, 3, 4, 5, 6})
// h := (*ListNode)(unsafe.Pointer(head))
func MakeListNode(nums []int) *ListNode {
	head := new(ListNode)
	p := head
	for i := 0; i < len(nums); i++ {
		p.addToTail(nums[i])
		p = p.Next
	}
	return head.Next
}
