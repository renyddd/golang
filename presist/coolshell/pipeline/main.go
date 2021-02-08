package main

// https://coolshell.cn/articles/21228.html

import (
	"fmt"
)

func echo(in []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range in {
			out <- v
		}
		close(out)
	}()
	return out
}

func sqr(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if v % 2 == 0 {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}

var nums = []int{1, 2, 3, 4, 5, 6, 7}

func NotGraceful() {
	ch := sum(sqr(echo(nums)))
	for n := range ch {
		fmt.Printf("%v\t ", n)
	}
	fmt.Printf("\n")
}

type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

func PipeLine(nums []int, echo EchoFunc, pipefns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for _, f := range pipefns {
		ch = f(ch)
	}
	return ch
}

func Graceful() {
	ch := PipeLine(nums, echo, sqr, odd, sum)
	for n := range ch {
		fmt.Printf("%v\t", n)
	}
}

func main() {
	Graceful()
}