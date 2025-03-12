package main

import "fmt"

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
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	l1 := createLinkedList([]int{2, 4, 3}) 
	l2 := createLinkedList([]int{5, 6, 4}) 

	result := addTwoNumbers(l1, l2)

	
	printLinkedList(result) 
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	arr1 := []int{}
	arr2 := []int{}

	if l1 == nil {
		return nil
	}
	if l2 == nil {
		return nil
	}
	arr1 = append(arr1, l1.Val)
	arr2 = append(arr2, l2.Val)
	if l1.Next != nil{
		arr1 = append(arr1, l1.Next.Val)
	}
	if l2.Next != nil{
		arr2 = append(arr2, l2.Next.Val)
	}


	fmt.Println(arr1)
	fmt.Println(arr2)
	return nil
}
