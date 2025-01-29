package main

import "fmt"

func main(){

	result := []int{3,2,2,3}

	k := removeElement(result, 2)

	fmt.Println(k, result[:k]) 
	fmt.Println(result)
}


// removeElement removes all occurrences of 'val' from 'nums' in-place
// and returns the count of remaining elements.
func removeElement(nums []int, val int) int {
    
    k := 0

    for _, n := range nums {
        if n != val {
            nums[k] = n
			k++
        } 
    }
    return k
}