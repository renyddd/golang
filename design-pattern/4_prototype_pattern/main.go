package main

import "fmt"

// https://refactoringguru.cn/design-patterns/prototype/go/example#example-0

type inode interface {
	print(string)
	clone() inode
}

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation, f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation, f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode

	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	f1 := &file{"file 1"}
	f2 := &file{"file 2"}
	f3 := &file{"file 3"}

	folder1 := &folder{
		children: []inode{f1},
		name:     "folder 1",
	}

	folder2 := &folder{
		children: []inode{f2, f3, folder1},
		name:     "folder 2",
	}

	folder2.print("    ")

	fmt.Println(" --- clone ---")
	folder2.clone().print("    ")
}
