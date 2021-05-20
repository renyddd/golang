package main

type Interface interface {
	doStuff()
}

type Product struct{}

func (p *Product) doStuff() {}

type ConcreteProductA struct {
	Product
}

func NewConcreteProductA() *ConcreteProductA { return &ConcreteProductA{} }

type ConcreteProductB struct {
	Product
}

func NewConcreteProductB() *ConcreteProductB { return &ConcreteProductB{} }

func Creator(productType string) Interface {
	switch productType {
	case "a":
		return NewConcreteProductA()
	case "b":
		return NewConcreteProductB()
	default:
		panic("wrong error")
	}
}
