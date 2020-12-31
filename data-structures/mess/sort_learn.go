package main

import "fmt"

func InsertionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
				fmt.Printf("%d times: ", i)
				fmt.Println(a)
			} else {
				break
			}
		}
		a[j+1] = value
		fmt.Printf("%d times: ", i)
		fmt.Println(a)
	}
}

func main() {
	a := []int{4, 5, 6, 3, 2, 1}
	InsertionSort(a, 6)
}
