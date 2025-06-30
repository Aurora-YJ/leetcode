package main

import "fmt"

func main() {
	nums := []int{-1,-2,3,4}
	k := 3

	result := maxSubsequence(nums, k)

	fmt.Println("Max subsequence of length", k, ":", result)
}

func maxSubsequence(nums []int, k int) []int {

	if k > len(nums) {
		return []int{}
	}

	lastres := 0

	ll := []int{}

	for i := 0; i < len(nums); i++ {
		res := 0
		res += nums[i]
		v := 0
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			v++
			if v  <k {
				res += nums[j]
			} else {
				break
			}
			
			fmt.Println("jjj", res)
		}
		fmt.Println("-------")
		if res > lastres {
			lastres = res
			res = 0
		} 
	}
	fmt.Println("kk", lastres)
	return ll

}
