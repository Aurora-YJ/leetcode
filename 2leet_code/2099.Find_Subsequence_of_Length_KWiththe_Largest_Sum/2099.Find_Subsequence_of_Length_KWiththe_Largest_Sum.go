package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, -2, 3, 4}
	k := 3

	result := maxSubsequence(nums, k)

	fmt.Println("Max subsequence of length", k, ":", result)
}

func maxSubsequence(nums []int, k int) []int {

	if k > len(nums) {
		return []int{}
	}

	mapp := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		mapp[nums[i]] = append(mapp[nums[i]], nums[i])

		res := []int{}
		for j := 0; j < len(nums); j++ {
			res = append(res, nums[j])
		}
		newRes := checker(res, k-1, nums[i])
		mapp[nums[i]] = append(mapp[nums[i]], newRes...)

	}

	hh := 0
	jj := 0
	result := []int{}
	for _,v := range mapp {

		for _,h := range v {
			hh += h
		}


		if hh > jj {
			jj = hh
			result = v
		}
		hh = 0

	}

	fmt.Println("jj", jj)

	return result

}

func checker(res []int, k int, intger int) []int {
	result := []int{}
	for k > 0 {
		helper := 0
		index := 0

		for i := 0; i < len(res); i++ {
			if res[i] == intger {
				continue
			}
			if helper == 0 {
				helper = res[i]
				index = i
				continue
			}

			if helper < res[i] {
				helper = res[i]
				index = i
			}

		}
		res = append(res[:index], res[index+1:]...)
		result = append(result, helper)
		k--
	}
	return result
}
