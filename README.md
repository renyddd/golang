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
