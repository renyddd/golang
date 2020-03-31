package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
	// fmt.Printf("%s\n", runtime.Version())
	fmt.Println("你好 Привет мир Hola mundo Olá, mundo.")

	a1 := 666
	var a2 int
	var a3 int = 999
	fmt.Println(a1, a2, a3)
	fmt.Printf("a1 type: %T\n", a1)

	const Pi = 2.1415926
	// const identifier [type] = value

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

}
