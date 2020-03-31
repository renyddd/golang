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
此种声明时赋值，编译器亦可自动推断其类型。函数体内声名局部变量时应使用简短格式：
```go
a2 := 666
a3, b3, c3 := 666, 999, "hello" 
```
交换连个变量可写为：
```go
a, b = b, a
```
空白标识符```_```可用于抛弃指，其为一个只写变量你无法得到它的值
```go
_, b = 5, 7
```

