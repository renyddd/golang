package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int8
	var a1 int32
	var b int
	var ba  int64
	fmt.Println(reflect.TypeOf(a).Size())
	fmt.Println(reflect.TypeOf(a1).Size())
	fmt.Println(reflect.TypeOf(b).Size())
	fmt.Println(reflect.TypeOf(ba).Size())
}