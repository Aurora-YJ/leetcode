package main

import "fmt"

func main() {
	nums := []int{50,-75 , 7}
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


func mysort(res []int , in int) []int {
	for i := 0; i < len(res); i++ {
		if in < res[i] {
			res = append(res[:i] , append([]int{in} , res[i:]...)...)
			return res
		} 
	}
	
	res = append(res, in)
	return  res
}

func checker(res []int, k int) []int {
	result := []int{}
	for k > 0 {
		helper := 0
		in := 0
		for i := 0; i < len(res); i++ {
			if helper == 0 {
				helper = res[i]
				in = i
				continue
			}
			
			if helper < res[i] && !jj(result, i){
				helper = res[i]
				in = i
			
			}
			// ila mal9itix li kbar mno 

			
		}
		
		if jj(result , in) {
			continue
		} else {
			if len(result) == 0 {
				result = append(result, in)
			} else {
				result = mysort(result, in)
			}
		}
		fmt.Println("gg", result)

		k--
	}
	return result
}

func jj(res []int, j int) bool{
	for _, v := range res{
		if v == j {
			return true
		}
	}
	return false
}