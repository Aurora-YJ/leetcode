package main

import (
	"fmt"
)

func main() {
	digits := []int{3,5,7}
	res := findEvenNumbers(digits)
	fmt.Println("---->",res)
}

func findEvenNumbers(digits []int) []int {
	var res []int
	mapp := make(map[int]bool)

	for i := range digits {
		for j := range digits {
			for k := range digits {
				if i != j && j != k  && k != i {
					r := digits[i] * 100 + digits[j] * 10 + digits[k]
					if r >= 100 && r < 1000 && r % 2 == 0 {
						mapp[r] = true
					}
				}
			}			

		}
	}

	for v := range mapp {
		res = append(res, v)
	}

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res)-1-i; j++ {
			if  res[j] > res[j+1] {
				res[j] , res[j+1] = res[j+1] , res[j]
			}
		}
	}

	return res
}
