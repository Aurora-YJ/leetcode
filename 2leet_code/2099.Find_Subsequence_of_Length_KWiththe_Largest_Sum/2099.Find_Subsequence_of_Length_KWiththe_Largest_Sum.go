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

	lenn := len(nums)
	lastres := 0


	ll := []int{}

	for m := 0; m < k; m++ {
		ll =append(ll, nums[m])
		lastres += nums[m]
	}

	lenn--

	j := 1
	for k <= lenn {
		res := 0
	    result := []int{}


		H:=0
		for i := j; i < len(nums); i++ {
			if H  < k {
				res += nums[i]
				H++
				result =append(result, nums[i])
			} else  {
				break
			}
		}

		if res > lastres {
			lastres = res
			res = 0
			ll = result
		} 

		lenn--
		j++
	}

	return ll

}
