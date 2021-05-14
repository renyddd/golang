package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoes := nikeFactory.makeShoe()
	nikeShirts := nikeFactory.makeShirt()

	adidasShirts := adidasFactory.makeShirt()

	fmt.Println(nikeShoes, adidasShirts, nikeShirts)
}
