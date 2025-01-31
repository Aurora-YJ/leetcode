package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{1, &ListNode{3, &ListNode{5,&ListNode{5, &ListNode{1, &ListNode{5, nil}}}}}}
	list2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}

	list3 := mergeTwoLists(list1, list2)

	printList(list3)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	d := &ListNode{}
	c := d

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			c.Next = list1
			list1 = list1.Next

		} else {
			c.Next = list2
			list2 = list2.Next
		}
		c = c.Next
	}

	if list1 != nil {
		c.Next = list1
	} else {
		c.Next = list2
	}
	return d.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}