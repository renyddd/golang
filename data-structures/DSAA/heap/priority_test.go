package heap

import (
	"fmt"
	"testing"
)

type Item struct {
	value    string // item 的值
	priority int    // 优先级
	index    int    // 在堆中的索引
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// 我们需要得到优先级最高的而非最小的
	return pq[i].priority > pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[:n-1]
	return item
}

func TestPriorityQueue(t *testing.T) {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}

	Init(&pq)

	item := &Item{
		value:    "orange",
		priority: 1,
	}
	Push(&pq, item)

	for pq.Len() > 0 {
		item := Pop(&pq).(*Item)
		fmt.Println(item.value, item.priority)
	}
}
