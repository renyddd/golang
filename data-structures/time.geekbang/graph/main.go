package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	// the num of vertex
	n int
	adj []*list.List
}

func NewGraph(n int) *Graph {
	g := &Graph{
		n: n,
		adj: make([]*list.List, n),
	}
	for i, _ := range g.adj {
		g.adj[i] = list.New()
	}
	return g
}

func (g *Graph) AddEdge(s, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}
	prev := make([]int, g.n)
	for index := range prev {
		prev[index] = -1
	}
	var queue []int
	visited := make([]bool, g.n)
	visited[s] = true
	isFound := false
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedlist := g.adj[top]
		for e := linkedlist.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}

	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Println("Not Found")
	}
}

func printPrev(a []int, s, t int) {
	if t == s || a[t] == 1 {
		fmt.Print(t)
	} else {
		printPrev(a, s, a[t])
		fmt.Print(t)
	}
}