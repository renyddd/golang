package main

import "fmt"

// https://coolshell.cn/articles/21214.html
// 反转控制

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(x int) {
	set.data[x] = true
}

func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}

func (set *IntSet) contains(x int) bool {
	return set.data[x]
}

// Undoable set

type UndoAbleIntSet struct {
	IntSet    // Embedding
	functions []func()
}

func NewUndoAbleIntSet() UndoAbleIntSet {
	return UndoAbleIntSet{NewIntSet(), nil}
}

func (set *UndoAbleIntSet) Add(x int) { // overwrite
	if !set.contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoAbleIntSet) Delete(x int) { // overwrite
	if set.contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoAbleIntSet) Undo() error {
	if len(set.functions) == 0 {
		return fmt.Errorf("no undo functions.")
	}
	index := len(set.functions) - 1
	if f := set.functions[index]; f != nil {
		f()
		set.functions[index] = nil // for garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}

func main() {
	set := NewUndoAbleIntSet()
	set.Add(3)
	set.Add(6)
	set.Undo()
	set.Delete(3)
	set.Undo()
	fmt.Println(set)
}
