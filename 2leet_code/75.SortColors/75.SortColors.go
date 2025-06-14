package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println("result--->", sortColors(nums))

}

func sortColors(nums []int) []int {
	clr1, clr2, clr3 := 0, 0, len(nums)-1

	for clr2 <= clr3 {
		if nums[clr2] == 0 {
			helper1 := nums[clr1]
			nums[clr1] = nums[clr2]
			nums[clr2] = helper1
			clr1++
			clr2++
		} else if nums[clr2] == 1 {
			clr2++
		} else if nums[clr2] == 2 {
			helper2 := nums[clr3]
			nums[clr3] = nums[clr2]
			nums[clr2] = helper2
			clr3--
		}
	}

	return nums

}
