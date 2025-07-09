package main

import "fmt"

func main() {
	nums := []int{2,1,3,3}
	k := 2

	result := maxSubsequence(nums, k)

	fmt.Println("Max subsequence of length", k, ":", result)
}

func maxSubsequence(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}

	res := []int{}

	mapp := make(map[int]int)

	res = append(res, nums...)

	for k > 0 {
		index := checker(res)
		mapp[res[index]]++

		res = append(res[:index], res[index+1:]...)
		k--
	}
	result := []int{}
	for h := range mapp {

		for i := 0; i < len(nums); i++ {
			if h == nums[i] {
				result = append(result, h)
			}
			if mapp[h] >= 1 {
				mapp[h]--
			} else {
				break
			}
		}
	}

	return result
}

func checker(res []int) int {
	help := res[0]
	in := 0
	for i := 1; i < len(res); i++ {
		if help < res[i] {
			help = res[i]
			in = i
		}
	}
	return in
}
