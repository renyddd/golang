package main

import "fmt"

// 对接口的理解：为即将满足该接口所包含的方法的结构体而准备的一些共性的方法
// 注意：在那些共性的方法中的参数即为该接口类型；可参见 sort 包

type Interface interface {
	Len() int
	Swap(i, j int)
	DeleteLast() interface{}
}

func DeleteTheSecondToLast(i Interface) interface{} {
	if i.Len() < 2 {
		return i
	}
	last := i.Len() - 1
	i.Swap(last, last-1)
	return i.DeleteLast()
}

type People struct {
	Name string
	Age  int
}

type Peoples []People

func (ps Peoples) Len() int { return len(ps) }
func (ps Peoples) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func (ps Peoples) DeleteLast() interface{} {
	return ps[:ps.Len()-1]
}

type Ints []int

func (i Ints) Len() int      { return len(i) }
func (i Ints) Swap(k, j int) { i[k], i[j] = i[j], i[k] }

// 参见如下的实现错误的函数，这就将改变最终的正确性
// func (i Ints) Swap(k, j int) { i[k], i[j] = i[k], i[j] }

func (i Ints) DeleteLast() interface{} {
	return i[:i.Len()-1]
}

func main() {
	ps := Peoples{
		{"Bob", 3},
		{"Tom", 5},
		{Name: "Zik", Age: 7},
	}
	fmt.Println(ps, ps.Len())

	res := DeleteTheSecondToLast(ps).(Peoples)
	fmt.Println(res)

	is := Ints{1, 3, 5, 7, 9}
	fmt.Println(is)
	fmt.Println(DeleteTheSecondToLast(is))
}
