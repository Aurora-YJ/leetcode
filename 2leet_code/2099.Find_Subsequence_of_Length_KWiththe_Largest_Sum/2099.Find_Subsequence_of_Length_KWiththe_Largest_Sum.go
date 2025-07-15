package main

import (
	"fmt"
)

func main() {
	nums := []int{7,3,9,6}
	k := 2

	result := maxSubsequence(nums, k)

	fmt.Println("Max subsequence of length", k, ":", result)
}

func maxSubsequence(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}

	res := []int{}

	return res
}
