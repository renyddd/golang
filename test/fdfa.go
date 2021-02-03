package test

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, p, q)
	var t1, t2 *TreeNode

	for len(queue) > 0 {
		t1, t2 = queue[0], queue[1]
		queue = queue[2:]

		if t1.Left == nil && t1.Right == nil && t2.Left == nil && t2.Right == nil {
			continue
		}


		queue = append(queue, t1.Left, t2.Left, t1.Right, t2.Right)
	}

	return true
}