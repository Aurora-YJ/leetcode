package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 6, 3, 8}
	target := 10

	res := numSubseq(nums, target)
	fmt.Println("result is :", res)
}

func numSubseq(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	conter := 0

	for i := 0; i < len(nums); i++ {

		res := []int{}
		res2 := []int{}
		for j := i; j < len(nums[i:]); j++ {
			res = append(res, nums[j])

			if checker(res, target) {
				conter++
			}

			if j != i {
				res2 = append(res2, nums[i])
				res2 = append(res2, nums[j])
				if checker(res2, target) {
					conter++
				}

				res2 = []int{}

			}

		}

	}

	return conter
}

func checker(nums []int, t int) bool {
	sum := nums[0] + nums[len(nums)-1]
	if sum <= t {
		fmt.Println("--->", nums)
		return true
	}
	return false
}
