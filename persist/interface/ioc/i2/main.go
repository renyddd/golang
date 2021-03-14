package main

import "fmt"

// https://coolshell.cn/articles/21214.html
// 反转依赖

type Undo []func()

func (u *Undo) Add(f func()) {
	*u = append(*u, f)
}

func (u *Undo) Undo() error {
	functions := *u
	if len(functions) == 0 {
		return fmt.Errorf("no undo functions.")
	}
	index := len(functions) - 1
	if f := functions[index]; f != nil {
		f()
		functions[index] = nil // for garbage collection
	}
	*u = functions[:index]
	return nil
}

// 将 undo 嵌入 IntSet

type IntSet struct {
	data map[int]bool
	undo Undo
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool), undo: []func(){}}
}

func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Add(x int) {
	if !set.contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
	delete(set.data, x)
}

func (set *IntSet) contains(x int) bool {
	return set.data[x]
}

func main() {
	set := NewIntSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Undo()
	fmt.Println(set)
	set.Delete(5)
	fmt.Println(set)
	set.Undo()
	fmt.Println(set)

	/*
		{map[1:true 2:true] [0x10a96c0 0x10a96c0]}
		{map[1:true 2:true] [0x10a96c0 0x10a96c0 <nil>]}
		{map[1:true 2:true] [0x10a96c0 0x10a96c0]}
	*/
}
