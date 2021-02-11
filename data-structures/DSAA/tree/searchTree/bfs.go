package main

// 这里只是尝试实现
// BFS 用队列来维护遍历顺序
func BFS(root *TreeNode) []ElementType {
	var curNode *TreeNode
	queue := make([]*TreeNode, 0)
	res := make([]ElementType, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		curNode = queue[0]
		queue = queue[1:]

		if curNode != nil {
			res = append(res, curNode.Element)
			queue = append(queue, curNode.Left, curNode.Right)
		}
	}

	return res
}
