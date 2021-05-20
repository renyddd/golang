package main

import "testing"

func TestClient(t *testing.T) {
	adidasFactory, _ := getSportsFactory("adidas")

	t.Log(adidasFactory.makeShirt())
	t.Log(adidasFactory.makeShoe())

	nikeFactory, _ := getSportsFactory("nike")

	t.Log(nikeFactory.makeShirt())
	t.Log(nikeFactory.makeShoe())
}
