package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nm := removeDuplicates(nums)
	fmt.Println(nm)
}
func removeDuplicates(nums []int) int {
	nm := 0
	mmap := make(map[int]int)

	for _, s := range nums {
		mmap[s] = 0
	}

	for _, s := range nums {
		if mmap[s] == 0 {
			mmap[s]++
			nums[nm] = s
			nm++
		}
	}

	return nm
}
