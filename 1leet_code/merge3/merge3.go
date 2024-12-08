package main

import "fmt"

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println(arr)
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)
}


func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(arr1 []int , arr2 []int) []int {
	var res []int
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			res = append(res, arr2[j])
			j++
		} else {
			res = append(res, arr1[i])
			i++
		}
	}
	res = append(res, arr1[i:]...)
	res = append(res, arr2[j:]...)

	return res
}