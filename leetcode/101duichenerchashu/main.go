package main

import "github.com/renyddd/golang/leetcode/treenode"

func main() {
	return
}

func isSymmetric(root *treenode.TreeNode) bool {
	return recursionCheck(root.Left, root.Right)
}

func recursionCheck(left, right *treenode.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		// 左右一边不为空时一定不对称
		return false
	}
	return left.Val == right.Val && recursionCheck(left.Left, left.Right) && recursionCheck(right.Left, right.Right)
}

func isSymmetricLoop(root *treenode.TreeNode) bool {
	l, r := root, root
	q := []*treenode.TreeNode{}
	q = append(q, l)
	q = append(q, r)

	for len(q) > 0 {
		l, r = q[0], q[1]
		q = q[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		q = append(q, l.Right)
		q = append(q, r.Left)

		q = append(q, r.Right)
		q = append(q, l.Left)
	}
	return true
}