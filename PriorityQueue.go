package mygo

import (
	"container/heap"
)

// Item 实现优先队列中的一个元素(一项).
type Item struct {
	Value    int // 值
	Priority int // 优先级
	Index    int // 索引
}

// PriorityQueue 实现接口.
type PriorityQueue []*Item

// Len 返回优先队列长度.
func (pq *PriorityQueue) Len() int { return len(*pq) }

// Less 比较两个项优先级.
func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Priority < (*pq)[j].Priority
}

// Swap 交换两个项.
func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].Index = i
	(*pq)[j].Index = j
}

// Push 向优先队列中添加项.
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

// Pop 移除并返回优先级最高的项.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// Update 修改优先级和值的item在优先级队列中的位置.
func (pq *PriorityQueue) Update(item *Item, value int, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
