package main

import (
	"fmt"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *ByteCounter) Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, format, args...)
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	var name = "dolly"
	c.Printf("hello, %s\n", name)
}