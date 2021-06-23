package main

/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，
只保留原始链表中没有重复出现的数字。
返回同样按升序排列的结果链表。

示例 1：
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

示例 2：
输入：head = [1,1,1,2,3]
输出：[2,3]

提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	tail := &ListNode{Val: -1}
	dummyHead := tail
	nmap := make(map[int]int, 0)

	for p := head; p != nil; p = p.Next {
		if _, ok := nmap[p.Val]; ok {
			nmap[p.Val]++
		} else {
			nmap[p.Val] = 0
		}
	}

	// 注意 map 读取时的无序性
	sortedNums := make([]int, 0)
	for k, v := range nmap {
		if v == 0 {
			sortedNums = append(sortedNums, k)
		}
	}

	for p := 1; p < len(sortedNums); p++ {
		for i := p; i > 0; i-- {
			if sortedNums[i] < sortedNums[i-1] {
				sortedNums[i], sortedNums[i-1] = sortedNums[i-1], sortedNums[i]
			}
		}
	}

	for _, v := range sortedNums {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}

	return dummyHead.Next
}

// deleteDuplicates2 采用双指针的方式进行遍历，当第一个指针判断出 Next 为重复元素时，
// 则用第二个指针来找到第一个不重复的节点
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{Val: -1, Next: head}
	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val != x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
