package main

import "fmt"

func main() {
	nums := []int{500 , 100, 444}
	k := 2

	result := maxSubsequence(nums, k)

	fmt.Println("Max subsequence of length", k, ":", result)
}

func maxSubsequence(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}

	indexRes := checker(nums, k)

	newRes := []int{}
	for i := 0; i < len(indexRes); i++ {
		newRes = append(newRes, nums[indexRes[i]])
	}

	return newRes
}

func mysort(res []int) {
	if len(res) < 2 {
		return 
	}
	if res[len(res)-1] < res[len(res)-2] {
		swp := res[len(res)-1]
		res[len(res)-1] = res[len(res)-2]
		res[len(res)-2] = swp
	}
}

func checker(res []int, k int) []int {
	result := []int{}
	for k > 0 {
		helper := res[0]
		in := 0
		
		for i := 1; i < len(res); i++ {
		
			if !jj(result, i) && helper < res[i] {
				helper = res[i]
				in = i

			} 
			// ila mal9itix li kbar mno

			fmt.Println("iiiiiiiiiiiiii", helper, in)
		}

		if jj(result, in) {
			continue
		} else {
			result = append(result, in)
			fmt.Println("-->", result)
			// if len(result) == 0 {
			// 	result = append(result, in)
			// } else {
			// 	 mysort(result)
			// }
		}
		fmt.Println("gg", result)

		k--
	}
	return result
}

func jj(res []int, j int) bool {
	for _, v := range res {
		if v == j {
			return true
		}
	}
	return false
}
