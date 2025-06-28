package main

import "fmt"

func main() {
	nums := []int{3, 5, 2, 9, 1}
	k := 3

	result := maxSubsequence(nums, k)

	fmt.Println("Max subsequence of length", k, ":", result)
}

func maxSubsequence(nums []int, k int) []int {

	res := 0
	lastRes := 0
	cont := 0
	i := 0
	for  i < len(nums){
		if cont != k {
			res += nums[i]
			cont++

		} else {
			if res > lastRes{
				lastRes = res
				res = 0
				cont = 0
				i = i - k -1
			} else  {
				res = 0
				cont = 0
				i = i - k -1
			}
		}

	}
	fmt.Println("jj" , lastRes, res)
	return []int{lastRes}

}
