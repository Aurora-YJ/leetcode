package main

import (
	"fmt"
)

func main() {
	// nums := []int{-56, -214, -345, 952, 529, -294, -325, -924, -27, -41, 716, -67, -65, -481}
	nums := []int{
		-16, -13, 8, 16, 35, -17, 30, -8, 34, -2, -29, -35, 15, 13, -30, -34, 6, 15, 28, -23, 34, 28, -24, 15, -17, 10, 31, 32, -3,
		-36, 19, 31, -5, -21, -33, -18, -23, -37, -15, 12, -28, -40, 1, 38, 38, -38, 33, -35, -28, -40, 4, -15, -29, -33, -18, -9, -29, 20, 1,
		36, -8, 23, -34, 16, -7, 13, 39, 38, 7, -7, -10, 30, 9, 26, 27, -37, -18, -25, 14, -36, 23, 28, -15, 35, -9, 1,
	}

	// nums := []int{2, 5, 3, 5}
	k := 8
	/*

				Output
		[35,34,34,38,38,36,39,38]
		Expected
		[35,34,38,38,36,39,38,35]

	*/

	result := maxSubsequence(nums, k)
	fmt.Println("Max subsequence of length", k, ":", result)
}

func maxSubsequence(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}
	n := len(nums)
	res := []int{}

	res = append(res, nums...)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	mapp := make(map[int]int)

	for _, v := range nums[:k] {
		mapp[v]++
	}
	result := []int{}

	for _, v := range res {
		for i := 0; i < len(nums[:k]); i++ {
			if mapp[v] < 1 {
				break
			} 
			if nums[:k][i] == v {
				result = append(result, v)
				mapp[v]--
				break
			}
		}
	}

	return result
}
