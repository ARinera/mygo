package mygo

import (
    "errors"
    "fmt"
)

// LinkNode 定义链表节点结构体
type LinkNode struct {
    Val  int
    Next *LinkNode
}

// NewLinkList NewListList 通过[]int创建一个链表
func NewLinkList(nums []int) (*LinkNode, error) {
    if len(nums) == 0 {
        err := errors.New("nums is empty")
        return nil, err
    }

    dummy := &LinkNode{}
    curr := dummy

    for _, v := range nums {
        curr.Next = &LinkNode{Val: v}
        curr = curr.Next
    }

    return dummy.Next, nil
}

// PrintLinkList 打印链表
func PrintLinkList(head *LinkNode) error {
    if head == nil {
        err := errors.New("listnode is nil")
        return err
    }

    curr := head
    for curr != nil {
        fmt.Printf("%d ", curr.Val)
        curr = curr.Next
    }
    fmt.Println()
    return nil
}
