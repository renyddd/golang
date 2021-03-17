// main 这里借助 go 官方的实现：https://golang.org/src/container/heap/heap.go
// 需要关注完全二叉树以数组存储时的性质
package heap

import (
	"sort"
)

// 堆是一颗完全二叉树，其节点都是其子树上的最小节点
// 堆也是最常用的一种实现优先队列的方式

// Interface 注意如下接口中的 pop，push 方法（绑定了结构体的）是为了该包中堆的实现
// 也是从自定义的元素中，真正向末尾添加 or 删除首元素的方法
// 要想从堆中添加或移除元素请使用（包名）heap.Pop or heap.Push
type Interface interface {
	sort.Interface
	// 实现向尾添加
	Push(x interface{})
	// 实现从尾删除
	Pop() interface{}
}

func Init(h Interface) {
	// 堆化
	n := h.Len() // 为 sort 中接口而实现
	// 从非叶子层开始
	for i := n/2 - 1; i >= 0; i-- {
		// 下滤
		down(h, i, n)
	}
}

// down 这里的下滤描述的是从 i0 节点开始的不断向下 percolate
func down(h Interface, i0, n int) bool {
	i := i0
	for {
		// 从左孩子开始比较
		j1 := 2*i + 1
		// 注意此处的等于号，否则会因此数组越界 panic
		if j1 >= n || j1 < 0 { // 当整型溢出时 j1 < 0
			break
		}
		j := j1 // 左孩子
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			// 当其右孩子比左孩子更小时
			j = j2 // 右孩子
		}
		// 以上操作目的是从当前节点的左右孩子中知道一个最小值，给 j

		// 此时 j 已代表为待处理的左节点或右节点
		if !h.Less(j, i) {
			// 注意这是小顶堆，需要将较小的叶子节点交换上来
			break
		}
		// 用左右孩子中的最小值来和跟节点交换
		h.Swap(i, j)
		// 这里的赋值终会导致循环判断条件不符合
		// j 由 j1,j2 而选择出来的左节点 or 右节点，赋值给 i
		// 进行 for 内的下一次 percolate
		// 因为以 Init 中的调用为例，i0 参数由数组中值开始递减为 0
		// 若在初始化的切片中堆顶元素为最大值，且堆顶元素的操作只会在 i0 为 0 时涉及到
		// 交换完最大的堆顶元素后，则必然会导致堆的性质被破坏，因此需要继续下滤
		i = j
	}
	return i > i0
}

// Push 将元素 x 加入到堆中，时间复杂度 O(log n)
// 始终将待插入元素追加至尾部，即二叉堆的最后一个叶子节点，以保证完全二叉树的性质不会被破坏
func Push(h Interface, x interface{}) {
	h.Push(x)
	// 这里默认你已经成功插入了元素在末尾（append）
	// 并且长度信息也已经更新
	// 传递给上滤的参数就是末尾元素，也就是 x
	up(h, h.Len()-1)
}

// up 将末尾元素不断与其父节点进行比较，若更小则交换
func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

// Pop 移除并且返回最小值元素
// 时间复杂度是 O(log n)
// 注意，因保证堆的结构性质 故在删除首元素时需要先交换首位元素
// 因此 Interface 中的 Pop 方法交给用户实现的也是删除尾元素
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	// 首先将堆顶元素与末尾元素进行交换
	// 再从堆顶开始下滤，以保持（最小堆的）堆序性质
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// Remove

// Fix
