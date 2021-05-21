package main

func doSliceReverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func doSliceReverse1(a []int) {
	i, j := 0, len(a)-1

	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func doSliceReverse2(a []int) []int {
	if len(a) == 0 {
		return a
	}
	return append(doSliceReverse2(a[1:]), a[0])
}
