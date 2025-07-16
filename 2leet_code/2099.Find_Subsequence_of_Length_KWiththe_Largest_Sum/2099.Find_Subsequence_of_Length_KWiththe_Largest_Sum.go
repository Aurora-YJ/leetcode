package main

import (
	"fmt"
)

func main() {
	nums := []int{-56, -214, -345, 952, 529, -294, -325, -924, -27, -41, 716, -67, -65, -481}
	k := 5

	/*

		[952,716,-67,-65,-481]
		Expected
		[952,529,-27,-41,716]

	*/

	result := maxSubsequence(nums, k)

	fmt.Println("Max subsequence of length", k, ":", result)
}

func maxSubsequence(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}

	res := []int{}

	toRemove := len(nums) - k
	for i := 0; i < len(nums); i++ {
		for len(res) > 0 && res[len(res)-1] < nums[i] && toRemove > 0 && len(nums[i:]) > toRemove {
			res = res[:len(res)-1]
			toRemove--
		}
		res = append(res, nums[i])
	}

	return res[:k]
}
