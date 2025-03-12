package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func main() {
	l1 := createLinkedList([]int{2, 4, 3})
	l2 := createLinkedList([]int{5, 6, 4})

	result := addTwoNumbers(l1, l2)

	fmt.Println(result)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) []int {
	arr1 := []int{}
	

	for l1 != nil || l2 != nil {
		sum := l1.Val + l2.Val
		arr1 = append(arr1, sum)
		l1 = l1.Next
		l2 = l2.Next

	}

	fmt.Println(arr1)
	

	return nil
}
