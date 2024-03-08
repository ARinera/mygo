package mygo

import (
    "errors"
    "sync"
)

// ArrayQueue 数组队列
type ArrayQueue struct {
    array []int
    size  int
    lock  sync.Mutex
}

// LinkQueue 链表队列
type LinkQueue struct {
    root *LinkNode
    size int
    lock sync.Mutex
}

// Add 数组队列入队
func (q *ArrayQueue) Add(v int) {
    q.lock.Lock()
    defer q.lock.Unlock()
    q.array = append(q.array, v)
    q.size++
}

// Add 链表队列入队
func (q *LinkQueue) Add(v int) {
    q.lock.Lock()
    defer q.lock.Unlock()
    node := &LinkNode{Val: v, Next: q.root}
    q.root = node
    q.size++
}

// Remove 数组队列出队
func (q *ArrayQueue) Remove() (int, error) {
    q.lock.Lock()
    defer q.lock.Unlock()
    if q.size == 0 {
        return 0, errors.New("queue is empty")
    }
    v := q.array[0]
    q.array = q.array[1:]
    q.size--
    return v, nil
}

// Remove 链表队列出队
func (q *LinkQueue) Remove() (int, error) {
    q.lock.Lock()
    defer q.lock.Unlock()
    if q.size == 0 {
        return 0, errors.New("queue is empty")
    }
    v := q.root.Val
    q.root = q.root.Next
    q.size--
    return v, nil
}
