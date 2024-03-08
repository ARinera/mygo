package mygo

import (
    "errors"
    "sync"
)

// ArrayStack 数组栈
type ArrayStack struct {
    array []int      // 底层切片
    size  int        // 栈的元素数量
    lock  sync.Mutex // 为了并发安全使用的锁
}

// LinkStack 链表栈
type LinkStack struct {
    root *LinkNode  // 链表起点
    size int        // 栈的元素数量
    lock sync.Mutex // 为了并发安全使用的锁
}

// Push 数组栈入栈
func (s *ArrayStack) Push(v int) {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.array = append(s.array, v)
    s.size++
}

// Push 链表栈入栈
func (s *LinkStack) Push(v int) {
    s.lock.Lock()
    defer s.lock.Unlock()

    if s.root == nil {
        s.root = new(LinkNode)
        s.root.Val = v
    } else {
        preNode := s.root
        newNode := new(LinkNode)
        newNode.Val = v
        newNode.Next = preNode
        s.root = newNode
    }
    s.size++
}

// Pop 数组栈出栈
func (s *ArrayStack) Pop() (int, error) {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.size == 0 {
        return 0, errors.New("stack is empty")
    }
    v := s.array[s.size-1]
    s.array = s.array[:s.size-1]
    s.size--
    return v, nil
}

// Pop 链表栈出栈
func (s *LinkStack) Pop() (int, error) {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.size == 0 {
        return 0, errors.New("stack is empty")
    }
    v := s.root.Val
    s.root = s.root.Next
    s.size--
    return v, nil
}

// Peek 数组栈查看栈顶元素
func (s *ArrayStack) Peek() (int, error) {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.size == 0 {
        return 0, errors.New("stack is empty")
    }
    return s.array[s.size-1], nil
}

// Peek 链表栈查看栈顶元素
func (s *LinkStack) Peek() (int, error) {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.size == 0 {
        return 0, errors.New("stack is empty")
    }
    return s.root.Val, nil
}

// Size 数组栈查看栈大小
func (s *ArrayStack) Size() int {
    return s.size
}

// Size 链表栈查看栈大小
func (s *LinkStack) Size() int {
    return s.size
}

// IsEmpty 数组栈判断是否为空
func (s *ArrayStack) IsEmpty() bool {
    return s.size == 0
}

// IsEmpty 链表栈判断是否为空
func (s *LinkStack) IsEmpty() bool {
    return s.size == 0
}
