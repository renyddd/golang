nowcoder 牛客网刷体记录
=========

在刚开始训练时，本地提供完整的可运行代码。
最好每章都能有必要的注释，或者加上一些 go 语言语法的拓展部分。



## 读取用户输入

使用`fmt`包中的`Scanln`函数：

```go
var n1, n2 int
fmt.Scanln(&n1, &n2)
```

或者使用`bufio`包缓冲读写，待补充。

近期时间有限，一定要用在紧迫的技能上！



## ListNode
### [合并链表](./ListNode_Merge.go)

使用 `new` 为结构体变量分配内存，分配当前类型所属的零直：

```go
var head = new(ListHead)
head := new(ListHead)
```

或者使用混合字面量语法（composite literal syntax）`&struct1{1,nil}`，其为一种简写但底层还是在调用`new`，此时赋值的字段必须按照顺序来写。

当然这样是等价的`new(ListHead)`, `&ListHead{}`。

```go
head := &ListHead{Val: 5, Next: nil}
```

为类型定义结构体工厂，工厂名以 new 或者 New 开头，一次来简单的实现构造函数：

```go
func NewListNode(val int, next *ListNode) *ListNode {
        if val < 0 { 			// 强迫 Val 直大于零
                return nil
        }
        return &ListNode{Val: val, Next: next}
}
```

同样运用可见性规则（当标识符以大写字母开头则其对象可被外部包代码所使用），可以强制使用工厂函数。

