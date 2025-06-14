package main

import "fmt"

func main() {
	tests := [][]int{
		{3, 3, 3},   // equilateral
		{3, 4, 4},   // isosceles
		{3, 4, 5},   // scalene
		{1, 2, 3},   // none
		{5, 10, 25}, // none
	}

	for _, sides := range tests {
		fmt.Printf("Triangle sides %v → %s\n", sides, triangleType(sides))
	}
}

func triangleType(nums []int) string {
	if len(nums) != 3 {
		return "invalid"
	}
	n1 , n2 , n3 := nums[0] , nums[1] , nums[2]
	if n1+n2 <= n3 || n1+n3 <= n2 || n2+n3 <= n1 {
		return "none"
	}
 	mapp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		mapp[nums[i]] = true
	}

	if len(mapp) == 1 {
		return "equilateral"
	} else if len(mapp) == 3 {
		return "scalene"
	} else {
		return "Isosceles"
	}
}
