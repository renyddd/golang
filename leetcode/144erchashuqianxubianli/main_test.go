package main

import "testing"

func TestAppendNil(t *testing.T)  {
	data := make([]*int, 0)
	t.Log(data)
	data = append(data, nil)
	t.Log(len(data), data)
}