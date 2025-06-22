package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 1, 2, 1}
	fmt.Println("result--->", isArraySpecial(nums))

}

func isArraySpecial(nums []int) bool {

	if len(nums) == 0 {
		return true
	}

	help := nums[0] % 2 == 0

	for i := 1; i < len(nums); i++ {
		isgood := nums[i] %2==0
		if help == isgood {
			return  false
		}
		help = isgood
		
	}
	return true
}
