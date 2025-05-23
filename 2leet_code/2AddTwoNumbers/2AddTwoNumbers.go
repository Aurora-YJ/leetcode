package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}



func main() {
    l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}

    l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

    result := addTwoNumbers(l1, l2)

    fmt.Print("Result: ")
    for result != nil {
        fmt.Print(result.Val)
        result = result.Next
    }
    fmt.Println()
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    newNode := &ListNode{}
    res := newNode
    md := 0
    for l1 != nil || l2 != nil || md != 0 {
        nmb1 , nmb2 := 0 , 0

        if l1 != nil {
            nmb1 = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            nmb2 = l2.Val
            l2 = l2.Next
        }

        sum := nmb1 + nmb2 + md
        md = sum / 10
        dv := sum % 10
        res.Next = &ListNode{Val: dv}
        res = res.Next
    }
    return newNode.Next
}