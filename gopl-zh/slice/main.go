package main

import "fmt"

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func AddOneToEachElementAndDoubleIt(slice []byte) {
	for i := range slice {
		slice[i]++
	}
	slice = append(slice, slice...)
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	*slicePtr = (*slicePtr)[:len(*slicePtr)-1]
}

type ptrByteSlice []byte

func (p *ptrByteSlice) SubtractOneFromLength() {
	*p = (*p)[:len(*p)-1]
}

func PtrSliceDouble(slicePtr *[]byte) {
	*slicePtr = append((*slicePtr), (*slicePtr)...)
}

func (p *ptrByteSlice) SubtractHalf() {
	*p = (*p)[:len(*p)/2]
}
func main() {
	s1 := []byte{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s1, len(s1), cap(s1))
	PtrSliceDouble(&s1)
	fmt.Println(s1, len(s1), cap(s1))
	s2 := ptrByteSlice(s1)
	s2.SubtractHalf()
	fmt.Println(s2, len(s2), cap(s2))
}

func mainT2() {
	s1 := []byte{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s1, len(s1))
	PtrSubtractOneFromLength(&s1)
	fmt.Println(s1, len(s1))
	s2 := ptrByteSlice(s1)
	s2.SubtractOneFromLength()
	fmt.Println(s2, len(s2))
}

func mainT1() {
	s1 := []byte{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s1)
	AddOneToEachElement(s1)
	fmt.Println(s1, len(s1))
	AddOneToEachElementAndDoubleIt(s1)
	fmt.Println(s1, len(s1))
}
