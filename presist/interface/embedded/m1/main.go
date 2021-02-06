package main

import "fmt"

type withName struct {
	name string
}

type City struct {
	withName
}

type Country struct {
	withName
}

type Printable interface {
	PrintStr()
}

func (w withName) PrintStr() {
	fmt.Println(w.name)
}

func main() {
	ci := City{withName{"city1"}}
	co := Country{withName{"country1"}}
	cList := []Printable{ci, co}

	for _, v := range cList {
		v.PrintStr()
	}
}

