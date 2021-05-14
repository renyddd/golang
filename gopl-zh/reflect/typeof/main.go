package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func main() {
	x := 1
	fmt.Println(Any(&x))
}

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func T2() {
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Printf("%T\n", v)
	fmt.Println(v.String())

	t := v.Type()
	fmt.Println(t.String())

	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)
}

func T1() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String()) // int
	fmt.Println(t)          // int

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File
	fmt.Printf("%T\n", 3)
}
