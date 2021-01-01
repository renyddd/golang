package main

import "testing"

func TestNewGraph(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(1,3)
	t.Log(g)
}