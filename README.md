# 项目学习
- context
- net/http
- https://github.com/spf13/tri/tree/master/todo
- https://pkg.go.dev/k8s.io/client-go

# 新知识点

理解`GOPATH`的重要性，亦即 Go 语言源码的组织方式。

以`go build`命令为例，搞清楚开始匹配的其实相对路径。

### import 包管理

环境如下（首先确保如下环境变量设置正确）：

```bash
➜ go env -w GO111MODULE=auto
➜ pwd
/home/renyddd/work/src/github.com/renyddd
➜ mkdir demox
➜ cd demox
```

编辑文件`demox.go`如下：

```go
package main

import "fmt"

func main() {
        fmt.Println("hello come from package main!")
}
```

表明该代码属于 main 包，运行后即可得到相应输出。

>  [working directory is not part of a module 问题解决](https://stackoverflow.com/questions/61921282/golang-cannot-find-module-providing-package-package-name-working-directory-is)，这些报错无法解决的原因还是因为你对其基础的不了解。

再进行代码的拆分，将输出功能移至 libx 库中：

```bash
➜ mkdir lib
➜ cd lib  
```

编辑 demox_libx.go 文件如下：

```go
package libx

import "fmt"

func Hello() {
        fmt.Println("come from package libx")
}
```

再通过相对路径对该包进行安装，便会生成如下文件结构：

```bash
➜ go install github.com/renyddd/demox/libx

➜ tree /home/renyddd/work
├── pkg
│   ├── linux_amd64
│   │   └── github.com
│   │       └── renyddd
│   │           └── demox
│   │               └── libx.a
```

再进行修改 demox.go 文件为：

```go
package main

import (
        "github.com/renyddd/demox/lib"
)

func main() {
        libx.Hello()
}
```





# 旧文

Go 或者你可以称其为 Golang，是由谷歌团队以及开源社区的贡献者们开发的开源编程语言。2007 年 9 月 Go 的设计者之中就包括肯·汤普逊，并于两年后宣布推出。

[https://golang.org/](https://golang.org/) 是 Go 的官网；

[https://golang.google.cn/](http://docscn.studygolang.com/) 如官网不存在你也可以选择访问这里；

[https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.1.md](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.1.md) 一本很版很棒的开源教程

接着在你的家目录下为 Go 创建一个工作目录，结构如下：

```go
work
├── bin
│   └── hello
├── pkg
└── src
    └── github.com
        └── USERNAME
            └── hello
                └── hello.go

```

该 work 工作目录中的 src 存放着你工程中的源文件，bin 目录中存放着编译安装之后的可执行文件，pkg 中则为包文件。

> 你可以把 GOPATH 简单理解成 Go 语言的工作目录，它的值是一个目录的路径，也可以是多个目录路径，每个目录都代表 Go 语言的一个工作区（workspace）。

有关 go 程序的子命令，你都可以在命令行下获得帮助：

```bash
go help <topic>
```

目前你最常用到的就属 go run 了，它会为你编译并运行你所指定的源码文件；go 有标准编码风格，也为你提供了 go fmt 来格式化你的程序文件。
## 程序结构

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```

经典的第一步，你的任一可执行 go 程序都必须在第一行指明其属于 main 包；第二行导入实现了格式化输入输出的 fmt 包；主函数必须无参无返回值。

### 变量
你可以通过 var 关键字后跟标识符及其类型，完成变量的声明，

```go
var a int
var p1, p2 *int
var s string
```

或者你也可以将变量的声明于初始化结合起来，

```go
var a int = 9
var s string = "Hello, world"
```

或者你也可以让 go 编译器自行推断变量类型，

```go
var a = 9
var p = &a
```

或者你还可以使用简短格式，

```go
i := 1
s := "Hi"
```

由于 go 语言同时赋值特性的存在，你可以一次同时初始化多个变量，

```go
a, b, c := 1, 2.0, "three"
```

当你需要交换两变量时，已不需要再借助第三变量，

```go
a, b = b, a
```

还要说到的是 go 语言不准你声明了变量却不使用，这会导致你在编译时错误。

### 分支
go 语言完全去掉了 if、for 语句条件判断两侧的圆括号，并严格规定了花括号的使用格式，

```go
if 1 != 2 {
	fmt.Println("True")
} else {
	fmt.Println("Impossible")
}
```

你也可以在 if 语句的条件判断中添加一个初始化语句，

```go
if a := (1 != 2); a {
	fmt.Println(a)
}
```

go 只为你提供了强大的 for 循环，

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

你可以按需求省去分号中任一部分的内容，或者全部省去变为一个死循环，

```go
for {
	fmt.Println("Hi")
}
```

还有 for range 结构，可以帮你快速迭代一个集合，并返回每次的索引以及其值，

```go
str := "Hello, world"
for i, v := range str {
	fmt.Println(i, v)
}
```

如果后续你不需要用到其中的索引变量，但又要避免编译时的变量未使用报错，即可使用下划线以丢掉，

```go
for _, v := range str {   // ...
```

### 函数
go 函数定义由 func 关键字开始，并按需给出形参类型个数，与返回值类型及个数，

```go
func sum2(a, b int) int {
	return a + b
}
```

是的，go 函数可以返回多个值（后面都省掉了包文件部分，）

```go
func main() {
	fmt.Println(twoNums(1))
}

func twoNums(a int) (i, j int) {
	i = a - 1
	j = a + 1
	return i, j
}
```

并且 go 函数还能接收变长参数，只要当其最后一个参数为 ...type 的形式时即可，

```go
func main() {
	hi()
	hi("Alice")
	hi("Alice", "Bob")
}

func hi(s ...string) {
	if len(s) == 0 {
		return
	}
	fmt.Printf("%T \n", s)
	// 输出为：[]string，即为切片类型
	for _, v := range s {
		fmt.Println("Hi,", v)
	}
}
```

在用三个点接受变长参数的情况下，其实相当于是在函数中的一个某类型的切片（slice）变量中存储了参数值；后文会提到切片，因此你才能使用 for-range 进行迭代。

如果你选择空接口（interface）作为变长参数的类型，那么你的函数将可以接受任意长度任意类型的变量，

```go
package main

import "fmt"

type Any interface{}

func main() {
	hiAny("Alice", true, 1, 2.0)
}

func hiAny(a ...Any) {
	for _, v := range a {
		fmt.Println("Hi,", v)
	}
	fmt.Printf("%T \n", a)
}
```

有关接口与 type 的使用，都放在后文中说，接下来是

### 匿名函数
直接理解匿名函数（anonymous function）就是不包含名字的函数，它除了不需名字外就跟普通函数的定义方式一样了，还是以计算两数和为例

```go
func(a, b int) int {
	return a + b
}
```

那么现在的问题是，就既然它没有名字那而你又需要调用。你有两种选择，第一是当你需要时就定义，定义完就直接用，

```go
res := func(a, b int) int {
	return a + b
}(1, 2)
```

直接在函数体后跟上小括号就表示调用，并传好要计算的数值同时还要接收好计算结果。

你的第二选择，是将该匿名函数保存至一变量当中；

```go
f := func(a, b int) int {
	return a + b
}
f(1, 2)
```

好在你不用担心变量 f 的类型，编译器会明白它是函数类型类型的变量 —— 这是有点绕，func(int, int) int 这就是变量 f 的类型。

有意思的来了，既然如此那匿名函数也就能像变量一样作为另一函数的返回值，看着有层次的写法是这样的，

```go
func outer() (func(int, int) int) {
	fmt.Println("something in outer")
	return func(a, b int) int {
		return a + b
	}
}
```

这样写没错，但是 go fmt 会自动为你去掉 outer 函数返回值部分的小括号，后面的代码中会遵循这种写法。其中 outer 函数输出部分亦为函数体，放在这里只是想说明外层函数当然不止仅仅用于返回求和函数这一个作用。下来是主函数，

```go
func main() {
	fsum := outer()
	fmt.Println(fsum(1, 2))
}
```

对应前面，变量 fsum 的类型应为 outer() 函数的返回值类型：func(int, int) int ，接收两整形形参返回一整形的函数类型；接着再调用它并完成一与二的求和运算。

上面的例子好像除了麻烦并没有带来什么别的东西，不过这种返回一个函数的函数也被成为工厂函数，当你需要一系列相似的函数时就会用到。这里以不同的打招呼方式为例，当你遇见 Alice 和 Bob 的时候要说 hi，而遇见世界的时候要说你好，

```go
package main

import "fmt"

func main() {
	sayHiTo := MakeGreetPrefix("Hi, ")
	sayHelloTo := MakeGreetPrefix("Hello, ")

	sayHiTo("Alice")
	sayHiTo("Bob")
	sayHelloTo("world")
}

func MakeGreetPrefix(greeting string) func(string) {
	return func(name string) {
		fmt.Println(greeting + name)
	}
}
```

就像这样你无需将重复的 Print 实现两次，你只需要通过工厂函数来实现每次不同的招呼方式即可。

### 数组与切片
数组就是同一元素的集合，你可以像下面这样声明一包含 5 个整形默认值为零的数组，并输出其类型，

```go
func main() {
	var a [5]int
	fmt.Printf("%T\n", a)
}
// output： [5]int
```

从结果你会发现，在 go 中连数组的长度都算作了数组类型。数组是固定长度你无法为它重新重新定义大小，但是动态长度的切片（slice）在 go 中更为常用。

```go
var a = [5]int{1, 2, 3, 4, 5}
s1 := a[1:3]
s2 := a[:]
```

通过同一个数组的两个不同索引来界定一个切片时，你创建的切片门都会指向那个关联数组；因此传参给 go 函数时使用数组会造成额外的复制浪费，而使用切片则传递了索引。还有内建函数 make，在指明切片长度后会为你创建一全零数组同时返回切片，

```go
s := make([]int, 10)
```

这是取自官网的图片与说明：切片是一数组的描述符，它包含着一个指向关联数组的指针、长度（切片涉及到的数组元素）以及容量（关联数组大小。）
![https://blog.golang.org/slices-intro](https://img-blog.csdnimg.cn/20200610100818982.png =300x)
![https://blog.golang.org/slices-intro](https://img-blog.csdnimg.cn/20200610100735777.png =300x)
## 结构与方法
结构（struct）是字段（field）的集合；字段为组成结构体的数据，并可通过 .（点）访问其字段。

```go
type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{"Alice", 4}
	fmt.Println(p1.name)
	p1.age = 5
	fmt.Println(p1.age)
}
```

go 在通过指针访问结构字段时，你不需要自己进行解引用，

```go
p2 := new(Person)
p2.name = "Bob"
fmt.Println(p2)
```

go 中没有对象的概念，也就自然没有所谓构造函数一说，但你可以通过以 New 开头命名的结构工厂来初始化对象，

```go
func NewPerson(name string, age int) *Person {
	if age < 0 {
		return nil
	}
	return &Person{name, age}
}

func main() {
    p3 := NewPerson("Bob", 7)
}
```

当结构中包含一个或多个没有名字的字段时，这些字段被称为（anonymous field）匿名字段或内嵌字段，其只有类型名是必须的。当一结构体作为匿名字段被内嵌在另一结构体时，便可以与面向对象中的继承相比较。

```go
type Man struct {
	Person
	gender string
}

func main() {
	p4 := Man{Person{"Bob", 7}, "male"}
	fmt.Println(p4)
}
```

## 方法
方法 method 定义于 go 的类型之上，即在定义函数的关键字 func 与函数名之间加上一个接收者变量；意为这个方法是作用在这个接收者类型之上的。

```go
func (m Man) Eat() {
	fmt.Println("I ate an apple.")
}
```

当使用值类型的接收者时，表现会跟普通的函数一样，操作的是原数据的一份拷贝。或者你也可以将接受者类型定义为指针，这样便可直接修改指针所指向的值，

```go
func (m *Man) GrowUp(n int) {
	m.age = m.age + n
}

func main() {
	p4 := Man{Person{"Bob", 7}, "male"}

	p4.Eat()
	p4.GrowUp(1)
    
	fmt.Println(p4)
}
```

### 接口
对接口（interface）类型的定义是一组方法集，之后的类型有实现了同名方法的就可以认为它也实现了这个接口，因此接口类型的变量就可以容下该类型的变量。

现在上面的例子中添加一个 Tiger 结构，与人结构一样，老虎也需要有吃这个方法，

```go
type Tiger struct {
	name string
	age  int
}

func (t Tiger) Eat() {
	fmt.Println("I do not eat apples.")
}
```

再来定义一个名为 Eater 的接口，其中就包含有人与老虎结构共实现了的 Eat() 这个方法，

```go
type Eater interface {
	Eat()
}
```

创建好接口类型变量后，就可以让他容下实现了同名方法的 Tiger 结构了，再当在接口类型变量上调用 Eat() 方法时，就会表现出 Tiger 作为接收者的方法了，

```go
var eatIntf Eater = Tiger{"BagCat", 3}
eatIntf.Eat()
```

当然还可以用到循环之中，这就有点像是所谓的多态了，

```go
t1 := Tiger{"BagCat", 3}
eatIntf := []Eater{p4, t1}

for _, e := range eatIntf {
	e.Eat()
}
```

本文的最后就到了前文提到过的空接口了，接口类型说的是谁实现了我有的方法，我就能容下谁；既然任何类型都至少满足了无方法，那么空接口也就自然能容下任何值了。
