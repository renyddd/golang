package main

// ref: https://leetcode-cn.com/problems/clone-graph/

type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraphO1 进行深度优先
func cloneGraphO1(node *Node) *Node {
	// 原始节点 -> 克隆节点
	visited := make(map[*Node]*Node, 0)
	// 注意局部变量的命名！
	// 注意局部变量的命名！
	// 注意局部变量的命名！
	var doClone func(curNode *Node) *Node
	doClone = func(curNode *Node) *Node {
		if curNode == nil {
			return nil
		}

		if _, ok := visited[curNode]; ok {
			return visited[curNode]
		}

		cloneNode := &Node{curNode.Val, []*Node{}}
		visited[curNode] = cloneNode

		for _, n := range curNode.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, doClone(n))
		}

		return cloneNode
	}

	return doClone(node)
}

// cloneGraphO2 使用队列来进行广度优先遍历
func cloneGraphO2(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node, 0)
	// 首先进行克隆第一个节点
	visited[node] = &Node{node.Val, []*Node{}}
	nodeQueue := []*Node{node}

	for len(nodeQueue) > 0 {
		curNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		for _, v := range curNode.Neighbors {
			if _, ok := visited[v]; !ok {
				visited[v] = &Node{v.Val, []*Node{}}
				nodeQueue = append(nodeQueue, v)
			}
			visited[curNode].Neighbors = append(visited[curNode].Neighbors, visited[v])
		}
	}

	return visited[node]
}