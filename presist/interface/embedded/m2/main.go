package main

import "fmt"

type withTypeName struct {
	Type, Name string
}

type Printable interface {
	PrintStr()
}

func (w withTypeName) PrintStr() {
	fmt.Println(w.Type, w.Name)
}

type City struct {
	withTypeName
}

func NewCity(name string) City {
	return City{withTypeName{"City", name}}
}

type Country struct {
	withTypeName
}

func NewCountry(name string) Country {
	return Country{withTypeName{"Country", name}}
}

func main() {
	ci := NewCity("city1111")
	co := NewCountry("coutry22222")

	cList := []Printable{ci, co}
	for _, v := range cList {
		v.PrintStr()
	}
}
