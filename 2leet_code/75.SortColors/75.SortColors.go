package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println("result--->", sortColors(nums))

}

func sortColors(nums []int) []int {
	c1, c2, c3 := 0, 0, len(nums)-1

	for c2 <= c3 {
		switch nums[c2] {
		case 0:
			nums[c1], nums[c2] = nums[c2], nums[c1]
			c2++
			c1++
		case 1:
			c2++
		case 2:
			nums[c3], nums[c2] = nums[c2], nums[c3]
			c3--
		}
	}
	return nums
}
