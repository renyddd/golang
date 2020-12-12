//package main
//
//import (
//	"fmt"
//	"io"
//)
//
//type Store interface {
//	Open(string) (io.ReadWriteCloser, error)
//}
//
//type StorageType int
//
//const (
//	DiskStorage StorageType = 1 << iota
//	TempStorage
//	MemoryStorage
//)
//
//func NewStore(t StorageType) Store {
//	switch t {
//	case MemoryStorage:
//		return newMem
//	case DiskStorage:
//		return newDisk
//	default:
//		return newTem
//	}
//}
//
//func main() {
//
//	fmt.Println(c1, c2, c3)
//}