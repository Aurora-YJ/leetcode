package main

import "fmt"

func main(){

	r1 := []int{1,3}
	r2 := []int{2,4}
	
	fmt.Printf("%f\n",findMedianSortedArrays(r1,r2))
}

func findMedianSortedArrays(r1, r2 []int) float64 {
	var res []int
	res =append(res, r1...)
	res =append(res, r2...)

	for i:=1; i < len(res); i++ {
		key := res[i]
		j := i - 1

		for j >= 0 && res[j] > key {
			res[j+1] = res[j]
			j--
		}
		res[j+1] = key
	}
	
	if len(res)%2 == 0 {
		fmt.Println(res[len(res)/2])
		fmt.Println(res[len(res)/2 - 1])
		b := float64(res[len(res)/2] + res[len(res)/2 - 1])
		return float64(b/2)
	}
	return float64(res[len(res)/2])
}