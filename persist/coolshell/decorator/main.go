package main

import "fmt"

func Hello(s string) {
	fmt.Println("Hello", s)
}

func Decorate(fn func(string)) func(string) {
	return func(is string) {
		fmt.Println("Start")
		fn(is)
		fmt.Println("End")
	}
}

func main() {
	Decorate(Hello)("Tom")
}
