## Make your golang src directory.
Install the golang form [golang.google.cn/dl/](golang.google.cn/dl/).

Make the **work** directory in your $HOME directory, like the following.
```
work
├── bin
│   └── hello
├── pkg
└── src
    └── github.com
        └── renyddd
            ├── hello
            │   └── hello.go
            └── README.md
```
Then edit your /etc/profile setting up the environment variables, using **sudo**.
```bash
export GOPATH=$HOME/work
export PATH=$PATH:$GOPATH/bin
```

## The general structure of golang
```go
package main

import (
    "fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
}

func main() {
    var a int
    Func1()
    // ...
    fmt.Println(a)
}

fmt (t T) Method1() {
    // ...
}

func Func1() { // exported function Func1
    // ...
}
```

## Declaration syntax
```go
// var identifier type
var a, b *int
var bo bool
var str string
```
系统为该变量自动赋予其类型的零值，命名遵循驼峰命名法;当你的全局变量需要被外部包所使用时，则将首字母大写（Printf 可以在 fmt 包外被使用）
```go
var a1 = 15
var b1 bool = false
```
> 但是 Go 编译器的智商已经高到可以根据变量的值来自动推断其类型，这有点像 Ruby 和 Python 这类动态语言，只不过它们是在运行时进行推断，而 Go 是在编译时就已经完成推断过程。

此种声明时赋值，编译器亦可自动推断其类型。函数体内声名局部变量时应使用简短格式：
```go
a2 := 666
a3, b3, c3 := 666, 999, "hello" 
// 只能被用在函数体内，而不可以用于全局变量的声明与赋值。
```
交换连个变量可写为：
```go
a, b = b, a
```
空白标识符```_```可用于抛弃指，其为一个只写变量你无法得到它的值
```go
_, b = 5, 7
```
注意字符串与基本数据类型，这里没有写道。

指针: 可用来传递变量的引用，但不允许进行指针计算
```go
var i = 5
var p *int
p = &i
fmt.Printf("An integer: %d, its location is: %p\n", i, &i)
```

## 控制结构

```go
if condition {
    // do somethig
}
```
if 即使当代码块之间只有一条语句时，大括号也不可被省略，{ 左大括号必须和关键字在一行。 
带初始化语句的情况：
```go
if initialization; condition {
    // do something
}
```
switch var1 可以是任何类型，val1 与 val 则可以是同类的任意值，你可以同时测试 case val1，val2， val3 测试多个值。
在分支执行完代码后则推出，可以用 fallthrough 去继续执行后续分支。
```go
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```
循环里只有 for 可用，基本用法：
```go
fro i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
同时使用多个计数器，则会得益于 go 的平行赋值特性：
```go
for i, j := 0, N; i < j; i, j = i+1, j-1 {
    // ...
}
```
for-range 结构，用于迭代一个集合，并同时返回其元素值于下表索引。
```go
str := "hello, world"
for pos, char := range str {
    fmt.Printf("%c is in the %d position\n")
}
```

## 函数
> Go 函数的功能非常强大，以至于被认为拥有函数式编程语言的多种特性。
go 里面有三种类型的函数：
- 普通的带有名字的函数
- 匿名函数或者 lambda 函数
- 方法 Method
除 main()、init() 外其他类型都可以有参数于返回值。函数的参数于返回值及他们的类型统称为函数签名。
函数可将其他函数作为其参数，只要被调函数返回值个数类型于顺序都于调用函数实参一致。
不支持函数重载。
go 使用按值传递来传递参数，也就是传递参数的副本。如需直接修改参数值需要传递其地址。
```go
package main

import "fmt"

func main() {
    fmt.Println(Multiply3Nums(1, 2, 3))
}

func Multiply3Nums(a, b, c int) int {
    return a * b * c
}

func return2Nums(a int) (int, int) {
    return a + 1, a - 1
}
```
### 匿名函数
当我们不希望给函数起名字的时候，可以使用匿名函数。这样的函数不能独立存在，但可以赋值给某个变量，或者直接进行调用。
```go
package main

import "fmt"

func main() {

	func() {
		sum := 0
		for i := 1; i < 1e6; i++ {
			sum += i
		}
		fmt.Println(sum)
	}() // this is calling.

	f()

	func(s string) {
		fmt.Println(s)
	}("hello, world")
}

func f() {
	for i := 0; i < 5; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
```
匿名函数同样被成为闭包（函数式语言的术语），[配合 defer 使用。](https://github.com/renyddd/the-way-to-go_ZH_CN/blob/master/eBook/06.8.md)

一个返回值为另一个函数的函数可以被成为**工厂函数**，数以一个工厂函数而不是为每种情况都书写一个函数：
```go
package main

import (
	"fmt"
	"strings"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpg := MakeAddSuffix(".jpg")

	fmt.Println(addBmp("lena"))
	fmt.Println(addJpg("lena"))
}
```

## 数组
声明格式：
```go
var identifier [len]type
```
(其他部分)[https://github.com/renyddd/the-way-to-go_ZH_CN/blob/master/eBook/07.1.md]
```go
package main

import "fmt"

func f(a [3]int) {
	a[2] = 10
	fmt.Println(a)
}

func fp(a *[3]int) {
	fmt.Println(a)
	a[0] = 99
}

func main() {
	var ar = [3]int{1, 2, 3}
	f(ar)
	fp(&ar)
	f(ar)
}
```
output: 
```go
[1 2 10]
&[1 2 3]
[99 2 10]
```

go 语言中的数组是一种**值类型**，而非 C 中指向首元素的指针。所以可如下方式创建:
```go
var arr1 = new([5]int)
arr2 := *arr1
```
赋值方式：
```go
var a1 = [5]int{18, 20, 15, 22, 16}
var a2 = [...]int{5, 6, 7, 8, 22}
// 若忽略 ... 则将变为切片
var a3 = [5]string{3:"Chris", 4:"Ron"}
// 只有 3 和 4 被赋值，其余为空
```
### 将数组传递给函数
传递数组指针、使用数组的切片可避免内存消耗，不过 go 通常使用切片。

## 切片 Slice

切片是对数组一个连续片段的引用，该数组我们称为相关数组，通常是匿名的，所以是一个引用类型。切片提供一个相关数组的动态窗口。

切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度，切片是一个长度可变的数组。

切片的 ```cap()``` 函数可计算其容量即切片最长可达多少。```len()``` 获取其长度。

若多个切片表示同一个数组的片段，则他们共享数据；
```go
var identifier []type
```
切片的声明不需说明长度，在为初始化前默认为 nil，长度为 0。
```go
var slice1 []type = arr1[start:end]
```
若为 arr1[:] 则取其全部元素。如下两种方式生成由 1，2，3 组成的切片：
```go
s := [3]int{1, 2, 3}[:]
// or
s := []int{1, 2, 3}
```
将切片扩容至他的大小上限，s = s[:cap(s)]

### 切片结构

切片在内存中的组织方式，实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度及切片容量。

### 将切片传递给函数

如果你的一个函数需要对数组进行操作，那你可能总是需要将其参数声明为切片。
```go
func newElement(a []int) {    
    a[2] = 100    
}    
    
func main() {    
    a := []int{3, 2, 5}    
    fmt.Println(a)    
    newElement(a)    
    fmt.Println(a)    
}  
// output
// [3 2 5]
// [3 2 100]
```
### make 创建切片

还未定义相关数组时，通过 make() 函数来创建一个切片，同时创建好相关数组。
```go
slice1 := make([]type, len)
```

### make() 与 new() 的区别
**待补充**

### for-range 循环
第一个返回数组或切片的索引，第二个返回该索引位置的值，均为局部变量。
```go
for index, value := range slice1 {
    ...
}
```
切片的追加与复制：
```go
slice1 := make([]int, 5)
fmt.Println(slice1)
fmt.Printf("%T\n", slice1)

slice2 := append(slice1, 4, 3, 3, 3, 3)
fmt.Println(slice2)

s3 := make([]int, 20)
n := copy(s3, slice2)
fmt.Println(n)
fmt.Println(s3)

// output
[0 0 0 0 0]
[]int
[0 0 0 0 0 4 3 3 3 3]
10
[0 0 0 0 0 4 3 3 3 3 0 0 0 0 0 0 0 0 0 0]
```

## Map

map 是一种元素对的无序集合，也称为关联数组或字典；声明如下
```go
var map1 map[keytype]valuetype
```
map 为引用类型，可用 make 分配
```go
var map1 = make(map[keytype]valuetype)
```
使用实例：
```go
func main() {
    m1 := make(map[int]int)
    for i := 1; i < 3; i++ {
        m1[i] = i * 100
    }

    m2 := map[int]string{1: "one", 2: "tow"}

    fmt.Println(m1)
    fmt.Println(m2)

    v := m1[1]
    fmt.Println(v)

    if _, isPresent := m2[2]; isPresent {
        fmt.Println(m2[2])
    }
}
```

## 结构 struct 与方法 method

go 通过类型别名 alias types 和结构体的形式支持用户自定义类型，或者叫做自定义类型。一个带属性的结构体试图表示一个现实世界中的实体。结构体是复合类型 composite types，它把数据聚集在一起。结构体也是值类型，可通过 new 函数来创建。

组成结构体类型的哪些数据称为字段 field。每个字段都有一个类型名和一个名字。go 中没有类的概念，因此结构体更为重要。

```go
type identifier struct {
    field1 type1
    field2 type2
    ...
}
```
声明结构体 T 类型变量，
```go
var s T
// or
var t *T = new(T)
```
以上两种方式中，t 与 s 都被称做类型 T 的一个实例 instance 或对象 object。

为结构体字段赋值，获取结构体字段的值，都是用点号。在 go 中叫选择器 selector。无论变量为一个结构体类型还是结构体类型指针，都是用同样的选择器符来引用结构体字段。

或者直接初始化结构体变量：
```go
type struct1 struct {
    i1  int
    f1  float32
    str string
}

func main() {
    var ms struct1
    ms.i1 = 1
    ms.f1 = 3.14
    ms.str = "hello, world"
    fmt.Println(ms)
    fmt.Printf("%T\n", ms)

    m1 := new(struct1)
    m1.i1 = 1
    m1.f1 = 3.14
    m1.str = "hello, world"
    fmt.Println(m1)
    fmt.Printf("%T\n", m1)

    m2 := struct1{66, 66.666, "Hi!"}
    fmt.Println(m2)
    fmt.Printf("%T\n", m2)

    m3 := &struct1{66, 66.666, "Hi!"} // 底层仍会调用 new()
    fmt.Println(m3)
    fmt.Printf("%T\n", m3)
}
```
### 内存布局
结构体和它所包含的数据在内存中是以连续块形式存在的，即使嵌套有其他结构体。

### 结构体工厂
go 不支持面向对象中的构造函数，但可以为类型定义一个工厂，按惯例，工厂的名字以 new 或 New 开头。假设有 File 结构体类型：
```go
type File struct {
    fd int
    name string
}
```
下为该结构体对应的工厂方法，他返回一个指向结构体实例的指针：
```go
func NewFile(fd int, name, string) *File {
    if fd < 0 {
    return nil
    }
    return &File{fd, name}
}
```
然后这样调用它:
```go
f := NewFile(10, "./test.txt")
```
你可以说是工厂实例化了类型的一个对象。

**强制使用工厂** 可应用可见性规则，来禁止使用 new() 函数：
```go
package main
import "matrix"

wrong := new(matrix.matrix) // 编译失败，matrix 为私有
right := matrix.NewMatrix(...) // 实例化的唯一方式
```

### 匿名字段
结构体可以包含一个或多个匿名字段（或内嵌字段），即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身也可以是一个结构体类型。在一个结构体中每种数据类型只能有一个匿名字段。

可以粗略的将这个和面向对象语言中的继承概念作比较，在 go 中，组合更受青睐。
### 方法 method
go 的方法是作用在接收者 receiver 上的一个函数，接收者是某种类型的变量。接收者可以是几乎任何类型，但不能是一个接口类型，因为接口是一个抽象定义，但是方法却是具体实现。

一个类型加上它的方法等价于面向对象中的一个类。区别为 go 中类型的代码和绑定在它上面的方法的代码可以不放置在一起，要求是在同一个包中。

因为方法是函数，所以同样不允许重载，即对一个类型只能有一个给定名称的方法。但具有同样名字的方法可以在多个接收者类型上存在。

定义方法的格式如下：
```go
func (recv receiverType) methodName(parameterList) (returnValueLisrt) { ... }
```
如果 recv 是 receiver 的实例，Method1 是它的方法名，那么调用遵循选择器符号 recv.Method1()；
如果 recv 是一个指针，go 会自动解引用；
如果方法不需要 recv 的值，则可通过 ```_``` 舍弃它。

当一个匿名类型被内嵌在结构体时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型**继承**了这些方法：将父类型放在子类型中来实现亚型。

定制类型的字符串形式输出，一种可阅读性和打印性的输出，为类型定义 String() 方法，它会被用在 fmt.Printf() 中生成默认输出。


## 接口
go 中灵活的接口可以实现很多面向对象的特性。接口提供了一种说明对象行为的方式：如果谁能搞定这件事，它就可以用在这儿。

接口定义了一组方法 method set，但是这些方法不包含实现代码，他们是抽象的；接口里也不能包含变量。定义格式如下：
```go
type Namer interface {
    Method1(param_list) return_value
    Method2(param_list) return_value
}
```
按照约定，只包含一个方法的接口的名字由 [e]r 后缀组成；或者以 able 结尾。go 中的接口通常有 0 个最多 3 个方法。

go 中接口可以有值，一个接口类型变量或一个接口值，是一个多字 multiword 数据结构。
```go
var ai Namer
```
ai 它本质上是一个指针，虽然不完全是一回事。指向接口值的指针是非法的，会导致代码错误。

类型（比如结构体）实现接口方法集中的方法，每一个方法的实现说明了此方法是如何作用于该类型的：即实现接口，同时方法集也构成了该类型的接口。
```c
// 接口类型变量结构说明
ai : [ recevier | method table ptr ]
```
实现了 Namer 接口类型的变量可以赋值给 ai（接收者值），此时方法表中的指针会指向被实现的接口方法。如果另一个类型（也实现了该接口）的变量被赋值给 ai，这二者也会随之改变。

几个概念：类型不需要显式声明它实现了某个接口，即接口被显式地实现；多个类型可以实现同一个接口。
实现某个接口的类型（除了实现接口方法外）可以有其他的办法。
一个类型可以实现多个接口。
接口类型可以包含一个实例的引用，该实例的类型实现了此接口（接口是动态类型）。
即使接口在类型之后才定义，二者处于不同的包中，被单独编译，只要实现了接口中的方法，它就实现了此接口。

**多态**：根据当前的类型选择正确的方法，或者说：同一类型在不同的实例上似乎表现出不同的行为。go 版本的多态，接口变量包含一个指向某类型变量的引用，通过它可以调用该类型上的方法（接口中有），这使得此方法更具有一般性。

```go
package main

import "fmt"

type Shaper interface {
    Area() float32
}

type Square struct {
    side float32
}

func (s *Square) Area() float32 {
    return s.side * s.side
}

type Rectangle struct {
    length, width float32
}

func (r *Rectangle) Area() float32 {
    return r.width * r.length
}

func main() {
    r := Rectangle{5, 3}
    q := &Square{5}

    shapes := []Shaper{&r, q}
    for n, _ := range shapes {
        fmt.Println(shapes[n].Area())
    }
}
```
在调用 shapes[n].Area() 时，只知道 shapes[n] 是一个 Shaper 对象，最后摇身一变成为一个 Square 或 Rectangle 对象，还表现出了相应的行为。

**一应仔细看的例子： **

```go
package main

import "fmt"

type stockPosition struct {
    ticker     string
    sharePrice float32
    count      float32
}

func (s stockPosition) getValue() float32 {
    return s.sharePrice * s.count
}

type car struct {
    cmake string
    model string
    price float32
}

func (c car) getValue() float32 {
    return c.price
}

type valuable interface {
    getValue() float32
}

func showValue(asset valuable) {
    fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
    var o valuable = stockPosition{"GOOG", 3.14, 4}
    showValue(o)
    o = car{"BMW", "M3", 66500}
    showValue(o)
}

```
两个类型 sotckPosition 和 car 都有一个 getValue() 方法，接着我们还有一个 valuable 的接口。再定义一个使用 valuable 类型作为参数的函数 showValue() 函数，并在其接口变量完成了调用。

类型断言 type assertion 用于检测一接口类型变量的动态类型，即运行时在变量中存储的值的实际类型。测试某时刻 varI 是否包含类型 T 的值：
```go
if v, ok := varI.(T); ok {
    Process(v)
    return
}
```
之后来理解，断言时，某类型方法的绑定的为指针时， areaIntf.(*Square)在断言类型前需要加星号的问题。 impossible type assertion: Square does not implement Shaper (Area method has pointer receiver)。

[更多详情参照：](https://github.com/renyddd/the-way-to-go_ZH_CN/blob/master/eBook/11.6.md)

使用 sort 标准库排序：[中文源码文档](http://docscn.studygolang.com/src/sort/sort.go?s=7327:7346#L265)

以自定义类型 Person 的 Age 为排序为例，
```go
package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

type Persons []Person

func (p Persons) Len() int           { return len(p) }
func (p Persons) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p Persons) Swap(i, j int)      { p[i].Age, p[j].Age = p[j].Age, p[i].Age }

func main() {
    a := Persons{{"Alice", 18}, {"Bob", 15}, {"Eve", 5}}
    fmt.Println(a)
    sort.Sort(a)
    fmt.Println(a)
}
```
如果此时试着将 Persons 类型的定义改为 []Person 类型定义，则会有如下报错：
```go
// a := []Person{{"Alice", 18}, {"Bob", 15}, {"Eve", 5}}
cannot use a (type []Person) as type sort.Interface in argument to sort.Sort:
    []Person does not implement sort.Interface (missing Len method)
```
同样的道理，将在下例使用标准库中的 IntSlice 时重现。先给出你的程序代码：
```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    a := sort.IntSlice{6, 3, 4, 1, 9, 0, 2, 5, 7, 8}
    fmt.Println(a)

    sort.Sort(a)
    fmt.Println(a)
}
```
此处可能产生的问题还是声明待排序数组处，官方源码中对 Len(), Less(), Swap() 的实现都是在 IntSlice 类型之上，而非 []int；如果已用 []int 进行了声明，则需要在调用 sort.Sort() 函数为其接口类型变量传参前进行强制类型转化。

以下是 sort 库中的代码只摘取部分，你可以在上面连接中找到：

```go
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}

func Sort(data Interface) {
    n := data.Len()
    quickSort(data, 0, n, maxDepth(n))
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p IntSlice) Sort() { Sort(p) }
```
